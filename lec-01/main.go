package main

import "fmt"

func main() {
	str := "hieu"

	sum := 0
	for _, v := range str {
		sum += int(v)
	}
	sum = sum % 10

	fmt.Println(sum)
}
