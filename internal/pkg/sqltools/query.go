package sqltools

import (
	"context"
	"database/sql"
)

func RunQuery(ctx context.Context, db *sql.DB, action func(ctx context.Context, conn *sql.Conn) error) error {
	conn, _ := db.Conn(ctx)
	defer conn.Close()

	err := action(ctx, conn)
	if err != nil {
		return err
	}

	return nil
}
