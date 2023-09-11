package main

import (
	"fmt"
)

func main() {
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
