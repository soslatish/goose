package database

import (
	"context"
	"fmt"
	"time"

	"goose/internal/config"

	"github.com/jackc/pgx/v5"
	"github.com/pressly/goose/v3"
)

func ConnectDB(cfg *config.Config) (*pgx.Conn, error) {
	dsn := fmt.Sprintf("postgres://%s:%s@%s:%d/%s?sslmode=%s",
		cfg.Database.User,
		cfg.Database.Password,
		cfg.Database.Host,
		cfg.Database.Port,
		cfg.Database.Name,
		cfg.Database.SSLMode,
	)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	conn, err := pgx.Connect(ctx, dsn)
	if err != nil {
		return nil, err
	}

	if err := conn.Ping(ctx); err != nil {
		return nil, err
	}

	return conn, nil
}

func RunMigrations(conn *pgx.Conn, migrationsDir string) error {
	goose.SetBaseFS(nil)
	goose.SetDialect("postgres")

	// хз че он ругается, нужна помощь

// 	db := &GooseDBAdapter{conn: conn}

// 	if err := goose.Up(db, migrationsDir); err != nil {
// 		return err
// 	}

// 	return nil
// }

type GooseDBAdapter struct {
	conn *pgx.Conn
}

func (pg *GooseDBAdapter) Exec(query string, args ...interface{}) (goose.Result, error) {
	res, err := pg.conn.Exec(context.Background(), query, args...)
	return GooseResult{res}, err
}

func (pg *GooseDBAdapter) Query(query string, args ...interface{}) (goose.Rows, error) {
	rows, err := pg.conn.Query(context.Background(), query, args...)
	return GooseRows{rows}, err
}

type GooseResult struct {
	pgx.CommandTag
}

func (r GooseResult) LastInsertId() (int64, error) {
	return 0, fmt.Errorf("LastInsertId not supported by pgx")
}

func (r GooseResult) RowsAffected() (int64, error) {
	return int64(r.CommandTag.RowsAffected()), nil
}

type GooseRows struct {
	pgx.Rows
}

func (r GooseRows) Columns() ([]string, error) {
	fields := r.Rows.FieldDescriptions()
	columns := make([]string, len(fields))
	for i, field := range fields {
		columns[i] = field.Name
	}
	return columns, nil
}
