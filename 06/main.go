package main

import "fmt"

// func swap1(a,b int){
// 	tmp:=a
// 	a=b
// 	b=tmp
// }

func swap2(a, b *int) {
	tem := *a
	*a = *b
	*b = tem
}
func main() {
	a, b := 10, 20
	fmt.Println("a,b = ", a, b)

	swap2(&a, &b)

	fmt.Println("a,b = ", a, b)
}
