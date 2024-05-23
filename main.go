package main

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgogres"
	password = "postgogres"
	dbname   = "postgogres"
)

func main() {
	// connection string
	psqlconn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)

	// open database
	db, err := sql.Open("postgres", psqlconn)
	CheckError(err)

	// close database
	defer db.Close()

	// check db
	err = db.Ping()
	CheckError(err)

	initialize := true
	if initialize {
		SetupDatabase(db)
		InsertData(db)
	}

	QueryData(db)
	UpdateData(db)
	QueryData(db)
	DeleteData(db)
	QueryData(db)

	fmt.Println("Connected!")
}

func CheckError(err error) {
	if err != nil {
		panic(err)
	}
}

func UpdateData(db *sql.DB) {
	// update
	updateStmt := `update "Students" set "Name" = $1 where "Roll_Number" = $2`
	_, err := db.Exec(updateStmt, "Jacob", 20)
	CheckError(err)
}

func DeleteData(db *sql.DB) {
	// delete
	deleteStmt := `delete from "Students" where "Roll_Number" = $1`
	_, err := db.Exec(deleteStmt, 20)
	CheckError(err)
}

func QueryData(db *sql.DB) {
	// query
	rows, err := db.Query(`select * from "Students"`)
	CheckError(err)

	defer rows.Close()

	// iterate
	for rows.Next() {
		var id, roll int
		var name string
		err = rows.Scan(&id, &name, &roll)
		CheckError(err)
		fmt.Println(id, name, roll)
	}

}

func SetupDatabase(db *sql.DB) {
	createTableStmt := `
	CREATE TABLE IF NOT EXISTS "Students" (
		"ID" SERIAL PRIMARY KEY,
		"Name" VARCHAR(255) NOT NULL,
		"Roll_Number" INT NOT NULL
	)
	`
	_, err := db.Exec(createTableStmt)
	CheckError(err)
}

func InsertData(db *sql.DB) {
	// insert
	// hardcoded
	insertStmt := `insert into "Students"("Name", "Roll_Number") values('Jacob', 20)`
	_, err := db.Exec(insertStmt)
	CheckError(err)

	// dynamic
	insertDynStmt := `insert into "Students"("Name", "Roll_Number") values($1, $2)`
	_, err = db.Exec(insertDynStmt, "Jack", 21)
	CheckError(err)
}
