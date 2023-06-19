package vote

var (
	queryInsertVote = `insert into votes (author, thread, vote)
    values
        ($1, $2, $3)
        on conflict on constraint votes_pkey
            do update set vote = excluded.vote;`
)
