package main

import "fmt"

func main() {
	fmt.Println("  begin  ")
	var phoneTemp Phone
	phoneTemp = new(Iphone)
	var result = phoneTemp.call("dingsai")
	fmt.Println("  reuslt 1: ", result)

	phoneTemp = new(SanXingPhone)
	result = phoneTemp.call("dingsai")
	fmt.Println("  reuslt 2: ", result)
}

/**
定义接口
*/
type Phone interface {
	call(strReq string) string
}

/**
定义结构1
*/
type Iphone struct {
	Str2 string " iphone 打电话 1"
}

/**
实现类
*/
func (IphoneObj Iphone) call(strReq string) string {
	fmt.Println("  IphoneCall  ", strReq+IphoneObj.Str2)
	IphoneObj.Str2 = " Iphone 打电话2 "
	fmt.Println("  IphoneCall  ", strReq+IphoneObj.Str2)
	return strReq + " return" + IphoneObj.Str2
}

/**2
定义结构
*/
type SanXingPhone struct {
	Str2 string " iphone 打电话 1"
}

/**
实现类
*/
func (SanXingPhoneObj SanXingPhone) call(strReq string) string {
	fmt.Println("  SanXingPhoneCall  ", strReq+SanXingPhoneObj.Str2)
	SanXingPhoneObj.Str2 = " SanXingPhone 打电话2 "
	fmt.Println("  SanXingPhoneCall  ", strReq+SanXingPhoneObj.Str2)
	return strReq + " return" + SanXingPhoneObj.Str2
}
