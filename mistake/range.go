package main

import "fmt"

func main() {
	res := []int{1, 2, 3}
	for index, _ := range res {
		res[index] *= 10
	}
	fmt.Println("res:", res)
	// res: [10 20 30]
}
