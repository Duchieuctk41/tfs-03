package inter

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

type UserCase interface {
	Fetch(email string) *User
}

type User struct {
	ID    int    `json:"id" gorm:"PrimaryKey"`
	Email string `json:"email"`
	Name  string `json:"name"`
}

// fetch email
func (user *User) Fetch(email string) *User {
	DB.Where("email = ?", email).First(&user)

	return user
}

func main() {
	db, err := gorm.Open(mysql.Open("root:password@tcp(localhost:3306)/meo?parseTime=true"), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	DB = db
	db.AutoMigrate(&User{})
	initRouter()
}

// init router
func initRouter() {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/user/{email}", FindUser).Methods("GET")

	fmt.Println("server is running localhost:3000")
	http.ListenAndServe(":3000", router)
}

// find user by email
func FindUser(w http.ResponseWriter, r *http.Request) {
	email := mux.Vars(r)["email"]
	var user User

	var userCase UserCase
	userCase = &user

	userCase.Fetch(email)
}
