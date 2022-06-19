package main

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

type Book struct {
	id     int
	name   string
	author string
}

func dbOperations(db *sql.DB) {
	stmt, errstm := db.Prepare("INSERT INTO books (name, author, isbn) VALUES (?, ?, ?)")
	if errstm != nil {
		log.Fatal("Error inserting books into database...", errstm)
	}
	defer stmt.Close()
	stmt.Exec("A Tale of Two Cities", "Charles Dickens", 140430547)
	log.Println("Inserted book into database")

	rows, errows := db.Query("SELECT id, name, author FROM Books")
	if errows != nil {
		log.Fatal("Error quering books to the database...", errows)
	}
	defer rows.Close()
	var currBook Book
	for rows.Next() {
		if errScan := rows.Scan(&currBook.id, &currBook.name, &currBook.author); errScan != nil {
			log.Fatal("Error scanning row...")
		}
		log.Printf("ID: %d, Book: %s, Author: %s\n", currBook.id, currBook.name, currBook.author)
	}

	stmt, errstm = db.Prepare("UPDATE books SET name=? WHERE id=?")
	if errstm != nil {
		log.Fatal("Error updating books ...", errstm)
	}
	defer stmt.Close()
	stmt.Exec("The Tale of Two Cities", 1)
	log.Println("Updated book in database")

	stmt, errstm = db.Prepare("DELETE FROM books WHERE id=?")
	if errstm != nil {
		log.Fatal("Error deleting books ...", errstm)
	}
	defer stmt.Close()
	stmt.Exec(1)
	log.Println("Deleted book from database")
}

func main() {
	db, errdb := sql.Open("sqlite3", "./books.db")
	if errdb != nil {
        panic(errdb)
	}

	statemnt, errstm := db.Prepare("CREATE TABLE IF NOT EXISTS books (id INTEGER PRIMARY KEY, isbn INTEGER, author VARCHAR(64), name VARCHAR(64) NULL)")
	if errstm != nil {
        log.Fatal("Error creating the database:", errstm)
	} else {
		log.Println("Database created successfully!")
	}
	defer statemnt.Close()

	statemnt.Exec()
	dbOperations(db)
}
