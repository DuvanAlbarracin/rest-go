package main

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

type Book struct{
    id string
    name string
    author string
}

func dbOperations(db *sql.DB){
    stmt, errstm := db.Prepare("INSERT INTO books (name, author, isbn) VALUES (?, ?, ?)")
    if errstm != nil{
        log.Fatal("Error inserting books into database...", errstm)
    }
    defer stmt.Close()
    stmt.Exec("A Tale of Two Cities", "Charles Dickens", 140430547)
    log.Println("Inserted book into database")

}

func main(){
    db, errdb := sql.Open("sqlite3", "/books.db")
    if errdb != nil{
        panic(errdb)
    }

    statemnt, errstm := db.Prepare("CREATE TABLE IF NOT EXISTS books (id INTEGER PRIMARY KEY, isbn INTEGER, author VARCHAR(64), name VARCHAR(64) NULL))")
    if errstm != nil{
        log.Println("Error creating the database!")
    } else {
        log.Println("Database created successfully!")
    }
    defer statemnt.Close()

    statemnt.Exec()
    dbOperations(db)
}
