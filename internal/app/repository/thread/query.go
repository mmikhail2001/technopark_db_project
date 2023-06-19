package thread

var (
	// insert or update on table "threads" violates foreign key constraint "threads_forum_fkey"
	// insert or update on table "threads" violates foreign key constraint "threads_author_fkey"
	queryInsertThread = `insert into threads (title, message, author, forum, slug, created)
    values
        ($1, $2, $3, $4, $5, $6)
			on conflict on constraint threads_slug_unique_or_null_idx do nothing
            	returning id, title, author, forum, message, sum_votes, slug, created;`

	querySelectThreadBySlug = `select id, title, author, forum, message, sum_votes, slug, created
    from threads
        where slug = $1;`
	querySelectThreadByID = `select id, title, author, forum, message, sum_votes, slug, created
	from threads
		where id = $1;`
	querySelectThreadsByForumSlug_begin = `select t.id, t.title, t.author, 
		t.forum, t.message, t.sum_votes, t.slug, t.created
    		from threads t 
				where t.forum = $1 `

	queryUpdateThread = `
	UPDATE threads
	SET title   = COALESCE(NULLIF(TRIM($1), ''), title),
		message = COALESCE(NULLIF(TRIM($2), ''), message)
	WHERE id = $3
	RETURNING id, title, author, forum, message, sum_votes, slug, created;`
)
