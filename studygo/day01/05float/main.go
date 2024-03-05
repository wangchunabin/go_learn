package main

import "fmt"

// import "math"

//浮点型
func main() {
	// math.MaxFloat32
	f1 := 1.23456
	fmt.Printf("%T\n", f1) //默认GO语言中的浮点数都是float64类型
	f2 := float32(1.2345678901234567890123456789)
	fmt.Printf("%T\n", f2) //显示声明float32类型
	fmt.Println(f2)

}
