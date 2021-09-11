// Mô phỏng lại quá trình crawler các bộ phim trên web và lưu vào database

/* Cách thường */

// package main

// import (
// 	"fmt"
// 	"log"
// 	"time"
// )

// func main() {
// 	var arr [5]int
// 	timeStart := time.Now()

// 	crawl(&arr)
// 	create(arr)

// 	// estimate time run
// 	timeEnd := time.Now()
// 	log.Println(timeEnd.Sub(timeStart)) // 1.503092287s chậm hơn so với dùng goroutine và wg (do số loop ít nên chênh lệch ko nhiều)
// }

// func crawl(arr *[5]int) {
// 	for i := 0; i < 5; i++ {
// 		arr[i] = i
// 	}
// }

// func create(arr [5]int) {
// 	for i := 0; i < len(arr); i++ {
// 		fmt.Println(i)
// 		time.Sleep(time.Millisecond * 300)
// 	}
// }

/* Waitgroup*/

package main

import (
	"fmt"
	"log"
	"sync"
	"time"
)

func main() {
	ch := make(chan int)
	wg := sync.WaitGroup{}
	wg.Add(2) // cho waitgroup biết sử dụng bao nhiêu goroutines
	timeStart := time.Now()

	go crawl(&wg, ch)
	go create(&wg, ch)
	wg.Wait() // không cho cái này vô là nó chạy về đích luôn
	// time.Sleep(time.Second * 3) // viết cái này cũng đc, nhưng ko hay lắm, nó ko chạy tính time luôn
	// kiểm tra thời gian chạy
	timeEnd := time.Now()
	log.Println(timeEnd.Sub(timeStart)) // 1.502321918s
}

func crawl(wg *sync.WaitGroup, ch chan<- int) {
	defer wg.Done()
	for i := 0; i < 5; i++ {
		ch <- i
		time.Sleep(time.Millisecond * 300)
	}
	close(ch)
}

func create(wg *sync.WaitGroup, ch <-chan int) {
	defer wg.Done()
	for {
		i, open := <-ch
		if !open {
			return
		}
		fmt.Println(i)
	}
}
