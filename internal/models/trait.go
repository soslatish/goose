package models

import (
	"context"
	"encoding/json"
	"errors"

	"github.com/jackc/pgx/v5"
)

type Trait struct {
	ID            string          `json:"id"`
	NameID        int             `json:"name_id"`
	DescriptionID int             `json:"description_id"`
	Params        json.RawMessage `json:"params"`
}

func GetTraitByID(ctx context.Context, conn *pgx.Conn, traitID string) (*Trait, error) {
	var trait Trait
	err := conn.QueryRow(ctx, `
		SELECT id, name_id, description_id, params
		FROM traits
		WHERE id = $1
	`, traitID).Scan(
		&trait.ID, &trait.NameID, &trait.DescriptionID, &trait.Params,
	)

	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, nil
		}
		return nil, err
	}

	return &trait, nil
}
