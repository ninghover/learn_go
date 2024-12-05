// package main

// import "fmt"

// func main() {
// 	var a string = "123"
// 	var allType interface{} = a
// 	fmt.Println(allType)

// 	fmt.Printf("%T\n",allType)
// }

// package main

// import (
// 	"fmt"
// 	"io"
// 	"os"
// )

// func main() {
// 	tty, err := os.OpenFile("/dev/tty", os.O_RDWR, 0)
// 	// tty是*os.File类型
// 	// *os.File同时实现了io.Reader和io.Writer

// 	if err != nil {
// 		fmt.Println("open file error", err)
// 		return
// 	}
// 	var r io.Reader = tty 	// *os.File实现了io.Reader

// 	// 	r 的实际值是 tty，它的类型是 *os.File。
// 	// *os.File 同时实现了 io.Reader 和 io.Writer。
// 	// 因此，r 可以通过类型断言转换为 io.Writer。
// 	var w io.Writer = tty
// 	w = r.(io.Writer)

// 	w.Write([]byte("hello, world\n"))
// }

package main

import "fmt"

// interface的本质是指针
type Reader interface {
	ReadBook()
}

type Writer interface {
	WriteBook()
}

type Book struct {
}

func (b *Book) ReadBook() {
	fmt.Println("读书")
}

func (b *Book) WriteBook() {
	fmt.Println("写字")
}

/*func main() {
	b:=&Book{}
	b.ReadBook()
	b.WriteBook()
	fmt.Println("-----------------")

	var r Reader =b
	r.ReadBook()
	fmt.Println("----------------")

	var w Writer=r.(Writer)
	w.WriteBook()
}*/
