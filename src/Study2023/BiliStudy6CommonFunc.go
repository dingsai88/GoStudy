package main

import (
	"fmt"
	"strconv"
	"strings"
	"time"
)

func main() {

	//string 学习
	fmt.Println("string test: ", StringTest())
	fmt.Println("\n\n ")

	//时间学习
	fmt.Println("Date 学习: ", TimeTest())
	fmt.Println("\n\n ")

	//内置函数
	NeiZhiBuiltin()

}

/*
*
内置函数:builtin 内置的
*/
func NeiZhiBuiltin() {
	str := "123"
	fmt.Println(len(str))

	//*num 指针指向的value，&num指针自己的内存地址
	//返回传入类型的指针:
	num := new(int)
	fmt.Printf(" 类型%T 值:%v 地址:%v 指针是:%v",
		num, num, &num, *num)
}

func TimeTest() string {
	now := time.Now()
	fmt.Println("Date now: ", now)
	fmt.Println("Date now:年 ", now.Year())
	fmt.Println("Date now:月 ", now.Month())
	fmt.Println("Date now: 日", now.Day())
	fmt.Println("Date now:时 ", now.Hour())
	fmt.Println("Date now:分 ", now.Minute())
	fmt.Println("Date now: 秒", now.Second())

	fmt.Println("日期格式化字符串 ", now.Format("2006-01-02 15:04:05"))
	return now.GoString()
}

func StringTest() string {
	var str string = "abcedfgh"
	fmt.Println("string len=", len(str))
	r := []rune(str)
	for i := 0; i < len(r); i++ {
		fmt.Printf("rune r= %c \n", r[i])
	}

	//字符串转int
	var str2 string = "12345"
	n, _ := strconv.Atoi(str2)
	fmt.Println("string >> int ", n)

	//int转字符串
	str = strconv.Itoa(123)
	fmt.Println("  int >>string ", str)

	//字符串包含
	bo := strings.Contains(str2, "34")
	fmt.Println("  字符串包含 strings.Contains(str2,\"34\") :", bo)
	//字符串比较 区分不区分大小写
	bo = strings.EqualFold("aaa", "AAA")
	fmt.Println("  不区分大小写比较 strings.EqualFold :", bo)
	fmt.Println("  区分大小写比较 \"a\"==\"A\":", "a" == "A")

	return str
}
