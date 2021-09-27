package functional

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

type User struct {
	ID    int    `json:"id" gorm:"PrimaryKey"`
	Email string `json:"email"`
	Name  string `json:"name"`
}

func main() {
	// connect database
	db, err := gorm.Open(mysql.Open("root:password@tcp(localhost:3306)/meo?parseTime=true"), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	DB = db
	db.AutoMigrate(&User{})
	initRouter()
}

// Tìm kiếm user
func FindUser(w http.ResponseWriter, r *http.Request) {
	email := mux.Vars(r)["email"]
	var user User

	DB.Where("email = ?", email).First(&user)
	fmt.Println(user)
}

// init router
func initRouter() {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/user/{email}", FindUser).Methods("GET")

	fmt.Println("server is running localhost:3000")
	http.ListenAndServe(":3000", router)
}
