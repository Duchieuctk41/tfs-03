package main

import "fmt"

const ArraySize = 10

type HashTable struct {
	array [ArraySize]*bucket
}

type bucket struct {
	head *bucketNode
}

type bucketNode struct {
	key   string
	value User
	next  *bucketNode
}

type User struct {
	Email   string
	Name    string
	Address string
}

type ListUser []User

func (h *HashTable) Insert(key string, value User) {
	index := hash(key)
	h.array[index].insert(key, value)
}

func (b *bucket) insert(k string, v User) {
	if b.search(k) == "not found" {
		newNode := &bucketNode{key: k, value: v}
		newNode.next = b.head
		b.head = newNode
	} else {
		node := b.head
		for {
			if node != nil {
				d := node.value
				if d.Email == k {
					fmt.Println(k, " existed!")
					break
				} else {
					newNode := &bucketNode{key: k, value: v}
					newNode.next = b.head
					b.head = newNode
					break
				}
			}
			node = node.next
		}
	}
}

func (h *HashTable) Search(key string) interface{} {
	index := hash(key)
	return h.array[index].search(key)
}

func (b *bucket) search(k string) interface{} {
	currentNode := b.head
	for currentNode != nil {
		if currentNode.key == k {
			return currentNode.value
		}
		currentNode = currentNode.next
	}
	return "not found"
}

func Create() *HashTable {
	result := &HashTable{}
	for i := range result.array {
		result.array[i] = &bucket{}
	}
	return result
}

func hash(k string) int {
	sum := 0
	for _, v := range k {
		sum += int(v)
	}
	return sum % ArraySize
}

func main() {
	hashTable := Create()
	user1 := User{
		Email:   "Hieu@gmail.com",
		Name:    "Hieu",
		Address: "Dalat",
	}
	user2 := User{
		Email:   "Hieu@gmail.com",
		Name:    "Hieu",
		Address: "Dalat1",
	}

	listUser := ListUser{user1, user2}

	for _, v := range listUser {
		hashTable.Insert(v.Email, v)
	}

	fmt.Println("Looking for Hieu@gmail.com: ", hashTable.Search("Hieu@gmail.com"))
	fmt.Println("Looking for notfound@gmail.com: ", hashTable.Search("notfound@gmail.com"))
}
