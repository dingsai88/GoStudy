package main

import "fmt"

/*
*
指针
*/
func main() {
	fmt.Println("\nbegin\n")

	var i int64 = 1
	//i变量的类型:int64 value:1   内存地址:0xc00000a0e0
	fmt.Printf("i变量的类型:%T value:%v   内存地址:%p", i, i, &i)
	fmt.Println("\n ")

	//ptr1指针类型返回地址， *ptr1返回指针的value  &ptr1:返回本指针的地址
	var ptr1 *int64 = &i
	//i变量的类型:int64 value:1   内存地址:0xc00000a0e0
	fmt.Printf("ptr1变量是指针类型:%T value:%v   内存地址:%p  指针原本的值是:%v  指针自己的地址是:%v", ptr1, ptr1, ptr1, *ptr1, &ptr1)
	fmt.Println("\n ")
	//用指针修改value
	*ptr1 = 20
	//ptr1变量是指针类型:*int64 value:0xc00000a0e0   内存地址:0xc00000a0e0  指针原本的值是:20  指针自己的地址是:0xc00005a028
	fmt.Printf("ptr1变量是指针类型:%T value:%v   内存地址:%p  指针原本的值是:%v  指针自己的地址是:%v", ptr1, ptr1, ptr1, *ptr1, &ptr1)
	fmt.Println("\n ")

	//用户输入 获取终端输入，用户输入
	fmt.Println("请输入 ")
	var impStr string
	//到换行结束输入;必须传入地址
	fmt.Scanln(&impStr)
	fmt.Println("您输入的是 " + impStr)

	//用户输入 获取终端输入，用户输入
	fmt.Println("请录入学生的年龄，姓名，成绩，是否是VIP ")
	var age int
	var name string
	var score float32
	var isVIP bool
	//到换行结束输入;必须传入地址
	fmt.Scanf("%d %s %f %t", &age, &name, &score, &isVIP)
	//您输入的 12 dingsai 33.4 true
	fmt.Printf("您输入的 %v %v %v %v ", age, name, score, isVIP)

}
