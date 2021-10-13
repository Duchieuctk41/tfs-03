
import (
	"context"
	"fmt"
	"time"
)

func main() {
	ctx, _ := context.WithTimeout(context.Background(), time.Second*1)
	doSomething(ctx)

}

func doSomething(ctx context.Context) {
	canceledChannel := make(chan bool)

	go func() {
		select {
		case <-ctx.Done():
			fmt.Println(ctx.Err())
			canceledChannel <- true
			return
		}
	}()

	isCanceledChannel := <-canceledChannel
	if isCanceledChannel {
		close(canceledChannel)
		return
	}
	time.Sleep(time.Second * 10)
	fmt.Println("end")
}
