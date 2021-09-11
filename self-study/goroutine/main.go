// Xuất ra màn hình console

/* Code đồng bộ */

// package main

// import (
// 	"fmt"
// 	"time"
// )

// func main() {
// 	hello("world")
// 	hello("hieuchunhat")
// }

// func hello(name string) {
// 	for i := 0; i < 5; i++ {
// 		fmt.Println("hello ", name)
// 		time.Sleep(time.Millisecond * 300)
// 	}
// }

/* Code concurrency với goroutine */

// package main

// import (
// 	"fmt"
// 	"time"
// )

// func main() {
// 	go hello("world")
// 	go hello("hieuchunhat")
// 	time.Sleep(time.Second * 2) // thêm sleep không chương trình chạy nhanh quá, goroutine ko kịp chạy
// }

// func hello(name string) {
// 	for i := 0; i < 5; i++ {
// 		fmt.Println("hello ", name)
// 		time.Sleep(time.Millisecond * 300)
// 	}
// }

/* Viết goroutine dùng thêm channel */

// package main

// import (
// 	"fmt"
// 	"time"
// )

// func main() {
// 	ch := make(chan string)
// 	go hello(ch, "world")
// 	go hello(ch, "hieuchunhat")
// 	for {
// 		str, open := <-ch
// 		if !open {
// 			return
// 		}
// 		fmt.Println("hello ", str)
// 	}
// }

// func hello(ch chan<- string, name string) {
// 	for i := 0; i < 5; i++ {
// 		ch <- name
// 		time.Sleep(time.Millisecond * 300)
// 	}
// 	close(ch)
// }

/* Select */

package main

import (
	"fmt"
	"time"
)

func main() {
	ch1 := make(chan string)
	ch2 := make(chan string)
	out := make(chan int)
	go hello(ch1, out, "world")
	go hello(ch2, out, "hieuchunhat")
	for {
		select {
		case a := <-ch1:
			fmt.Println("hello ", a)
		case b := <-ch2:
			fmt.Println("hello ", b)
		case <-out:
			fmt.Println("bye bai")
			return
		}
	}

}

func hello(ch chan string, out chan int, name string) {
	for i := 0; i < 5; i++ {
		ch <- name
		time.Sleep(time.Millisecond * 300)
	}
	out <- 0
}
