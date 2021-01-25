package main

import "fmt"

func main() {
	var test1, test2 []int
	test1 = []int{1, 2, 3}
	test2 = make([]int, len(test1),10)
	fmt.Printf("len :%d,cap:%d",len(test2),cap(test2))
	copy(test2, test1)
	fmt.Println("test2 的值:", test2)
}

// test2 的值: [1 2 3]
