package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

func main() {
	//параметры коннектра
	connStr := "postgresql://postgres:secret@localhost:5432/postgres?sslmode-disabled"

	//открытие коннекта
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	//проверка
	err = db.Ping()
	if err != nil {
		log.Fatal("Ошибка подключения к бд:", err)
	}
	fmt.Println("прокнуло, коннект есть")

	//создание бдшки тимс
	rows, err := db.Query("SELECT id, names, city, coach_id FROM team ")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	fmt.Println("BD TEAM:")
	for rows.Next() {
		var id, coach_id int
		var names, city string
		if err := rows.Scan(&id, &names, &city, &coach_id); err != nil {
			log.Fatal(err)
		}
		fmt.Println(id, names, city, coach_id)
	}
	if err := rows.Err(); err != nil {
		log.Fatal(err)
	}
}
