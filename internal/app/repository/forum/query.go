package forum

var (
	queryInsertForum = `insert into forums (slug, title, author)
		values
			($1, $2, $3)
				on conflict(slug) do nothing
					returning title, author, slug, count_posts, count_threads;`
	querySelectForum = `select slug, title, author, count_posts, count_threads
						from forums
							where slug = $1;`
)
