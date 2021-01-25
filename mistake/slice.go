package main

import "fmt"

func main() {
	data := cutting()
	fmt.Printf("data's len:%v,cap:%v\n", len(data), cap(data))
}

func cutting() []int {
	val := make([]int, 1000)
	fmt.Printf("val's len:%v,cap:%v\n ", len(val), cap(val))

	res := make([]int, 10)
	copy(res, val)
	return res
}

//val's len:1000,cap:1000
// data's len:10,cap:10
