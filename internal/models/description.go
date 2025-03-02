package models

import (
	"context"
	"encoding/json"
	"errors"

	"github.com/jackc/pgx/v5"
)

type Description struct {
	ID            string          `json:"id"`
	NameID        int             `json:"name_id"`
	DescriptionID int             `json:"description_id"`
	BaseParams    json.RawMessage `json:"base_params"`
	UpgradeStep   json.RawMessage `json:"upgrade_step"`
}

func GetDescriptionByID(ctx context.Context, conn *pgx.Conn, descID string) (*Description, error) {
	var desc Description
	err := conn.QueryRow(ctx, `
		SELECT id, name_id, description_id, base_params, upgrade_step
		FROM descriptions
		WHERE id = $1
	`, descID).Scan(
		&desc.ID, &desc.NameID, &desc.DescriptionID, &desc.BaseParams, &desc.UpgradeStep,
	)

	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, nil
		}
		return nil, err
	}

	return &desc, nil
}
