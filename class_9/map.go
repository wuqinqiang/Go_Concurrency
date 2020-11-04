package main

import (
	"fmt"
	"reflect"
)

func main() {
	m := make(map[string]interface{})
	m["name"] = "wuqq"
	m["age"] = "test1"
	m["ag2"] = "test2"
	m["age3"] = "test3"

	for item := range m {
		fmt.Println("item:", reflect.ValueOf(item))
	}
}
