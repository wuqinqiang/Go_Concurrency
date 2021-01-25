package main

import "fmt"

func main() {
	s := "test"
	bytes := []byte(s)
	bytes[0] = 'T'
	fmt.Println("s的值:", string(bytes))
	// s的值: Test
}
