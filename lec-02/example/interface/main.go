package main

import "fmt"

type Speaker interface {
	Speak() string
}

type Cat struct {
	Say string
}

type Human struct {
	Greet string
}

func (cat Cat) Speak() string {
	return cat.Say
}

func (human Human) Speak() string {
	return human.Greet
}

func main() {
	var sp Speaker

	cat := Cat{"mèo méo meo mèo meo"}
	sp = cat
	fmt.Println(sp.Speak())

	human := Human{"Hello World"}
	sp = human
	fmt.Println(sp.Speak())
}
