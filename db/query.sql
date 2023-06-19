-- запросы

-- 1 /forum/create
-- Создание форума

-- ====== [ START ] ======
insert into forums (slug, title, author) 
    values 
        ('social-democratic', 'test-title-4', 'jfrank')
            on conflict(slug) do nothing
                returning title, author, slug, count_posts, count_threads;
-- если вернулись данные, то 201 и возвращаем их

-- если ошибка ERROR: insert or update on table "forums" violates foreign key constraint "forums_author_fkey"
-- то вернуть ошибку и код 404

-- если ни одной строки не вернулось, то запросить 
select slug, title, author, count_posts, count_threads
    from forums
        where slug = 'medical-early';

-- и вернуть 409

-- ====== [ END ] ======
select nickname from users limit 10;
select * from forums;
select * from posts where parent <> 0 limit 200 offset 2000;




-- 2 /forum/{slug}/details
-- Получение информации о форуме

-- ====== [ START ] ======
select slug, title, author, count_posts, count_threads
    from forums
        where slug = 'social-democratic';

-- если вернулась строка, возвращаем 200 
-- если ничего не вернулось, возвращаем ошибку 404 

-- ====== [ END ] ======


-- 3 /forum/{slug}/create
-- Создание ветки

select count(*) from 
ALTER SEQUENCE threads_id_sec RESTART WITH 4615793;

-- ====== [ START ] ======

insert into threads (title, message, author, forum, slug, created)
    values
        ('test-thread-title', 'test-message', 'fischerkari_', 'heart-beyond-happen', 'test-slug-2', '2023-06-08')
            on conflict on constraint threads_slug_unique_or_null_idx do nothing
                returning id, title, author, forum, message, sum_votes, slug, created;

INSERT INTO table1 (col1_1, col1_2, col1_3)
SELECT 'A', 'B', 0
WHERE EXISTS (SELECT * FROM table2 where col2_2 = 'B')
    AND NOT EXISTS (SELECT * FROM table1 where col1_1 = 'A');

select id, title, author, forum, message, sum_votes, slug, created
    from threads
        where slug = 'test-slug';

-- если вернулась строка ответ 201

-- ERROR: insert or update on table "threads" violates foreign key constraint "threads_author_fkey"
-- ERROR: insert or update on table "threads" violates foreign key constraint "threads_forum_fkey"
-- на эти ошибки ответ 404

-- ошибка существования ветки можно проверить только по слагу, но он не передается, а может и передается (НА ТЕСТАХ УЗНАЕМ)

-- ====== [ END ] ======


-- 4 /forum/{slug}/users
-- Пользователи данного форума
-- Получение списка пользователей, у которых есть пост или ветка обсуждения в данном форуме.

-- "current-eye-live"
-- "investment-player"
-- "admit-many-weight"
-- "each-hair-test"
-- "skin-rate-more"
-- "you-son-sign"
-- "what-significant"
-- "against-draw"
-- "heart-during-region"
-- "treat-lose-able"
-- "sometimes-east-me"
-- "believe-talk-ok"
-- "street-upon"
-- "protect-action"
-- "none-serve-position"
-- "wrong-which"
-- "she-somebody-lead"
-- "near-rise-single"
-- "heart-beyond-happen"

-- ====== [ START ] ======

select uf.nickname, uf.fullname, uf.about, uf.email
		from users_forums uf
            where uf.forum = 'fcay1Wv-335G8' 
				and uf.nickname > 'b'
                    order by uf.nickname
                        limit 500;
-- ====== [ END ] ======


-- старый запрос
select uu.nickname, uu.fullname, uu.about, uu.email from
	(select u.nickname nickname, u.fullname, u.about, u.email
		from users u
			join posts p on p.author = u.nickname and p.forum = 'skin-rate-more' 
				and lower(u.nickname) > 'b'
	union 
	select u.nickname nickname, u.fullname, u.about, u.email
		from users u
			join threads t on t.author = u.nickname and t.forum = 'skin-rate-more' 
				and lower(u.nickname) > 'b') uu
	order by (lower(uu.nickname))
		limit 500;
        

