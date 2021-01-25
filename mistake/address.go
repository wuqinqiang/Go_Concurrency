package main

import "fmt"

//type User struct {
//	Name string
//	Age  int
//}

func (u *User) GetName() {
	fmt.Println(u.Name)
}

type Name interface {
	GetName()
}
//只要值是可寻址的，就可以对值调用指针接收器方法。换句话说，在某些情况下，您不需要该方法的值接收器版本。
func main() {
	//u := User{Name: "wuqinqiang"}
	//u.GetName()
	var u Name = User{Name: "wuqinqiang"}
	u.GetName()
}
