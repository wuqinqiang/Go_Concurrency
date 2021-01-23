package main

import "fmt"

func main() {
	res := []int{1, 2, 3}
	var out []*int

	for _, item := range res {
		out = append(out, &item)
	}
	fmt.Println("values:", *out[0], *out[1], *out[1])
	fmt.Println("values address:", out[0], out[1], out[1])
}
