package main

import "fmt"

//类型定义---通过type关键字的定义，NewInt就是一种新的类型，它具有int的特性。
type NewInt int

//类型别名
type MyInt = int

func main() {
	var a NewInt
	var b MyInt

	fmt.Printf("type of a:%T\n", a)
	fmt.Printf("type of b:%T\n", b)

}
