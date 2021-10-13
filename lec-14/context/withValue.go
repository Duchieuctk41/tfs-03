
import (
	"context"
	"fmt"
)

func main() {
	ctx := context.WithValue(context.Background(), "number", 10)
	A(ctx)
}

func A(ctx context.Context) {
	if value := ctx.Value("number"); value != nil {
		ctx := context.WithValue(ctx, "number1", 10/2)
		B(ctx)
	}
}

func B(ctx context.Context) {
	value1 := ctx.Value("number").(int)
	value2 := ctx.Value("number1").(int)
	fmt.Println(value1 + value2)
}
