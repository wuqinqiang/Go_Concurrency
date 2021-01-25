package main

import "fmt"

type User struct {
	Name string
	Age  int
	//IsChild func(age int) bool
}

//如果可以将每个结构字段与相等运算符进行比较，则可以使用相等运算符 = = 来比较结构变量。
// 如果任何结构字段不具有可比性，则使用相等运算符将导致编译时错误。注意，只有当数组的数据项具有可比性时，数组才具有可比性。

func main() {
	user1 := User{}
	user2 := User{}
	fmt.Println(user1 == user2)
	// invalid operation: user1 == user2 (struct containing func(int) bool cannot be compared)
}
