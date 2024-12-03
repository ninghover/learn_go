// 声明slice的四种方式
package main

import "fmt"

func main() {
	// 1.
	slice1 := []int{1, 2, 3}
	fmt.Printf("len = %d, slice1 = %v\n", len(slice1), slice1)

	//2.
	// var slice2 []int //声明
	// if slice2 == nil {
	// 	fmt.Println("slice2是个空slice")
	// } else {
	// 	fmt.Println("slice2非空")
	// }
	// slice2 = make([]int, 3) //开辟空间
	slice2 := make([]int, 3)
	fmt.Printf("len = %d, slice2 =%v\n", len(slice2), slice2)

	//3.
	var slice3 []int = make([]int, 5)
	fmt.Printf("len = %d, slice3 =%v\n", len(slice3), slice3)

	//4.
	slice4 := make([]int, 3)
	fmt.Printf("len = %d, slice4 =%v\n", len(slice4), slice4)

	//5.传递一个长度和容量
	slice5 := make([]int, 6, 8)
	slice5[0], slice5[1] = 1, 2
	fmt.Printf("len=%d,cap=%d,slice5=%v\n", len(slice5), cap(slice5), slice5)
	slice5 = append(slice5, 1)
	fmt.Printf("len=%d,cap=%d,slice5=%v\n", len(slice5), cap(slice5), slice5)
	slice5 = append(slice5, 2)
	fmt.Printf("len=%d,cap=%d,slice5=%v\n", len(slice5), cap(slice5), slice5)

	//6.类似python这样的切片
	slice6 := slice5[:2]
	fmt.Printf("len=%d,cap=%d,slice6=%v\n", len(slice6), cap(slice6), slice6)
	fmt.Printf("%d %d\n", &slice5[0], &slice6[0]) //实际上是指向同一片内存地址，一个也会影响到另一个

	//7.copy
	slice7 := make([]int, 2, 3)
	copy(slice7, slice5) //把slice5拷贝到slice7中，拷贝len(slice7)个
	fmt.Printf("len=%d,cap=%d,slice7=%v\n", len(slice7), cap(slice7), slice7)
}
