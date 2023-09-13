package main

import (
	"fmt"
	"time"
)

/*
*
流程控制
*/
func main() {
	fmt.Println("begin")

	/* 局部变量定义 */
	var a int = 100
	/* 判断布尔表达式 */
	if a < 20 {
		fmt.Printf("a 小于 20\n")
	} else if a < 50 {
		fmt.Printf("a 小于 50\n")
	} else {
		fmt.Printf("a 不小于 20\n")
	}

	var grade string = "B"
	switch {
	case grade == "A":
		fmt.Printf("优秀!\n")
	case grade == "B", grade == "C":
		fmt.Printf("良好\n")
	default:
		fmt.Printf("差\n")
	}

	//定义两个管道
	ch1 := make(chan int)
	ch2 := make(chan int)
	//开启线程
	go GoSelectTest(1, ch1)
	go GoSelectTest(2, ch2)
	//等待返回值
	for i := 0; i < 3; i++ {
		select {
		case msg1 := <-ch1:
			fmt.Println("收到", msg1)
		case msg2 := <-ch2:
			fmt.Println("收到", msg2)
		}
		if i >= 1 {
			fmt.Println("走了break")
			break
		}
	}

	//测试goto 和 label1;不打印 b和c
	fmt.Println("aaaaa")
	goto label1
	fmt.Println("bbbb")
	fmt.Println("ccccc")
label1:
	fmt.Println("dddd")

}

// 管道返回
func GoSelectTest(num1 int, ch chan int) int {
	time.Sleep(time.Duration(num1) * time.Second)
	ch <- num1
	return num1
}
