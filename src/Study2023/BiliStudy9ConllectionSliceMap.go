package main

import "fmt"

/*
*
Go的切片是在数组之上的抽象数据类型，所以创建的切片始终都有一个数组存在。
*/
func main() {
	fmt.Println("begin")
	sliceTest()
	fmt.Println("\nend 1\n")
	sliceTest2()
	fmt.Println("\nend 1\n")
	mapTest()
}

/*
*

2、切片的四种创建方式
var slice1 = make([]int,5,10)
var slice2 = make([]int,5)
var slice3 = []int{}
var slice4 = []int{1,2,3,4,5}
相对应的：

slice1 := make([]int,5,10)
slice2 := make([]int,5)
slice3 := []int{}
slice4 := []int{1,2,3,4,5}

1、数组的三种申明方式
var identifier [len]type
var identifier = [len]type{value1, value2, … , valueN}
var identifier = […]type{value1, value2, … , valueN}
相对应的：
identifier := [len]type{}
identifier := [len]type{value1, value2, … , valueN}
identifier := […]type{value1, value2, … , valueN}
*/
func sliceTest() {
	var arrys1 = [6]int{1, 2, 3, 4, 5, 6}
	fmt.Println(arrys1)
	fmt.Printf("类型:%T \n\n", arrys1)
	//含下标为1-到-下标不为3
	sliceTemp := arrys1[1:3]
	fmt.Println(sliceTemp)

	var slice2 []int = make([]int, 11, 20)
	fmt.Printf("类型:%T \n\n", slice2)

	/* 创建切片 */
	numbers := []int{0, 1, 2, 3, 4, 5, 6, 7, 8}
	printSlice(numbers)

	/* 打印原始切片 */
	fmt.Println("numbers ==", numbers)

	/* 打印子切片从索引1(包含) 到索引4(不包含)*/
	fmt.Println("numbers[1:4] ==", numbers[1:4])

	/* 默认下限为 0*/
	fmt.Println("numbers[:3] ==", numbers[:3])

	/* 默认上限为 len(s)*/
	fmt.Println("numbers[4:] ==", numbers[4:])

	numbers1 := make([]int, 0, 5)
	printSlice(numbers1)

	/* 打印子切片从索引  0(包含) 到索引 2(不包含) */
	number2 := numbers[:2]
	printSlice(number2)

	/* 打印子切片从索引 2(包含) 到索引 5(不包含) */
	number3 := numbers[2:5]
	printSlice(number3)
}

func printSlice(x []int) {
	fmt.Printf("len=%d cap=%d slice=%v\n", len(x), cap(x), x)
} /**
增加数据
*/
func sliceTest2() {
	var numbers []int
	printSlice(numbers)

	/* 允许追加空切片 */
	numbers = append(numbers, 0)
	printSlice(numbers)

	/* 向切片添加一个元素 */
	numbers = append(numbers, 1)
	printSlice(numbers)

	/* 同时添加多个元素 */
	numbers = append(numbers, 2, 3, 4)
	printSlice(numbers)

	/* 创建切片 numbers1 是之前切片的两倍容量*/
	numbers1 := make([]int, len(numbers), (cap(numbers))*2)
	fmt.Println("\n\n复制数据开始")
	/* 拷贝 numbers 的内容到 numbers1 */
	copy(numbers1, numbers)

	fmt.Println("复制数据:源数据:")
	printSlice(numbers)
	fmt.Println("复制数据:新数据:")
	printSlice(numbers1)
}

func mapTest() {
	// map创建
	var map1 = map[string]int{"key1": 1, "key2": 2}
	var map2 map[string]string = make(map[string]string)

	fmt.Println(map1)
	//加数据
	map1["ke3"] = 3
	/* map 插入 key - value 对,各个国家对应的首都 */
	map2["Google"] = "谷歌"

	fmt.Println(map1)
	//删除数据
	delete(map1, "key1")
	fmt.Println(map1)

	//遍历
	for key := range map1 {
		fmt.Println(key, " value", map1[key])
	}

	//遍历2
	for key, value := range map2 {
		fmt.Println("range遍历key 下标 ", key, value)
	}

	/**
	清空，
	1 遍历删除
	2.重新make
	*/
	map2 = make(map[string]string)
}
