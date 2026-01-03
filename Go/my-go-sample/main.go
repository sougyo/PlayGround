package main

import "fmt"

// Add returns the sum of two integers.
func Add(a, b int) int {
	return a + b
}

func main() {
	fmt.Println("Hello, Go World")
	result := Add(1, 2)
	fmt.Printf("1 + 2 = %d\n", result)
}
