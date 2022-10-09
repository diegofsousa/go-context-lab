package main

import (
	"context"
	"fmt"
)

func main() {
	ctx := context.WithValue(
		context.Background(),
		"key", "value",
	)
	showValue(ctx)
}

func showValue(ctx context.Context) {
	ctx2 := context.WithValue(ctx, "key2", "value2")
	fmt.Println(ctx2.Value("key"))
	fmt.Println(ctx2.Value("key2"))
}
