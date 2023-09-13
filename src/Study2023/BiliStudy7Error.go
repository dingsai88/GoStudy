package main

import (
	"fmt"
)

func main() {
	fmt.Println("错误开始")
	fmt.Println("正常计算1 ", errorTest(6, 3))
	fmt.Println("error计算 ", errorTest(6, 0))
	fmt.Println("正常计算 2", errorTest(6, 3))

}

/*
*
普通错误：不影响下边代码执行
*/
func errorTest(num1, num2 int) int {
	//相当于try catch了
	defer func() {
		err := recover()
		if err != nil {
			fmt.Println("捕获错误")
			fmt.Println("daniel err:", err)
		}
	}()

	result := num1 / num2
	return result
}
