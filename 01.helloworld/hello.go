package main //有main函数的，一定是main包

import "fmt"

func hello() {
	fmt.Println("hello, go")
}

func main() {
	fmt.Println("main func")
	hello()
}
