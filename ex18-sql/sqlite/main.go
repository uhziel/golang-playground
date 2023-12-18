package main

import (
	"context"
	"database/sql"
	"log"
	"os"

	_ "modernc.org/sqlite"
)

var createUserTableStmt = `
CREATE TABLE user (
    id bigint AUTO_INCREMENT,
    name varchar(255) NOT NULL DEFAULT '',
    password varchar(255) NOT NULL DEFAULT '',
    PRIMARY KEY (id)
)
`

var selectUserStmt = "SELECT name, password FROM user WHERE name=$1"

func main() {
	os.Remove("./user.db")

	ctx := context.Background()

	db, err := sql.Open("sqlite", "./user.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	if err := db.PingContext(ctx); err != nil {
		log.Fatal(err)
	}

	// create table: user
	if _, err := db.ExecContext(ctx, createUserTableStmt); err != nil {
		log.Fatal(err)
	}

	// insert
	result, err := db.ExecContext(ctx, "INSERT INTO user(name, password) VALUES($1, $2)", "zhulei", "123456")
	if err != nil {
		log.Fatal(err)
	}
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		log.Fatal(err)
	}
	lastInsertId, err := result.LastInsertId()
	if err != nil {
		log.Fatal(err)
	}
	log.Println("insert", "RowsAffected:", rowsAffected, "LastInsertId:", lastInsertId)

	// query
	rows, err := db.QueryContext(ctx, selectUserStmt, "zhulei")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	var user User
	for rows.Next() {
		rows.Scan(&user.Name, &user.Password)
		log.Println(user)
	}

	// statement
	stmt, err := db.PrepareContext(ctx, selectUserStmt)
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()
	var user2 User
	if err := stmt.QueryRowContext(ctx, "zhulei").Scan(&user2.Name, &user2.Password); err != nil {
		log.Fatal(err)
	}
	log.Println("user2", user2)

	// transaction
	tx, err := db.BeginTx(ctx, nil)
	if err != nil {
		log.Fatal(err)
	}
	result, err = tx.ExecContext(ctx, "UPDATE user SET password=$2 WHERE name=$1", "zhulei", "abcd")
	if err != nil {
		tx.Rollback()
		log.Fatal(err)
	}
	var user3 User
	if err := tx.QueryRowContext(ctx, selectUserStmt, "zhulei").Scan(&user3.Name, &user3.Password); err != nil {
		tx.Rollback()
		log.Fatal(err)
	}
	log.Println("user3", user3)

	if err := tx.Commit(); err != nil {
		log.Fatal(err)
	}

	log.Println("end!")
}

type User struct {
	Name     string
	Password string
}
