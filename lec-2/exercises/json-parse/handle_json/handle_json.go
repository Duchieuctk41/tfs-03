package handle_json

import (
	"encoding/json"
	"fmt"
	"strconv"
)

type user struct {
	Name  string `json:"name"`
	Age   int    `json:"age"`
	Lover string `json:"lover"`
}

func ReadFileJson() {
	strJson := `{"Name": "Hieu", "Age": 22, "Lover": "Ngoc Trinh"}`

	u := user{}
	json.Unmarshal([]byte(strJson), &u)
	objJson := map[string]string{
		"Name":  u.Name,
		"Age":   strconv.Itoa(u.Age),
		"Lover": u.Lover,
	}
	fmt.Println(objJson)
	fmt.Println(u)
	fmt.Print("Name: " + u.Name)
	fmt.Print(", Age: " + strconv.Itoa(u.Age))
	fmt.Print(", Lover: " + u.Lover)
}
