package main

import "fmt"

func main() {
	//a := [3]int{10, 20, 30}
	//changeArray(&a)
	//fmt.Println("a的值:", a)

	s := []int{10, 20, 30}
	changeSlice(s)
	fmt.Println("s的值是:",s)
}

//数组是值传递
func changeArray(items *[3]int) {
	items[2] = 50
}

func changeSlice(items []int) {
	items[2] = 50
}
