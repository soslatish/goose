package models

import (
	"context"
	"encoding/json"
	"errors"

	"github.com/jackc/pgx/v5"
)

type Melee struct {
	ID            string          `json:"id"`
	NameID        int             `json:"name_id"`
	DescriptionID int             `json:"description_id"`
	Params        json.RawMessage `json:"params"`
	UpgradeStep   json.RawMessage `json:"upgrade_step"`
}

func GetMeleeByID(ctx context.Context, conn *pgx.Conn, meleeID string) (*Melee, error) {
	var melee Melee
	err := conn.QueryRow(ctx, `
		SELECT id, name_id, description_id, params, upgrade_step
		FROM melee
		WHERE id = $1
	`, meleeID).Scan(
		&melee.ID, &melee.NameID, &melee.DescriptionID, &melee.Params, &melee.UpgradeStep,
	)

	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, nil
		}
		return nil, err
	}

	return &melee, nil
}