select u.nickname, u.fullname, u.about, u.email
    from users u, posts p, threads t, forums f
        where (p.author = u.nickname or t.author = u.nickname) and  p.forum = 'skin-rate-more'
            and u.nickname > 'b'
                order by lower(u.nickname)
                    limit 100;



-- 5 /forum/{slug}/threads
-- Список ветвей обсужления форума

-- ====== [ START ] ======

select t.id, t.title, t.author, t.forum, t.message, t.sum_votes, t.created
    from threads t 
        where t.forum = 'none-director-check'
            and t.created >= '2023-06-08'
                order by t.created
                    limit 100;

-- если ничего не выдало, то нужно проверять, форума нет или условие t.created >= '2023-06-08' такое, что под него ничего не попадает

select count(*) from forums where slug = 'none-director-check';

-- ====== [ END ] ======

-- 6 /post/{id}/details
-- Получение информации о посте

-- ====== [ START ] ======

select p.id, p.parent, p.author, p.message, p.is_edited, p.forum, p.thread, p.created
    from posts p
        where p.id = 101;
-- ====== [ END ] ======

-- если ничего не вывело, то возвращаем ошибку, что id неправильный

-- получение информации о user = p.author, forum = p.forum, thread = p.thread вызовом соответствующих функций



-- 7 /post/{id}/details
-- Обновление информации о посте

-- ====== [ START ] ======

update posts p set message = 'new message 2'
    where p.id = 102
        returning p.id, p.parent, p.author, p.message, p.is_edited, p.forum, p.thread, p.created;
-- ====== [ END ] ======

-- если ничего не вывело, то возвращаем ошибку, что id неправильный

-- 8 /service/clear
-- Очистка всех данных в базе

-- НЕ ПРОВЕРЯЛ

-- ====== [ START ] ======

begin;
delete from posts;
delete from votes;
delete from threads;
delete from forums;
delete from users;
commit;

TRUNCATE TABLE posts CASCADE;
TRUNCATE TABLE votes CASCADE;
TRUNCATE TABLE threads CASCADE;
TRUNCATE TABLE forums CASCADE;
TRUNCATE TABLE users CASCADE;


-- ====== [ END ] ======

-- ответ 200

-- 9 /service/status
-- Получение инфомарции о базе данных

-- ====== [ START ] ======

begin;
select count(*) from users;
select count(*) from forums;
select count(*) from threads;
select count(*) from posts;
commit;


select reltuples::bigint as estimate
from pg_class
where oid = 'posts'::regclass;

select reltuples::bigint as estimate
from pg_class
where oid = 'forums'::regclass;

select reltuples::bigint as estimate
from pg_class
where oid = 'threads'::regclass;

select reltuples::bigint as estimate
from pg_class
where oid = 'users'::regclass;
-- ====== [ END ] ======


-- 10 /thread/{slug_or_id}/create
-- Создание новых постов

select nickname from users limit 10;

-- Все посты, созданные в рамках одного вызова данного метода должны иметь одинаковую дату создания (Post.Created).
-- видимо, стоит дату создания в go приложении генерировать

-- ====== [ START ] ======

insert into posts (parent, thread, author, message)
    values (0, 10, 'cdavis', 'test-message-10'),
            (0, 10, 'cdavis', 'test-message-11'),
            (23232323, 121, 'cdavis', 'test-message-12')
                returning id, parent, author, message, is_edited, forum, thread, created;
-- ====== [ END ] ======

-- при добавлении без форума срабатывает триггер, который ищет форум
-- мб неоптимально

-- 201 вернулась запись
-- 404 ERROR: thread with id 121212121 does not exist
-- 409 ERROR: parent with id 23232323 does not exist

-- 11 /thread/{slug_or_id}/details
-- Получение информации о ветке обсуждения

-- ====== [ START ] ======

select id, title, author, forum, message, sum_votes, created
    from threads 
        where id = 42;

-- ====== [ END ] ======

-- не вывел ни одну строку -> ветка отсутствует 404


-- 12 /thread/{slug_or_id}/details
-- Обновление ветки

-- ====== [ START ] ======

update threads set title = 'new-test-title', message = 'new-test-message' 
    where id = 5
    returning id, title, author, forum, message, sum_votes, created;
