package main

import "fmt"

func sum(a int, b int) (ret int) {
	ret = a + b
	return ret
}

func main() {
	result1 := sum(1, 4)
	fmt.Printf("result1: %v\n", result1)
}
