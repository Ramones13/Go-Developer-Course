package main

import (
	"context"
	"fmt"
)

type ctxKey string

const someKey1 ctxKey = "some key1"
const someKey2 ctxKey = "some key2"

func main() {

	ctx := context.Background()
	ctx = context.WithValue(ctx, someKey1, "some value1")
	ctx = context.WithValue(ctx, someKey2, "some value2")
	do(ctx)

}

func do(ctx context.Context) {

	fmt.Printf("%v: %v\n", someKey1, ctx.Value(someKey1))
	fmt.Printf("%v: %v\n", someKey2, ctx.Value(someKey2))

}
