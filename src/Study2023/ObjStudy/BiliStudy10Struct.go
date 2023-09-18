package main

import "fmt"

type Teacher struct {
	id   int
	name string
	age  int
}
type Student struct {
	id   int
	name string
	age  int
}

// 基础数据类型结构体
type IntStruct int

func main() {
	fmt.Println("创建开始")
	structCreat()

	fmt.Println("\n创建结束\n")

	structCopy()
	fmt.Println("\ncopy结束\n")

	structBindMethod()
	fmt.Println("\n 结构体绑定方法结束\n")

}

func structCreat() {
	//创建方式1
	var t1 Teacher
	//创建方式2
	// var t2 Teacher=Teacher{1,"nnn",19}
	fmt.Println(t1)
	t1.id = 1
	t1.name = "daniel"
	t1.age = 18
	fmt.Println(t1)
	fmt.Println("\n指针 开始\n")

	//创建方式3 指针方式
	var t3Point *Teacher = new(Teacher)
	//创建方式4
	//	var t3Point *Teacher=&Teacher{}
	fmt.Println(t3Point)
	fmt.Println(*t3Point)
	(*t3Point).id = 2
	(*t3Point).name = "name"
	//go 底层编译器进行了优化，这样也可以复制
	t3Point.age = 20
	fmt.Println(*t3Point)
}

/*
*
1.个数类型相同可以强制类型转换
2.结构体起了别名，也要强制类型转换
*/
func structCopy() {
	var t2 Teacher = Teacher{1, "nnn", 19}
	var t3 Student
	//1.个数类型相同可以强制类型转换
	t3 = Student(t2)
	fmt.Println(t3)
	//2.结构体自己起了别名，也要强制类型转换
}

/*
*值传递
1.Teacher结构体绑定本方法
2.值传递，修改不影响原数据
*/
func (t Teacher) structBindMethod() string {
	t.name = "方法里的新名字"
	return t.name
}

/*
*
引用传递:指针类型传递，会变对象传递

(&t2).structBindMethodPoint()
*/
func (t *Teacher) structBindMethodPoint() string {
	(*t).name = "新名字2"
	return t.name
}

/*
*
基础数据类型的结构体绑定的值传递方法
*/
func (i IntStruct) intPrintTest() {
	fmt.Println(i)
}

/*
*
基础数据类型的结构体绑定的引用传递方法
*/
func (i *IntStruct) intPrintPointTest() {
	(*i) = (*i) + 200
	fmt.Println((*i))
}

/*
*
绑定方法：
1.值传递
2.引用传递
*/
func structBindMethod() {
	var t2 Teacher = Teacher{1, "name", 19}

	//值传递
	t2.name = "老名字"
	//结构体力的修改，不影响原值。 打印：方法里的新名字
	fmt.Println(t2.structBindMethod())
	//打印:老名字
	fmt.Println(t2.name)
	fmt.Println("\n指针方法绑定 开始\n")

	//引用传递
	//打印 新名字2
	fmt.Println((&t2).structBindMethodPoint())
	// 打印 新名字2
	fmt.Println(t2.name)
	fmt.Println(t2.structBindMethodPoint())

	fmt.Println("\n 基础数据类型，结构体+方法绑定 开始")
	//基础数据类型加绑定方法
	var int1 IntStruct = 8
	fmt.Println("打印看看 ", int1)
	//输出8
	int1.intPrintTest()
	fmt.Println("打印看看 ", int1)
	//输出 208
	int1.intPrintPointTest()
	fmt.Println("打印看看 ", int1)
	//输出 408
	(&int1).intPrintPointTest()
	fmt.Println("打印看看 ", int1)

}
