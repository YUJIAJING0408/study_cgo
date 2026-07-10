package main

import (
	"fmt"
	"study_cgo/example/simple_cgo5/pkg/sdk"
)

func main() {
	s := sdk.NewSdk()
	fmt.Printf("1 + 10 = %d\n", s.Add(1, 10))
	s.Greet("Go")
	s.Free()
}
