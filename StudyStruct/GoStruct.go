package main

import "fmt"

type Books struct {
	title   string
	author  string
	subject string
	book_id int
}

func main() {
	fmt.Print(Books{"1", "2", "3", 5}, " \n\n")
	fmt.Print(Books{title: "title", book_id: 3, author: "丁丁"}, " \n\n")

	var book1 Books
	book1.author = "yiming"
	book1.title = "jinpingmei"

	fmt.Println(book1)

	printBook(book1)

}

func printBook(book Books) {
	fmt.Printf("Book title : %s\n", book.title)
	fmt.Printf("Book author : %s\n", book.author)
	fmt.Printf("Book subject : %s\n", book.subject)
	fmt.Printf("Book book_id : %d\n", book.book_id)
}
