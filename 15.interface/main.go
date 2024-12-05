package main

import "fmt"

// interface{}是万能类型
func fun(arg interface{}) {
	fmt.Println(arg)

	// 断言arg数据类型是不是string
	// 断言可以理解为C中的类型转换

	// 1. 无ok检查, 失败panic
	// val:=arg.(string)
	// fmt.Println("val=",val)

	// 2. 有ok检查, 不会panic
	val, ok := arg.(string)
	if !ok {
		fmt.Println("arg的数据类型不是string")
	} else {
		fmt.Println("arg的数据类型是string")
		fmt.Printf("arg = %v , type = %T\n", val, val)
	}
	fmt.Println("---------------------------")
}

type Book struct {
	name string
}

func main() {
	// fun(123)
	fun("123")
	// book := Book{name: "go 语言"}
	// fun(book)
}
