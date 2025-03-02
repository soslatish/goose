package database

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5"
)

// проверяет, существуют ли в базе данных все ожидаемые таблицы
func VerifySchema(ctx context.Context, conn *pgx.Conn) error {
	expectedTables := []string{
		"cards",
		"user_card_relation",
		"descriptions",
		"traits",
		"melee",
		"translations",
	}

	for _, tableName := range expectedTables {
		var exists bool
		err := conn.QueryRow(ctx, `
			SELECT EXISTS (
				SELECT FROM information_schema.tables 
				WHERE table_schema = 'public' 
				AND table_name = $1
			)
		`, tableName).Scan(&exists)

		if err != nil {
			return fmt.Errorf("error checking if table %s exists: %w", tableName, err)
		}

		if !exists {
			return fmt.Errorf("table %s does not exist in the database", tableName)
		}
	}

	fmt.Println("Schema verification successful. All expected tables exist.")
	return nil
}
