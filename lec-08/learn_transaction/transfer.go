package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

var (
	ctx context.Context
	db  *sql.DB
)

func main() {
	var err error
	db, err = sql.Open("mysql", "root:password@tcp(localhost:3306)/paypay")
	if err != nil {
		panic(err.Error())
	}

	defer db.Close()

	transfer("user_A", "user_B", 333)
}

func transfer(userA string, userB string, amount float32) {
	// func (db *DB) Begin() (*Tx, error)
	tx, err := db.Begin()
	if err != nil {
		log.Fatal(err)
	}

	var (
		total_send    float32
		total_receive float32
	)

	// func (tx *Tx) QueryRow(query string, args ... interface{}) * Row
	row := tx.QueryRow(`SELECT total FROM send WHERE user_id = ?`, userA)
	err = row.Scan(&total_send)
	if err != nil {
		// func (tx *Tx) Rollback() error
		_ = tx.Rollback()
		fmt.Println("err: ", err)
		return
	}

	row = tx.QueryRow(`SELECT total FROM receive WHERE user_id = ?`, userB)
	err = row.Scan(&total_receive)
	if err != nil {
		// func (tx *Tx) Rollback() error
		_ = tx.Rollback()
		fmt.Println("err: ", err)
		return
	}

	total_send = total_send - amount
	total_receive = total_receive + amount

	fmt.Println("attempting to set total_receive: ", total_receive, " total_send: ", total_send)

	var result sql.Result
	// func (tx *Tx) Exec(query string, args ...interface{}) (Result, error)
	result, execErr := tx.Exec(`UPDATE send SET total = ? WHERE user_id = ?`, total_send, userA)
	rowsAffected, _ := result.RowsAffected()

	fmt.Println("update send execErr:", execErr, "rowsAffected:", rowsAffected)

	if execErr != nil || rowsAffected != 1 {
		//  func (tx *Tx) Rollback() error
		_ = tx.Rollback()
		return
	}

	// func (tx *Tx) Exec(query string, args ...interface{}) (Result, error)
	result, execErr = tx.Exec(`UPDATE receive SET total = ? WHERE user_id = ?`, total_receive, userB)
	rowsAffected, _ = result.RowsAffected()

	fmt.Println("update receive execErr:", execErr, "rowsAffected:", rowsAffected)

	if execErr != nil || rowsAffected != 1 {
		// func (tx *Tx) Rollback() error
		_ = tx.Rollback()
		return
	}

	// func (tx *Tx) Commit() error
	if err := tx.Commit(); err != nil {
		fmt.Println(err)
		return
	}
}
