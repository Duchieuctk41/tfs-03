package main

import "fmt"

const ArraySize = 7

// Hashtable structure
type HashTable struct {
	array [ArraySize]*bucket
}

// bucket structure
type bucket struct {
	head *bucketNode
}

type User struct {
	Email   string
	Name    string
	Address string
}

// bucketNode structure
type bucketNode struct {
	key   string
	value User
	next  *bucketNode
}

// Insert will take in a key and it to the hash table array
func (h *HashTable) Insert(key string, value User) {
	index := hash(key)
	h.array[index].insert(key, value)
}

// // Search will take in a key and return true if that key is stored in the hash table
func (h *HashTable) Search(key string) interface{} {
	index := hash(key)
	return h.array[index].search(key)
}

// // Delete will take in a key and delete it from the hash table
func (h *HashTable) Delete(key string) {
	index := hash(key)
	h.array[index].delete(key)
}

// insert will take in a key, create a node with the key and insert the node in the bucket
func (b *bucket) insert(k string, v User) {
	if b.search(k) == "not found" {
		newNode := &bucketNode{key: k, value: v}
		newNode.next = b.head
		b.head = newNode
	} else {
		fmt.Println(k, "already exist")
	}
}

// search will take in a key, return true if the bucket has that key
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

// delete will take in a key and delete the node from the bucket
func (b *bucket) delete(k string) {

	if b.head.key == k {
		b.head = b.head.next
		return
	}

	previousNode := b.head
	for previousNode.next != nil {
		if previousNode.next.key == k {
			// delete
			previousNode.next = previousNode.next.next
		}
		previousNode = previousNode.next
	}
}

// hash
func hash(key string) int {
	sum := 0
	for _, v := range key {
		sum += int(v)
	}
	return sum % ArraySize
}

// Init will create a bucket in each slot of the hash table
func Init() *HashTable {
	result := &HashTable{}
	for i := range result.array {
		result.array[i] = &bucket{}
	}
	return result
}

func main() {
	hashTable := Init()
	userName := User{
		Email:   "a@gmail.com",
		Name:    "Hieu",
		Address: "Da Lat",
	}
	hashTable.Insert(userName.Email, userName)
	fmt.Println("a@gmail.com", hashTable.Search("a@gmail.com"))
	fmt.Println("notfound@gmail.com", hashTable.Search("notfound@gmail.com"))
}