-- ====== [ END ] ======


-- вывел строку - 200
-- ни одной строки не вывел, значит ветки не существует 404

-- 13 /thread/{slug_or_id}/posts
-- Сообщения данной ветви обсуждения

-- ====== [ START ] ======

-- sort = flat 
select id, parent, author, message, is_edited, forum, thread, created 
    from posts
        where thread = 3 and id > 1689
            order by created
                limit 100;

-- sort = tree  

select id, parent, author, message, is_edited, forum, thread, created 
    from posts
        where thread = 3 and id > 1689
            order by path
                limit 10;

-- sort = parent_tree   

with root_posts as (
  select id
    from posts
      where parent = 0 
        and id > 4 and thread = 1
        order by id
          limit 5 
) select path
    from posts
      where path[1] in (select id from root_posts)
        order by path;


with root_posts as (
  select id
    from posts
      where parent = 0 
        and thread = 1
        order by id
          limit 100
) select path
    from posts
      where path[1] in (select id from root_posts)
        order by path;


select path
    from posts
	where path[1] in 
        (select id from posts 
            where thread = 1 and parent = 0 and path[1] >
				(select path[1] from posts where id = 2) 
        order by id asc limit 5) 
			order by path asc, id asc;

-- ====== [ END ] ======

-- если ничего не вернулось, то нужно проверить существование ветки, потому что условие id > [id] может не подходить ни под одну запись


-- 14 /thread/{slug_or_id}/vote
-- Проголосовать за ветвь обсуждения

select nickname from users limit 10;

-- ====== [ START ] ======

insert into votes (author, thread, vote)
    values
        ('brandon82', 100000, 1)
        on conflict on constraint votes_pkey
            do update set vote = excluded.vote;

-- 404 ветка не найдена - ERROR: insert or update on table "votes" violates foreign key constraint "votes_thread_fkey" 

-- присутствует триггер, который сразу обновит threads.sum_votes, однако всю запись thread все равно нужно запрашивает, так как ее надо вернуть

-- возможно стоит убрать триггер и обновлять запись через update returning
-- однако как мы узнаем, какой рейтинг стоял у пользователя до обновления (нужно сначала запросить, потом обновить....) ....
-- тут уж непонятно, что выбирать

select id, title, author, forum, message, sum_votes, created 
    from threads
        where id = 10; 

-- ====== [ END ] ======


select sum_votes from threads where id = 10;

select author, thread, vote from votes where author = 'brandon82' and thread = 10;

-- ====== [ END ] ======

-- 15 /user/{nickname}/create
-- Создание нового пользователя

insert into users (nickname, fullname, about, email)
    values
        ('test-nick-name', 'test-fullname', 'test-about', 'test@mail.ru')
            on conflict (nickname) do nothing
            returning nickname, fullname, about, email;

-- если ошибка, то возвращаем старого пользователя, а это новый запрос
-- ERROR: duplicate key value violates unique constraint "users_pkey"

select nickname, fullname, about, email
    from users 
        where nickname = 'test-nick-name' or email = 'test@mail.ru';

-- 16 /user/{nickname}/profile
-- Получение информации о пользователе

select nickname, fullname, about, email
    from users
        where nickname = 'test-nick-name';

-- если ничего не вывел, то ошибка

-- 17 /user/{nickname}/profile
-- Изменение данных о пользователе

update users set 
    fullname = 'test-fullname',
    about = 'test-about',
    email = 'test1@mail.ru'
            where nickname = 'test-nick-name';

-- пользователь не существует - ничего не выведется
-- конфиликт - ERROR: duplicate key value violates unique constraint "users_email_key"


insert into posts (parent, author, message, forum, thread, created)
    values  (1,'john37','We should be afraid of the Kraken.','protect-action',1066,'2023-06-13 13:04:29')
        returning id;

$1 = '0', $2 = 'john37', $3 = 'We should be afraid of the Kraken.', $4 = 'protect-action', $5 = '1066', $6 = '2023-06-13 13:04:29', $7 = '1', $8 = 'john37', $9 = 'We should be afraid of the Kraken.', $10 = 'protect-action', $11 = '1066', $12 = '2023-06-13 13:04:29'
