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
	Name   string
	Gender string
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
				if node.key == k {
					fmt.Println("existed")
					break
				} else {
					newNode := &bucketNode{key: k, value: v}
					newNode.next = b.head
					b.head = newNode
					break
				}
			}
		}
		node = node.next
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

func hash(k string) int {
	sum := 0
	for _, v := range k {
		sum += int(v)
	}
	return sum % ArraySize
}

func NewHashTable() *HashTable {
	result := &HashTable{}
	for i := range result.array {
		result.array[i] = &bucket{}
	}
	return result
}

func main() {
	hashTable := NewHashTable()

	user := User{
		Name:   "hieu",
		Gender: "femalemeo",
	}

	hashTable.Insert(user.Name, user)
	hashTable.Insert(user.Name, user)
	fmt.Println("looking for hieu: ", hashTable.Search("hieu"))
	fmt.Println("looking for notfound: ", hashTable.Search("notfound"))
}
