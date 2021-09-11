package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	// NewReader trả về  bộ đệm mới với kích thước mặc định là 4K 
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter text: ")
	text, _ := reader.ReadString('\n')
	fmt.Println(text)

	// NewScanner đọc từ bộ đệm có sẵn, giá trị giới hạn dòng 64k (4096 byte)
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print("Enter text: ")
	for scanner.Scan() {
		fmt.Println(scanner.Text())
		return
	}

	defer func() {
		fmt.Println("Stopped")
	}()
}
