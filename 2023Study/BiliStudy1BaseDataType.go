package main

import "fmt"
import "strconv"

var n1 = 222
var s1 = "string1"
var (
	n2 = 222
	s2 = "string1"
)

func main() {
	fmt.Println("begin")
	var (
		n3 = 33
		s4 = "444"
	)
	var n5 int = 214748364
	fmt.Println(n5)
	fmt.Println("\nend1\n")
	fmt.Println(n1, n2, s1, s2, n3, s4, n5)

	var c1 byte = 'a'
	var c2 byte = 'b'
	var c3 byte = 'c'
	fmt.Println(c1, c2, c3)
	fmt.Printf(" 原始字符 %c", c3)

	//bool
	var boT bool = true
	var boF bool = false
	fmt.Println(boT, boF)

	//string
	var str1 string = "/*\nPackage builtin provides documentation for Go's predeclared identifiers.\nThe items documented here are not actually in package builtin\nbut their descriptions here allow godoc to present documentation\nfor the language's special identifiers.\n*/\npackage builtin"
	fmt.Println(str1)
	var str2 string = "aa" +
		"bbb"
	fmt.Println(str2)

	// 其他类型 转换 字符串
	var f1 = 22
	var f1str = strconv.FormatInt(int64(f1), 2)
	fmt.Printf("数字:%d  转成二进制:%s", f1, f1str)
	fmt.Println("\n ")

	var strB string = "true"
	var booaa bool
	booaa, _ = strconv.ParseBool(strB)
	fmt.Printf("字符转bool换后  类型: %T  value: %v ", booaa, booaa)
	fmt.Println("\n ")

	// 字符串 转换   其他类型
	var strC string = "99"
	var i99 int64
	//10代码10进制，64代表int64
	//base指定进制（2到36）
	//bitSize指定结果必须能无溢出赋值的整数类型，0、8、16、32、64 分别代表 int、int8、int16、int32、int64；
	i99, _ = strconv.ParseInt(strC, 10, 32)
	//字符转int换后  类型: int64  value: 99
	fmt.Printf("字符转int换后  类型: %T  value: %v ", i99, i99)
	fmt.Println("\n ")

	var strD string = "99.9"
	var ff1 float64
	ff1, _ = strconv.ParseFloat(strD, 64)
	//字符转float64换后  类型: float64  value: 99.9
	fmt.Printf("字符转float64换后  类型: %T  value: %v ", ff1, ff1)

	fmt.Println("\n ")

	var strE string = "你好"
	var ff2 float64
	ff2, _ = strconv.ParseFloat(strE, 64)
	//错误字符转float64换后  类型: float64  value: 0
	fmt.Printf("错误字符转float64换后  类型: %T  value: %v ", ff2, ff2)
	fmt.Println("\n ")

}
