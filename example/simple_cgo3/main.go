package main

import (
	"fmt"
	"study_go_c/example/simple_cgo3/pkg/sdk/dynamic"
	"study_go_c/example/simple_cgo3/pkg/sdk/static"
)

func main() {
	res := static.Add(1, 10)
	fmt.Printf("1 + 10 = %d\n", res)
	res = dynamic.Add(1, 10)
	fmt.Printf("1 + 10 = %d\n", res)
}
