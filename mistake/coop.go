package main

import "fmt"


func main() {
	//var test1, test2 []int
	//test1 = []int{1, 2, 3}
	//copy(test2, test1)
	//fmt.Println("test2 的值:",test2)
	var test1, test2 []int
	test1 = []int{1, 2, 3}
	test2=make([]int,len(test1))
	copy(test2,test1)
	fmt.Println("test2 的值:",test2)
}
