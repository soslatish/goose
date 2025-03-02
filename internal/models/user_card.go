package models

import (
	"context"

	"github.com/jackc/pgx/v5"
)

type UserCard struct {
	UserID       int `json:"user_id"`
	CardID       int `json:"card_id"`
	UpgradeLevel int `json:"upgrade_level"`
}

func GetUserCards(ctx context.Context, conn *pgx.Conn, userID int) ([]UserCard, error) {
	rows, err := conn.Query(ctx, `
		SELECT user_id, card_id, upgrade_level
		FROM user_card_relation
		WHERE user_id = $1
	`, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var userCards []UserCard
	for rows.Next() {
		var uc UserCard
		if err := rows.Scan(&uc.UserID, &uc.CardID, &uc.UpgradeLevel); err != nil {
			return nil, err
		}
		userCards = append(userCards, uc)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return userCards, nil
}
