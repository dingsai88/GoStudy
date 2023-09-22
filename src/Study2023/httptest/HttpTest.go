package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)

type Person struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	Age  int    `json:"age"`
}

/*
*
I.get请求

http://127.0.0.1:8080/person?id=11&name=%E4%B8%81%E8%B5%9B&age18

I.post请求
curl -X POST -H "Content-Type: application/json" -d '{\"id\": 123, \"name\": \"John\", \"age\": 30}' http://localhost:8080/person
*/
func main() {
	fmt.Println("开始")
	http.HandleFunc("/person", handlePerson)
	fmt.Println("Server listening on port 8080...")
	http.ListenAndServe(":8080", nil)
}
func handlePerson(w http.ResponseWriter, r *http.Request) {
	fmt.Println("请求开始")
	if r.Method == "POST" {
		fmt.Println("收到请求 post")
		postPerson(w, r)
	} else if r.Method == "GET" {
		fmt.Println("收到请求 get")
		getPerson(w, r)
	} else {
		http.Error(w, "Invalid method", http.StatusMethodNotAllowed)
	}
}

/*
*
post请求
*/
func postPerson(w http.ResponseWriter, r *http.Request) {
	var p Person
	err := json.NewDecoder(r.Body).Decode(&p)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// 在这里对 p 进行业务逻辑处理，例如将其保存到数据库中等
	fmt.Println(p)
	fmt.Fprintf(w, "Created a new person. ID: %d, Name: %s, Age: %d\n", p.ID, p.Name, p.Age)
}

/*
*
get请求
*/
func getPerson(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	name := r.URL.Query().Get("name")
	age := r.URL.Query().Get("age")
	if id == "" || name == "" {
		http.Error(w, "Missing parameters", http.StatusBadRequest)
		return
	}
	idT, _ := strconv.Atoi(id)
	ageT, _ := strconv.Atoi(age)
	p := Person{ID: idT, Name: name, Age: ageT}
	fmt.Println(p)

	jsonResponse(w, &p)
}

/*
*
返回写入
*/
func jsonResponse(w http.ResponseWriter, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	enc := json.NewEncoder(w)
	enc.SetIndent("", "  ")
	enc.Encode(data)
}

/*
*
类型转换
*/
func atoi(s string) int {
	var res int
	for _, c := range s {
		res = res*10 + int(c-'0')
	}
	return res
}
