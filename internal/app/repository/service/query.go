package service

var (
	queryCountUsersStatistical = `select reltuples::bigint as estimate
		from pg_class
		where oid = 'users'::regclass;`
	queryCountPostsStatistical = `select reltuples::bigint as estimate
		from pg_class
		where oid = 'posts'::regclass;`
	queryCountForumsStatistical = `select reltuples::bigint as estimate
		from pg_class
		where oid = 'forums'::regclass;`
	queryCountThreadsStatistical = `select reltuples::bigint as estimate
		from pg_class
		where oid = 'threads'::regclass;`
	queryTrancate                = `TRUNCATE TABLE forums, posts, threads, users_forums, users, votes CASCADE;`
	queryCountUsers              = `select count(*) from users;`
	queryCountForumsThreadsPosts = `select count(*),  
								coalesce(sum(count_threads), 0), 
								coalesce(sum(count_posts), 0)
								from forums;`

	queryCountAllEntities = `
	SELECT (SELECT count(*) FROM forums) AS forums,
		   (SELECT count(*) FROM posts)  AS posts,
		   (SELECT count(*) FROM threads) AS threads,
		   (SELECT count(*) FROM users)  AS users`
)
