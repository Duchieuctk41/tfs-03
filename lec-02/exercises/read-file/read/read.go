package read
import (
	"fmt"
	"io/ioutil"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func ReadFile(fileName string) {
	file, err := ioutil.ReadFile(fileName)
	check(err)
	fmt.Print(string(file))
}

func WrFile(fileName string, str string) {
	myData := []byte(str)

	err := ioutil.WriteFile(fileName, myData, 0777)

	//handle this error
	if err != nil {
	// print it out
	fmt.Println(err)
	}
}
