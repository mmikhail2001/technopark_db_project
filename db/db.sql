drop schema public cascade;
create schema public;

create extension if not exists citext;

create unlogged table users (
    -- проверка регулярным выражением происходит в приложении
    nickname citext collate "ucs_basic" primary key,
    email citext not null unique,
    fullname text not null,
    about text not null
);

create unlogged table forums (
    slug citext primary key, 
    title text not null,
    author citext not null references users(nickname) on delete cascade,
    count_posts bigint default 0,
    count_threads bigint default 0
);

create unlogged table threads (
    id bigserial primary key,
    -- проверка регулярным выражением происходит в приложении
    author citext not null references users(nickname) on delete cascade,
    forum citext not null references forums(slug) on delete cascade,
    slug citext,
    title text not null,
    message text not null,
    created timestamp with time zone not null default now(),
    sum_votes bigint default 0
);

-- разрешено, но не работает on conflict: create unique index threads_slug_unique_or_null_idx on threads (slug) where slug is not null;

-- запрещено: alter table threads add constraint threads_slug_unique_or_null_idx unique (slug) where slug is not null;

-- работает ->
alter table threads add constraint threads_slug_unique_or_null_idx exclude (slug with =) where (slug is not null);

create unlogged table posts (
    id bigserial primary key,
    parent bigint not null,
    path bigint[] not null,
    author citext not null references users(nickname) on delete cascade,
    forum citext not null references forums(slug) on delete cascade,
    thread bigint not null references threads(id) on delete cascade,
    message text not null,
    is_edited boolean not null default false,
    created timestamp with time zone not null default now()
);

create unlogged table votes (
    author citext not null references users(nickname) on delete cascade,
    thread bigint not null references threads(id) on delete cascade,
    vote smallint not null constraint vote_check check (vote = -1 or vote = 1),
    primary key(author, thread)
);

create unlogged table users_forums (
    nickname citext collate "ucs_basic" not null references users(nickname) on delete cascade,
    email citext not null,
    fullname text not null,
    about text,
    forum citext not null references forums(slug) on delete cascade,
    unique(nickname, forum)
);

create index forums_author_idx on forums(author);

create index threads_author_idx on threads(author);
create index threads_forum_idx on threads(forum);
create index threads_forum_hash_idx on threads using hash (forum);
create index threads_created_idx on threads(created);

-- с ним на 40 мкс быстрее
create index threads_slug_hash_idx on threads using hash (slug);

create index posts_parent_idx on posts(parent);
create index posts_path_idx on posts(path);
create index posts_path_1_idx on posts((path[1]));

create index posts_path_1_path_idx on posts((path[1]), path); -- для сортировки при sort=parent_tree


create index posts_author_idx on posts(author);
create index posts_forum_idx on posts(forum);
create index posts_thread_idx on posts(thread);
create index posts_thread_hash_idx on posts using hash (thread);
create index posts_created_idx on posts(created);

create index users_nickname_idx on users(nickname);
create index users_email_idx on users(email);

create index users_forums_nickname_idx on users_forums(nickname);
create index users_forums_email_idx on users_forums(email);
create index users_forums_forum_idx on users_forums(forum);


-- обновление posts.path при вставке записи с известным parent

create or replace function update_posts_path()
returns trigger as $$
declare
  parent_path int[];
begin
  if new.parent = 0 then
    new.path = array[new.id];
  else
    select path into parent_path from posts where id = new.parent and thread = new.thread;
    if not found then
      raise exception 'parent with id % does not exist', new.parent;
    end if;
    new.path = parent_path || new.id;
  end if;
  return new;
end;
$$ language plpgsql;

create trigger update_posts_path_trigger
    before insert on posts
    for each row
    execute function update_posts_path();

-- обновление forums.count_posts при вставке записи в posts

create or replace function increment_forums_count_posts()
returns trigger as $$
begin
  update forums set count_posts = count_posts + 1 where slug = new.forum;
  return new;
end;
$$ language plpgsql;

create trigger insert_into_posts_trigger
after insert on posts
for each row
execute function increment_forums_count_posts();

-- обновление forums.count_threads при вставке записи в threads

create or replace function increment_forums_count_threads()
returns trigger as $$
begin
  update forums set count_threads = count_threads + 1 where slug = new.forum;
  return new;
end;
$$ language plpgsql;

create trigger insert_into_threads_trigger
after insert on threads
for each row
execute function increment_forums_count_threads();

-- обновление threads.sum_votes при вставке или обновлении записей в votes
-- 1. при обновлении записей в votes

create or replace function update_threads_sum_votes()
returns trigger as $$
begin
  if old.vote != new.vote then
    update threads set sum_votes = sum_votes - old.vote + new.vote where id = new.thread;
  end if;
  return new;
end;
$$ language plpgsql;

create trigger update_votes_trigger
after update on votes
for each row
execute function update_threads_sum_votes();

-- 2. при вставке записей в votes

create or replace function increment_threads_sum_votes()
returns trigger as $$
begin
  update threads set sum_votes = sum_votes + new.vote where id = new.thread;
  return new;
end;
$$ language plpgsql;

create trigger insert_into_votes_trigger
after insert on votes
for each row
execute function increment_threads_sum_votes();

-- обновление posts.is_edited при обновлении posts.message

create or replace function set_posts_is_edited()
returns trigger as $$
begin
  if old.message != new.message then
    new.is_edited = true;
  end if;
  return new;
end;
$$ language plpgsql;

create trigger update_posts_trigger
before update on posts
for each row
execute function set_posts_is_edited();

-- добавление записи в users_forums при добавлении записей в posts и threads

create or replace function update_users_forums_optimization() returns trigger as $$
declare
    _nickname text;
    _fullname text;
    _about    text;
    _email    text;
begin
    select u.nickname, u.fullname, u.about, u.email
      from users u
        where u.nickname = new.author
        into _nickname, _fullname, _about, _email;

    insert into users_forums (nickname, fullname, about, email, forum)
    values (_nickname, _fullname, _about, _email, new.forum)
    on conflict do nothing;
    return new;
end;
$$ language plpgsql;


create trigger insert_threads_trigger_for_user_forums_optimization
    after insert
    on threads
    for each row
execute procedure update_users_forums_optimization();

create trigger insert_posts_trigger_for_user_forums_optimization
    after insert
    on posts
    for each row
execute procedure update_users_forums_optimization();

-- обновление данных в users_forums при обновлении пользователя в users

create or replace function update_users_forums_after_update_user()
returns trigger as $$
begin
    update users_forums
    set fullname = new.fullname,
        about = new.about,
        email = new.email
    where nickname = new.nickname;

    return new;
end;
$$ language plpgsql;

create trigger update_users_forums_trigger
after update on users
for each row
execute function update_users_forums_after_update_user();


vacuum analyse;

