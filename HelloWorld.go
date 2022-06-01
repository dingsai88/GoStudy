package main

import "fmt"

/**
I.运行源文件:
go run HelloWorld.go

I.编译成二进制exe：
 go build HelloWorld.go

II.执行编译完的程序:HelloWorld

*/
func main() {
	//普通打印
	fmt.Println("Hello, World!")

	//占位符测试
	var stockcode = 1238
	var enddate = "2020-12-3144484"
	var url = "Code=%d&endDate=%s"
	var target_url = fmt.Sprintf(url, stockcode, enddate)
	fmt.Println("占位符测试" + target_url)

	//定义一个变量
	var str string = "stringstr"
	fmt.Println("定义单个变量1:" + str)
	fmt.Println("定义单个变量2:", str)

	//定义多个变量
	var a, b int = 1, 2
	fmt.Println("定义多个变量:", a+b)

	//变量未初始化
	var strNoInit string
	fmt.Println("变量未初始化string:" + strNoInit)

	var intNoInit int
	fmt.Println("变量未初始化int:", intNoInit)

	var boNoInit bool
	fmt.Println("变量未初始化bool:", boNoInit)

	var intVal int
	intVal2 := 1
	fmt.Println("变量未初始化intVal2:", intVal, intVal2)

	//变量未初始化
	var strEquals1 string = "abc"
	var strEquals2 string
	var strEquals3 *string
	//内存中将 i 的值进行了拷贝：
	strEquals2 = strEquals1
	fmt.Println("赋值判断1:", strEquals1, strEquals2)
	strEquals1 = "xyz"
	fmt.Println("赋值判断2:", strEquals1, strEquals2)

	strEquals3 = &strEquals1

	fmt.Println("赋值判断3 引用类型:", strEquals1, strEquals3)
	strEquals1 = "xyz2"
	fmt.Println("赋值判断3 引用类型2:", strEquals1, strEquals3)

}
