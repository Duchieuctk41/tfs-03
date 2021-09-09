package main

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

// check err
func failOnError(err error) {
	if err != nil {
		log.Println(err)
	}
}

// gửi tin nhắn đến email vừa mua hàng
func sendMail(db *sql.DB, done chan interface{}, email <-chan string, id chan int) {
	for {
		select {
		case mail := <-email:
			fmt.Printf("sending email to: %v", mail) // send tin nhắn cho email này
			sent(db, <-id)                           // update trạng thái đã gửi email
		case <-done:
			return
		}
	}
}

// tìm các email chưa đc gửi tin nhắn
func checkMail(db *sql.DB, unSentEmail chan<- string, ids chan int) {
	log.Println("check")
	rows, err := db.Query(`SELECT email, id FROM mail WHERE sent = false`) // tìm email chưa đc gửi
	failOnError(err)
	defer rows.Close()

	for rows.Next() {
		var email string
		var id int
		err := rows.Scan(&email, &id)
		failOnError(err)
		unSentEmail <- email
		ids <- id
	}
}

// update trạng thái đã send cho email
func sent(db *sql.DB, id int) {
	_, err := db.Exec(`UPDATE mail SET sent = true WHERE id = ?`, id)
	failOnError(err)
}

func main() {
	id := make(chan int)
	unSentEmail := make(chan string)
	done := make(chan interface{})

	// connect db
	db, err := sql.Open("mysql", "root:password@tcp(localhost:3306)/meow")
	failOnError(err)

	// sent mail
	go sendMail(db, done, unSentEmail, id)

	ticker := time.NewTicker(time.Minute) // 1 phút check 1 lần
	defer ticker.Stop()

	for {
		select {
		case <-done:
			fmt.Println("Done!") // sẽ chả bh chạy đến cái này
			return
		case <-ticker.C:
			checkMail(db, unSentEmail, id)
		}
	}
}
