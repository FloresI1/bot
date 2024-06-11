package config

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

const (
	tokenKey  = "token"
	chatIDKey = "chat_id"
)

// InitDB инициализирует базу данных и создает таблицу, если ее нет
func InitDB(filepath string) *sql.DB {
	db, err := sql.Open("sqlite3", filepath)
	if err != nil {
		log.Fatal(err)
	}

	createTable(db)
	return db
}

// createTable создает таблицу конфигурации
func createTable(db *sql.DB) {
	query := `CREATE TABLE IF NOT EXISTS config (
		key TEXT PRIMARY KEY,
		value TEXT
	);`
	_, err := db.Exec(query)
	if err != nil {
		log.Fatal(err)
	}
}

// LoadTokenAndChatID загружает токен и Chat ID из базы данных
func LoadTokenAndChatID(db *sql.DB) (string, string) {
	var token, chatID string

	rows, err := db.Query("SELECT key, value FROM config")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	for rows.Next() {
		var key, value string
		err = rows.Scan(&key, &value)
		if err != nil {
			log.Fatal(err)
		}

		if key == tokenKey {
			token = value
		} else if key == chatIDKey {
			chatID = value
		}
	}

	return token, chatID
}

// SaveConfig сохраняет токен и Chat ID в базу данных
func SaveConfig(db *sql.DB, key, value string) {
	_, err := db.Exec("INSERT OR REPLACE INTO config (key, value) VALUES (?, ?)", key, value)
	if err != nil {
		log.Fatal(err)
	}
}
