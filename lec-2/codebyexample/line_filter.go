package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	// user input then output in cmd (like cin >>, cout << in C++)
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter text: ")
	text, _ := reader.ReadString('\n')
	fmt.Println(text)

	// user input then toUpper string
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print("Enter text: ")
	for scanner.Scan() {
		ucl := strings.ToUpper(scanner.Text())
		fmt.Println(ucl)
		return
	}
	if err := scanner.Err(); err != nil {
		fmt.Fprintf(os.Stderr, "error", err)
		os.Exit(1)
	}
	defer func() {
		fmt.Println("Stopped")
	}()
}