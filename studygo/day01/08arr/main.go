package main

import "fmt"

func main() {
	//全局声明
	var arr0 [5]int = [5]int{1, 2, 3, 4, 5}
	var arr1 = [5]int{1, 2, 3, 4, 5}
	var arr2 = [...]int{1, 2, 3, 4, 5, 6}
	var str = [5]string{3: "hello world", 4: "tom"}
	fmt.Println(arr0)
	fmt.Println(arr1)
	fmt.Println(arr2)
	fmt.Println(str)

	for i := 0; i < len(arr0); i++ {
		fmt.Println(arr0[i])
	}

	//局部声明
	a := [3]int{1, 2}               //未初始化元素值为0
	b := [...]int{1, 2, 3, 4, 5, 6} //通过初始化值确定数组长度
	c := [5]int{2: 100, 4: 200}     //通过索引号初始化元素
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
}
