package main

import "fmt"

type Book struct {
	title string
	money int
}

func fun1(book Book) {
	// 是值传递
	book.title = "c++"
	fmt.Println("title=", book.title, ",money=", book.money)
}

func fun2(book *Book) {
	// 指针传递
	book.title = "c++"
	fmt.Println("title=", (*book).title, ",money=", book.money) // go语言的语法糖，可以自动解引用
}

func main() {
	var book Book
	book.title = "go语言圣经"
	book.money = 100
	fun1(book)
	fmt.Println("title=", book.title, ",money=", book.money)
	fmt.Println("=========================")
	fun2(&book)
	fmt.Println("title=", book.title, ",money=", book.money)
}
