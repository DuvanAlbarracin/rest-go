package dbutils

import (
	"database/sql"
	"log"
)

func createTable(dbDriver *sql.DB, query string) error{
    stmt, errstmt := dbDriver.Prepare(query)
    if errstmt != nil{
        return errstmt
    }
    defer stmt.Close()
    _, errExec := stmt.Exec()
    if errExec != nil{
        return errExec
    }
    return nil
}

func createTables(dbDriver *sql.DB, queries ...string) error{
    for _, query := range queries{
        stmt, errstmt := dbDriver.Prepare(query)
        if errstmt != nil{
            return errstmt
        }
        defer stmt.Close()
        _, errExec := stmt.Exec()
        if errExec != nil{
            return errExec
        }
    }
    return nil
}

func Initialize(dbDriver *sql.DB){
    //createTables(dbDriver, train, station, schedule)
    err := createTable(dbDriver, train)
    if err != nil{
        log.Fatal("Error creating table...", err)
    }

    err = createTable(dbDriver, station)
    if err != nil{
        log.Fatal("Error creating table...", err)
    }

    err = createTable(dbDriver, schedule)
    if err != nil{
        log.Fatal("Error creating table...", err)
    }

    log.Println("All tables created successfully!")
}
