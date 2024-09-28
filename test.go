package main

import (
	"fmt"
	"time"
)

func test(a []int) int {
	s := 0
	for i := 0; i < len(a); i++ {
		s += a[i]
	}
	return s
}

func main() {
	start := time.Now()
	arr := make([]int, 10)
	for k := range arr {
		arr[k] += 10
	}
	a := test(arr[:])
	fmt.Println(a)
	end := time.Now()
	sub := end.Sub(start)

	fmt.Printf("Costed: %d\n", sub)
}
