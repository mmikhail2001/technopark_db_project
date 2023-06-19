package sqltools

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	"github.com/mmikhail2001/technopark_db_project/internal/pkg/config"
	"github.com/pkg/errors"

	_ "github.com/jackc/pgx/stdlib"
)

func NewClientPostgres(cfg config.PostgresConfig) (*sql.DB, error) {
	dsn := fmt.Sprintf("user=%s dbname=%s password=%s host=%s port=%s sslmode=%s",
		cfg.User,
		cfg.DBName,
		cfg.Password,
		cfg.Host,
		cfg.Port,
		cfg.SSLmode)
	conn, err := sql.Open("pgx", dsn)
	if err != nil {
		return nil, errors.Wrapf(err, "not connect with dsn '%s'", dsn)
	}

	conn.SetMaxOpenConns(cfg.MaxOpenCons)

	wait := time.Now().Add(time.Second * 10)

	for time.Now().Before(wait) {

		err = conn.Ping()
		if err == nil {
			log.Println("ping sucessful")
			return conn, nil
		}

		time.Sleep(time.Second)
	}
	return conn, fmt.Errorf("connection failed")
}
