package utils

import (
	"database/sql"
	"fmt"
)

func SeedDB(db *sql.DB) {
	// Drop all tables
	_, err := db.Exec(`DROP TABLE IF EXISTS events, users, categories`)
	if err != nil {
		fmt.Println(err)
	}

	// Create users table
	_, err = db.Exec(`CREATE TABLE users (
		id SERIAL PRIMARY KEY,
		username VARCHAR(50) UNIQUE NOT NULL,
		password VARCHAR(255) NOT NULL,
		is_organiser BOOLEAN DEFAULT FALSE,
		is_admin BOOLEAN DEFAULT FALSE,
		created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
	)`)
	if err != nil {
		fmt.Println(err)
	}

	// Create Categories table
	_, err = db.Exec(`CREATE TABLE categories (
		id SERIAL PRIMARY KEY,
		name VARCHAR(100) NOT NULL
	)`)
	if err != nil {
		fmt.Println(err)
	}

	// Seed categories table
	_, err = db.Exec(`INSERT INTO categories (name) VALUES ('Music'), ('Food'), ('Art'), ('Sports'), ('Tech'), ('Business')`)
	if err != nil {
		fmt.Println(err)
	}

	// Create events table
	_, err = db.Exec(`CREATE TABLE events (
		id SERIAL PRIMARY KEY,
		title VARCHAR(100) NOT NULL,
		description TEXT NOT NULL,
		category_id INT,
		FOREIGN KEY (category_id) REFERENCES categories(id),
		user_id INT,
		FOREIGN KEY (user_id) REFERENCES users(id)
	)`)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("Successfully seeded database!")
}