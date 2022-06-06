package main

import "fmt"

func main() {
	map1 := make(map[int]float32)
	map1[1] = 1.0
	map1[2] = 2.0
	map1[3] = 3.0
	map1[4] = 4.0

	// 读取 key 和 value
	for key, value := range map1 {
		fmt.Printf("key is: %d - value is: %f\n", key, value)
	}

	// 读取 key
	for key := range map1 {
		fmt.Printf("key is: %d\n", key)
	}

	// 读取 value
	for _, value := range map1 {
		fmt.Printf("value is: %f\n", value)
	}

	fmt.Println("\n\n  初始化和遍历map \n\n")

	//无法使用，没有初始化的map ,nil map
	//	var countryCapitalMap map[string]string

	//初始化 遍历
	mapa := make(map[string]string)
	mapa["1"] = "one"
	mapa["2"] = "two"

	mapb := make(map[string]string)
	mapb["keya"] = "valuea"
	mapb["keyb"] = "valueB"
	mapb["keyc"] = "valueC"

	for keyStr, valueStr := range mapa {
		fmt.Println(keyStr, " map value: ", valueStr)
	}

	for keyStr, valueStr := range mapb {
		fmt.Println(keyStr, " map value: ", valueStr)
	}

	/*如果确定是真实的,则存在,否则不存在 */
	valueStr, ok := mapb["keyb"]
	fmt.Println(valueStr)
	fmt.Println(ok)
	if ok {
		fmt.Println("value 是", valueStr)
	} else {
		fmt.Println("value不存在 ")
	}

	fmt.Println("\n\n  删除 \n\n")

	/* 创建map */
	countryCapitalMap := map[string]string{"France": "Paris", "Italy": "Rome", "Japan": "Tokyo", "India": "New delhi"}

	fmt.Println("原始地图")

	/* 打印地图 */
	for country := range countryCapitalMap {
		fmt.Println(country, "首都是", countryCapitalMap[country])
	}

	/*删除元素*/

	delete(countryCapitalMap, "France")
	fmt.Println("法国条目被删除")

	fmt.Println("删除元素后地图")

	/*打印地图*/
	for country := range countryCapitalMap {
		fmt.Println(country, "首都是", countryCapitalMap[country])
	}

}
