// 1. Дана структура Human (с произвольным набором полей и методов).
// Реализовать встраивание методов в структуре Action от родительской структуры Human (аналог наследования).
package main

import "fmt"

type Human struct {
	name string
	age  int
}

func (h *Human) GetInfo() {
	fmt.Printf("name: %s, age: %d", h.name, h.age)
}

type Action struct {
	Human
	surname string
}

func (a *Action) GetInfo() {
	fmt.Printf("name: %s, age: %d, surname: %s", a.name, a.age, a.surname)
}

type Inforer interface {
	GetInfo()
}

func main() {
	human := Human{
		name: "John",
		age:  30,
	}

	action := Action{
		Human:   human,
		surname: "Doe",
	}

	// но, если мы объявим peoples := []Inforer{human, action}, то будет ошибка
	peoples := []Inforer{&human, &action}

	for _, p := range peoples {
		p.GetInfo()
	}
}
