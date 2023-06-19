package user

var (
	querySelectUserByNickname = `select nickname, fullname, about, email
    from users
        where nickname = $1;`
	querySelectUserByEmail = `select nickname, fullname, about, email
		from users
			where email = $1;`
	querySelectUserByNicknameOrEmail = `select nickname, fullname, about, email
		from users
			where nickname = $1 or email = $2;`

	querySelectUsersByForumSlug_begin = `select uf.nickname, uf.fullname, 
										uf.about, uf.email
										from users_forums uf
										where uf.forum = $1 `
	queryInsertUser = `insert into users (nickname, fullname, about, email)
    values
        ($1, $2, $3, $4)
		on conflict do nothing
            returning nickname, fullname, about, email;`
	queryUpdateUser = `
			UPDATE users
			SET fullname = COALESCE(NULLIF(TRIM($1), ''), fullname),
				about    = COALESCE(NULLIF(TRIM($2), ''), about),
				email    = COALESCE(NULLIF(TRIM($3), ''), email)
					WHERE nickname = $4
			RETURNING fullname, about, email, nickname;`
)
