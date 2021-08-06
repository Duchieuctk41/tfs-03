package main

import (
	"fmt"
)

type Animal interface {
	Speak()
}

type Action interface {
	HandShake()
}


type Communication interface {
	SayHi() string
}

// Embed Interface

type Human interface {
	Action
	Communication
}

type Dog struct {}
type Person struct {
	Name string
}

func (dog Dog) Speak() {
	fmt.Println("go go go")
}

func (person Person) HandShake() {
	fmt.Println("wave wave wave")
}

func (person *Person) SayHi() string {
	return "Hi everyone, my name is " + person.Name
}

func PrintSayHi(person Human) {
	fmt.Println(person.SayHi())
}

// Empty interface

func PrintEveryThing(everyThing interface{}) {
	fmt.Println(everyThing)
}

// Empty interface, giống khai báo void trong C++

func ReturnEveryThing(s string) interface{} {
	if s == "hieu thu bar" {
		return "hi " + s
	}
	return true
}

func main() {
	var chiHuaHua Animal = Dog{}
	chiHuaHua.Speak()

	var hieuThu2 Human =  &Person { Name: "Hieu thu high" }
	hieuThu2.HandShake()
	PrintSayHi(hieuThu2)

	PrintEveryThing(10)
	fmt.Println(ReturnEveryThing("hieu thu bar"))
}
