package main

import (
	"fmt"
	"math"
)

func getSequence() func() int {
	i := 0
	return func() int {
		i += 1
		return i
	}
}

/* 定义结构体 */
type Circle struct {
	radius float64
}

func main() {
	fmt.Println("\n  函数作为实参 开始 \n  ")
	/**
	TODO 1: 函数作为实参
	*/
	/* 声明函数变量 */
	getSquareRoot := func(x float64) float64 {
		return math.Sqrt(x)
	}

	/* 使用函数 */
	fmt.Println(getSquareRoot(9))

	fmt.Println(" \n 函数作为实参 结束 \n\n ")

	fmt.Println(" 函数闭包 开始 \n  ")

	/**
	TODO 2: 函数闭包
	*/
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

	fmt.Println(" \n 函数闭包 结束 \n\n ")

	fmt.Println(" 函数方法 开始 \n  ")
	/**
	TODO 3:函数方法
	*/
	var c1 Circle
	c1.radius = 10.00
	//getArea方法属于 Circle结构体对象的属性
	fmt.Println("圆的面积 = ", c1.getArea())

	fmt.Println(" \n 函数闭包 结束 \n\n ")
}

/**
该 method 属于 Circle 类型对象中的方法
*/
func (c Circle) getArea() float64 {
	//c.radius 即为 Circle 类型对象中的属性
	return 3.14 * c.radius * c.radius
}
