package main

import (
	"errors"
	"fmt"
)

func main() {
	fmt.Println("错误开始")
	fmt.Println("正常计算1 ", errorTest(6, 3))
	fmt.Println("error计算 ", errorTest(6, 0))
	fmt.Println("正常计算 2", errorTest(6, 3))
	fmt.Println("\n\n")
	recoverExample()
	fmt.Println("\n\n")

	err := errorTest2(0)
	if err != nil {
		fmt.Println("自定义错误:", err)
	}

	err = errorTest2(1)
	if err != nil {
		fmt.Println("自定义错误:", err)
	}
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

/*
*
defer延迟 、panic恐慌、recover组合使用
例子日志:
panic 恐慌开始
panic 测试 recover函数收到的异常信息:panic 抛出的异常信息
*/
func recoverExample() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("panic 测试 recover函数收到的异常信息:", r)
		}
	}()
	fmt.Println("panic 恐慌开始")
	//抛异常
	panic(" panic 抛出的异常信息")
	// 这行代码不会被执行
	fmt.Println("panic 此句不会执行")
}

/**
errors
自定义错误,抛异常2
*/

func errorTest2(num2 int) error {
	var num1 int = 10
	if num2 == 0 {
		return errors.New("除数不能为空")
	} else {
		result := num1 / num2
		fmt.Println("正常计算没有错误:", result)
		return nil
	}

}
