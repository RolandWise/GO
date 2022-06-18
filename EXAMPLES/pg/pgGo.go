package main

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "rusho1234"
	dbname   = "MyDB"
)

func main() {
	psqlconn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)

	db, err := sql.Open("postgres", psqlconn)
	CheckError(err)

	defer db.Close()

	// insert
	// hardcoded
	insertStmt := `insert into "Students"("Name", "Roll") values('John', 1)`
	_, e := db.Exec(insertStmt)
	CheckError(e)

	// dynamic
	insertDynStmt := `insert into "Students"("Name", "Roll") values($1, $2)`
	_, e = db.Exec(insertDynStmt, "Jane", 2)
	CheckError(e)

	// update
	updateStmt := `update "Students" set "Name"=$1, "Roll"=$2 where "id"=$3`
	_, e = db.Exec(updateStmt, "Mary", 3, 2)
	CheckError(e)

	// Delete
	deleteStmt := `delete from "Students" where id=$1`
	_, e = db.Exec(deleteStmt, 1)
	CheckError(e)

}

func CheckError(err error) {
	if err != nil {
		panic(err)
	}
}
