package main

import "fmt"

/*
*

I.数组
II.普通赋值
var numbers = [5]int{1, 2, 3, 4, 5}
numbers := [5]int{1, 2, 3, 4, 5}

II.不确定长度
var balance = [...]float32{1000.0, 2.0, 3.4, 7.0, 50.0}

balance := [...]float32{1000.0, 2.0, 3.4, 7.0, 50.0}

II.通过下标赋值
balance := [5]float32{1:2.0,3:7.0}
*/
func main() {
	var arrTest = [5]int{1, 2, 3, 4, 5}
	fmt.Println(arrTest)
	/**
	连续内容：根据类型不同，间隔的内存地址不一样,int8字节
	0xc00000e2a0
	0xc00000e2a8
	0xc00000e2b0
	0xc00000e2b8
	0xc00000e2c0
	*/
	fmt.Println(&arrTest[0])
	fmt.Println(&arrTest[1])
	fmt.Println(&arrTest[2])
	fmt.Println(&arrTest[3])
	fmt.Println(&arrTest[4])

	//数组遍历方法1
	for i := 0; i < len(arrTest); i++ {
		fmt.Println(arrTest[i])
	}

	//数组遍历2: range 范围遍历
	for key, value := range arrTest {
		fmt.Println("range遍历key 下标 ", key, value)
	}

	var arrTest2 = [10]int{1, 2, 3, 4, 5}
	//数组的类型 arrTest1:[5]int  arrTest2:[10]int   ,数组是带长度的
	fmt.Printf("数组的类型 arrTest1:%T  arrTest2:%T", arrTest, arrTest2)

}
