package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())

	timeOutCtx, _ := context.WithTimeout(ctx, time.Second*10) // thời gian sống là 10 giây

	time.AfterFunc(time.Second, func() { // 1s thì cancel context cha
		cancel()
	})

	// khi cancel context cha, context con tự động đc cancel do tính chất context tree
	select {
	case <-timeOutCtx.Done(): 
		fmt.Println("Done roi ne")
	}
}
