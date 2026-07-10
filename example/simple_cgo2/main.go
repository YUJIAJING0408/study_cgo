package main

import (
	"fmt"
	"study_go_c/example/simple_cgo2/pkg/gosdk"
)

func main() {
	res := gosdk.Add(1, 10)
	fmt.Printf("1 + 10 = %d\n", res)
	gosdk.Greet("Go")
}
