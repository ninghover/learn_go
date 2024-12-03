package main

import "fmt"

// 本质是一个指针
type Animal interface {
	Sleep()
	GetType() string
}

// Cat类
type Cat struct {
}

func (cat *Cat) Sleep() {
	fmt.Println("小猫在睡觉...")
}

func (cat *Cat) GetType() string {
	return "小猫"
}

// Dog类
type Dog struct {
}

func (dog *Dog) Sleep() {
	fmt.Println("小狗在睡觉...")
}

func (dog *Dog) GetType() string {
	return "小狗"
}

func fun(animal Animal) {
	animal.Sleep()
	fmt.Println(animal.GetType())
}

func main() {
	cat := Cat{}
	fun(&cat)

	dog := Dog{}
	fun(&dog)
}
