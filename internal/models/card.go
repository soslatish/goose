package models

import (
	"context"
	"errors"

	"github.com/jackc/pgx/v5"
)

type Card struct {
	ID         int    `json:"card_id"`
	Name       int    `json:"name"`
	Rareness   string `json:"rareness"`
	Class      string `json:"class"`
	TraitSkill string `json:"trait_skill"`
	MajorSkill string `json:"major_skill"`
	NFT        string `json:"nft"`
}

func GetCardByID(ctx context.Context, conn *pgx.Conn, cardID int) (*Card, error) {
	var card Card
	err := conn.QueryRow(ctx, `
		SELECT card_id, name, rareness, class, trait_skill, major_skill, nft, text_id, ru, en
		FROM cards
		WHERE card_id = $1
	`, cardID).Scan(
		&card.ID, &card.Name, &card.Rareness, &card.Class, &card.TraitSkill,
		&card.MajorSkill, &card.NFT,
	)

	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, nil
		}
		return nil, err
	}

	return &card, nil
}
