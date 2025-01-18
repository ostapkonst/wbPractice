// 21. Реализовать паттерн «адаптер» на любом примере.
package main

import "fmt"

// Про сам паттерн можно прочитать здесь: https://refactoring.guru/ru/design-patterns/adapter
// или тут: https://metanit.com/sharp/patterns/4.2.php

/*
Адаптер — это структурный паттерн проектирования, который позволяет объектам с несовместимыми интерфейсами работать вместе.

Адаптеры могут не только переводить данные из одного формата в другой, но и помогать объектам с разными интерфейсами работать сообща. Это работает так:

— Адаптер имеет интерфейс, который совместим с одним из объектов.
— Поэтому этот объект может свободно вызывать методы адаптера.
— Адаптер получает эти вызовы и перенаправляет их второму объекту, но уже в том формате и последовательности, которые понятны второму объекту.
*/

type Requester interface {
	Request(url string)
}

type Adaptee struct {
}

func (a Adaptee) SpecificRequest(url string) {
	// эмулируем отправку пакета по `url`
	fmt.Println("adaptee request:", url)
}

type Adapter struct {
	adaptee *Adaptee
}

func (a *Adapter) Request(url string) {
	a.adaptee.SpecificRequest(url)
}

func main() {
	// адаптер адаптирует Adaptee (структуру) к Requester (интерфейс)
	var r Requester = &Adapter{adaptee: &Adaptee{}}

	r.Request("http://example.com")
}
