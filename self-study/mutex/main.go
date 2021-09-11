// Tạo 1000 goroutine, mỗi goroutine tăng giá trị biến x lên 1 (kết quả mong đợi = 1000)

/* Cách viết bình thường */
// package main

// import (
// 	"fmt"
// 	"log"
// 	"time"
// )

// var x int = 0

// func main() {
// 	timeStart := time.Now()

// 	for i := 0; i < 1000; i++ {
// 		count()
// 	}
// 	fmt.Println(x) // 1000

// 	timeEnd := time.Now()
// 	log.Println(timeEnd.Sub(timeStart)) // 28.139µs

// }

// func count() {
// 	x = x + 1
// }

/* Viết với goroutine */

// package main

// import "fmt"

// var x int = 0

// func main() {
// 	for i := 0; i < 1000; i++ {
// 		go count()
// 	}
// Chương trình chạy mà ko đợi goroutine chạy xong, dẫn đến sai lệch dữ liệu
// 	fmt.Println(x) // 846
// }

// func count() {
// 	x = x + 1
// }

/* Thêm mutex */

// package main

// import (
// 	"fmt"
// 	"sync"
// )

// var x int = 0

// func main() {
// 	wg := sync.WaitGroup{}

// 	for i := 0; i < 1000; i++ {
// 		wg.Add(1)
// 		go count(&wg)
// 	}
// 	wg.Wait()

// 	// do các goroutines chạy đồng thời, đôi lúc chúng truy xuất vào x cùng 1 lúc (race condition),
// 	//dẫn đây sai lệch dữ liệu
// 	fmt.Println(x) // 926
// }

// func count(wg *sync.WaitGroup) {
// 	defer wg.Done()
// 	x = x + 1
// }

/* Sử dụng Mutex để ngăn race condition  */

// package main

// import (
// 	"fmt"
// 	"log"
// 	"sync"
// 	"time"
// )

// var x int = 0

// func main() {
// 	wg := sync.WaitGroup{}
// 	m := sync.Mutex{}
// 	timeStart := time.Now()
// 	for i := 0; i < 1000; i++ {
// 		wg.Add(1)
// 		go count(&wg, &m)
// 	}
// 	wg.Wait()
// 	fmt.Println(x)

// 	// estimate time
// 	timeEnd := time.Now()
// 	log.Println(timeEnd.Sub(timeStart)) // 844.346µs trong trường hợp này, chạy chậm hơn chạy tuần tự
// }

// func count(wg *sync.WaitGroup, m *sync.Mutex) {
// 	defer wg.Done()
// 	m.Lock()
// 	x = x + 1
// 	m.Unlock()
// }

/* Sử dụng channel để ngăn race condition */
package main

import (
	"fmt"
	"log"
	"sync"
	"time"
)

var x int = 0

func main() {
	wg := sync.WaitGroup{}
	ch := make(chan bool, 1) // channel buffer, để tạm thời lưu 1 byte vào bộ đệm

	timeStart := time.Now()

	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go count(&wg, ch)
	}
	wg.Wait()
	fmt.Println(x) // 1.312294ms

	timeEnd := time.Now()
	log.Println(timeEnd.Sub(timeStart))
}

func count(wg *sync.WaitGroup, ch chan bool) {
	defer wg.Done()
	ch <- true
	x = x + 1
	<-ch
}
