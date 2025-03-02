package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"path/filepath"

	"goose/internal/config"
	"goose/internal/database"
)

func main() {
	cfg, err := config.Load()
	if err != nil {
		log.Fatalf("Failed to load configuration: %v", err)
	}
	conn, err := database.ConnectDB(cfg)
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	defer conn.Close(context.Background())

	execPath, err := os.Executable()
	if err != nil {
		log.Fatalf("Failed to get  path: %v", err)
	}

	migrationsDir := filepath.Join(filepath.Dir(execPath), "..", "migrations")

	if err := database.RunMigrations(conn, migrationsDir); err != nil {
		log.Fatalf("Failed to run migration: %v", err)
	}

	fmt.Println("migration succeess")

	// get user card and get card

	// card, err := GetCardByID(ctx, conn, 17)
	// if err != nil {
	// 	log.Fatalf("Failed to get card: %v", err)
	// }
	// if card != nil {
	// 	fmt.Printf("Found card: %s (ID: %d)\n", card.EnText, card.ID)
	// }

	// userCards, err := GetUserCards(ctx, conn, 3)
	// if err != nil {
	// 	log.Fatalf("Failed to get user cards: %v", err)
	// }
	// fmt.Printf("User has %d cards\n", len(userCards))
}
