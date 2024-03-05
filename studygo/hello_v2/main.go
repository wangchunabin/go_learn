package main

import "fmt"

func main() {
	// a()
	var testArray [3]int
	var numArray = [3]int{1, 2, 3}
	var cityArray = [3]string{"上海", "北京", "重庆"}

	fmt.Println(testArray, numArray, cityArray)

	for i := 0; i < len(cityArray); i++ {
		fmt.Println(cityArray[i])
	}

	for index, value := range cityArray {
		fmt.Println(index, value)
	}

	//多维数组
	a := [3][2]string{
		{"1", "2"},
		{"4", "5"},
		{"6", "7"},
	}
	fmt.Println(a)

	b()

}

func a() {
	for i := 1; i <= 9; i++ {
		for j := 1; j <= i; j++ {
			fmt.Printf("%d*%d=%d\t", j, i, i*j)
		}
		fmt.Println()
	}
}

func b() {
	//声明切片
	fmt.Println("---b---切片------")
	var a []string
	var b = []int{}
	var c = []bool{false, true}
	var d = []bool{false, true}
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(d)
	fmt.Println(a == nil)
	fmt.Println(b == nil)
	// fmt.Println(c == d)
	e := make([]int, 5, 10)
	fmt.Println(e)
	fmt.Printf("%T\n", e)
	fmt.Println(len(e))
	fmt.Println(cap(e))
}
