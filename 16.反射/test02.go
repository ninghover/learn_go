// package main

// import(
// 	"fmt"
// 	"reflect"
// )

// func fun(arg interface{}){
// 	t:=reflect.TypeOf(arg)
// 	v:=reflect.ValueOf(arg)
// 	fmt.Printf("type=%v,value=%v\n",t,v)
// }

// func main() {
// 	fun(1.23)
// 	fun("abc")
// }

package main

import (
	"fmt"
	"reflect"
)

type User struct {
	name string
	id   int
}

func (user User) Call() {
	fmt.Println("user 说话")
}

func main() {
	user := User{"huang", 15}
	fun(user)
}

func fun(input interface{}) {
	inputType := reflect.TypeOf(input)
	inputValue := reflect.ValueOf(input)
	fmt.Printf("type=%v, value=%v\n", inputType, inputValue)


	// 通过成员变量
	for i:=0;i<inputType.NumField();i++{
		field:=inputType.Field(i)
		value:=inputValue.Field(i)
		fmt.Printf("name=%v, type=%v, value=%v\n",field.Name,field.Type,value)
	}

	// 获取方法列表
	for i:=0;i<inputType.NumMethod();i++{
		m:=inputType.Method(i)
		fmt.Printf("%s: %v\n",m.Name,m.Type)
	}
}
