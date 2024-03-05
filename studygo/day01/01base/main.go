package main

import "fmt" //fmt包中提供格式化，输出，输入的函数。

func main() {
	fmt.Print("hello world\n")
	//演示转义字符的使用 \t 制表符
	fmt.Println("Tom\tJack")
	//换行符
	fmt.Println("Tom\nJack")
	// \r 回车，从当前行的最前面开始输出，覆盖掉以前内容
	fmt.Println("天龙八部雪山飞狐\r张飞")

	fmt.Println("姓名\t年龄\t籍贯\t住址\n张三\t12\t河北\t北京")
}
