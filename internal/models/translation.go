package models

import (
	"context"
	"errors"

	"github.com/jackc/pgx/v5"
)

type Translation struct {
	TextID int    `json:"text_id"`
	Ru     string `json:"ru"`
	En     string `json:"en"`
}

func GetTranslationByID(ctx context.Context, conn *pgx.Conn, textID int) (*Translation, error) {
	var translation Translation
	err := conn.QueryRow(ctx, `
		SELECT text_id, ru, en
		FROM translations
		WHERE text_id = $1
	`, textID).Scan(
		&translation.TextID, &translation.Ru, &translation.En,
	)

	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, nil
		}
		return nil, err
	}

	return &translation, nil
}
