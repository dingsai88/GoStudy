package main

import (
	"fmt"
)

// 局部变量方法>init>main
var testFunc int = testOrder1()

func testOrder1() int {
	fmt.Println("第1个输出 局部变量初始化")
	return 1
}

func init() {
	fmt.Println("第2个输出 init")
}

func main() {
	fmt.Println("第3个输出 main")

	var a = 6
	var b = 4
	var result1, result2 = calc(a, b)
	//全返回值: 10 2
	fmt.Println("全返回值:", result1, result2)
	//只要一个返回值,不需要的_
	var result3, _ = calc(a, b)
	fmt.Println("部分返回值:", result3)

	//不固定入参
	var result4 = test(a, b)
	fmt.Println("不固定个参数:", result4)

	//值传递和引用传递
	var result5 = 55
	fmt.Println("传入前:", result5)
	pointTest(&result5)
	fmt.Println("传入后:", result5)

	//函数也可以是变量
	var result6 = 11
	//调用testFunc(&result6)  =pointTest(&result6)
	testFunc := pointTest
	testFunc(&result6)
	fmt.Println("传入后:", result6)

	//自定义数据类型 type myInt int
	type myInt int
	var myInt1 myInt = 30
	var tempMyInt1 int = 20
	tempMyInt1 = int(myInt1)
	fmt.Println("tempMyInt1:", tempMyInt1)

	//闭包：一共公共sum变量，调用多次sum可以复用
	biBao := getSum()
	fmt.Println(biBao(1))
	fmt.Println(biBao(2))
	fmt.Println(biBao(3))

	//defer 延迟
	fmt.Println("defer test: ", deferTest(1, 2))

}

// 两个返回值
func calc(a int, b int) (int, int) {
	var sum = a + b
	var minus = a - b
	return sum, minus
}

// 可变长参数
func test(args ...int) int {
	var sum int
	for i := 0; i < len(args); i++ {
		sum += args[i]
	}
	return sum
}

// 传入指针 地址
func pointTest(args *int) {
	*args = 99
}

/*
*
闭包可以保留上次引用的某个值
闭包：一共公共sum变量，调用多次sum可以复用
*/
func getSum() func(int) int {
	var sum int = 0
	return func(num int) int {
		sum = sum + num
		return sum
	}
}

/*
*
遇到defer，后边的diam 压入栈中(先进后出)，不执行
函数执行完毕，执行栈中语句(先进后出)
应用场景，需要关闭的资源可以随手写，退出时候自动帮忙退出
*/
func deferTest(num1, num2 int) int {
	defer fmt.Println("defer打印顺序3 num1=", num1)
	defer fmt.Println("defer打印顺序2 num2=", num2)
	var sum int = num1 + num2
	fmt.Println("defer打印顺序1 sum=", sum)
	return sum
}
