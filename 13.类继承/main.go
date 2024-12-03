package main

import "fmt"

type A struct {
	name string
	age  int
}

type B struct {
	A
	id int
}

func (a *A) eat() {
	fmt.Println("A的eat...")
}

func (a *A) fly() {
	fmt.Println("A的fly...")
}

func (b *B) fly() {
	fmt.Println("B的fly...")
}

func main() {
	a:=A{"huang",18}
	b:=B{A{"hao",18},20}
	a.eat()
	a.fly()
	b.eat()
	b.fly()
	fmt.Println("b.name=",b.name,",b.age=",b.age,",b.id=",b.id)
}
