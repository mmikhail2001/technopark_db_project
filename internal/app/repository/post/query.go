package post

var (
	querySelectPostByID = `select p.id, p.parent, p.author, p.message, p.is_edited, p.forum, p.thread, p.created
    from posts p
        where p.id = $1;`
	queryUpdateMessage = `update posts p set message = coalesce(nullif(trim($1), ''), message)
    	where p.id = $2
        	returning p.id, p.parent, p.author, p.message, p.is_edited, p.forum, p.thread, p.created;`
	queryInsertPosts_begin = `insert into posts 
								(parent, author, message, forum, thread, created)
    								values `
	querySelectPostsWithSort_Flat_Tree_Begin = `select id, parent, author, message, is_edited, forum, thread, created 
    from posts
        where thread = $1`
)
