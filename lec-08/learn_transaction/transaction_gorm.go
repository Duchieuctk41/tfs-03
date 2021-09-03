package main

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var db *gorm.DB

type Sender struct {
	UserID string
	Total  float32
}

type Receiver struct {
	UserID string
	Total  float32
}

func main() {
	var err error
	db, err = gorm.Open(mysql.Open("root:password@tcp(localhost:3306)/paypay?parseTime=true"), &gorm.Config{})

	if err != nil {
		panic(err.Error())
	}

	sqlDB, err := db.DB()
	defer sqlDB.Close()
	transfer("user_A", "user_B", 100)
}

func transfer(userA string, userB string, amount float32) {
	db.Transaction(func(tx *gorm.DB) error {

		var (
			user1 Sender
			user2 Receiver
		)

		if err := tx.Where("user_id = ?", userA).First(&user1).Error; err != nil {
			fmt.Println("canot find user1")
			return err
		}
		if err := tx.Where("user_id = ?", userB).First(&user2).Error; err != nil {
			fmt.Println("canot find user2")
			return err
		}

		user1.Total = user1.Total - amount
		user2.Total = user2.Total + amount

		if err := tx.Model(&user1).Where("user_id = ?", user1.UserID).Update("total", user1.Total).Error; err != nil {
			fmt.Println("canot update user1")
			return err
		}

		if err := tx.Model(&user2).Where("user_id = ?", user2.UserID).Update("total", user2.Total).Error; err != nil {
			fmt.Println("canot update user2")
			return err
		}

		fmt.Println("attempting to set total_receiver: ", user2.Total, " total_sender: ", user1.Total)

		return nil
	})

}
