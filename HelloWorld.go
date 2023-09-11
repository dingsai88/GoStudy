package main

import (
	"fmt"
	"math"
)

/*
*
I.运行源文件:
go run HelloWorld.go

I.编译成二进制exe：

	go build HelloWorld.go

II.执行编译完的程序:HelloWorld
*/
func main() {
	//普通打印
	fmt.Println("Hello, World!Hello, World!Hello, World!" +
		"Hello, World!Hello, World!" +
		"Hello, World!Hello, World!")
	var n1 int
	fmt.Println(n1)

	//占位符测试
	var stockcode int = 1238
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

	var retMax = max(111, 222)
	//占位符
	fmt.Printf("function函数 方法测试 占位符: %d  \n", retMax)

	aSwap, bSwap := swap("Google", "Runoob")
	fmt.Printf("函数方法返回多个值:  %s %s \n ", aSwap, bSwap)

	fmt.Println("函数 作为实参 开始")

	/* 声明函数变量 */
	getSquareRoot := func(x float64) float64 {
		return math.Sqrt(x)
	}

	/* 使用函数 */
	fmt.Println("函数 作为实参:", getSquareRoot(9))

	fmt.Println("\n\n 函数 闭包测试 开始")

	/* nextNumber 为一个函数，函数 i 为 0 */
	nextNumber := getSequence()

	/* 调用 nextNumber 函数，i 变量自增 1 并返回 */
	fmt.Println(nextNumber())
	fmt.Println(nextNumber())
	fmt.Println(nextNumber())

	/* 创建新的函数 nextNumber1，并查看结果 */
	nextNumber1 := getSequence()
	fmt.Println(nextNumber1())
	fmt.Println(nextNumber1())
	fmt.Println("函数 闭包测试结束 开始 \n\n")

}

/*
*
函数方法 返回多个值
*/
func swap(x, y string) (string, string) {
	return y, x
}

/* 函数返回两个数的最大值 */
func max(num1, num2 int) int {
	/* 定义局部变量 */
	var result int

	if num1 > num2 {
		result = num1
	} else {
		result = num2
	}
	return result
}

/*
*
闭包测试
*/
func getSequence() func() int {
	i := 0
	return func() int {
		i += 1
		return i
	}
}
