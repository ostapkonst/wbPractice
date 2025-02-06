// 24. Разработать программу нахождения расстояния между двумя точками, которые представлены в виде структуры Point с инкапсулированными параметрами x,y и конструктором.
package main

import (
	"fmt"
	"math"
)

// Структура Point с инкапсулированными полями x и y
type Point struct {
	x, y float64
}

// Конструктор для создания новой точки
func NewPoint(x, y float64) Point {
	return Point{x: x, y: y}
}

// Метод для получения значения x
func (p Point) GetX() float64 {
	return p.x
}

// Метод для получения значения y
func (p Point) GetY() float64 {
	return p.y
}

// Функция для вычисления расстояния между двумя точками
func Distance(p1, p2 Point) float64 {
	dx := p1.GetX() - p2.GetX()
	dy := p1.GetY() - p2.GetY()
	return math.Sqrt(dx*dx + dy*dy)
}

func main() {
	// Создаем две точки с помощью конструктора
	point1 := NewPoint(1.0, 2.0)
	point2 := NewPoint(4.0, 6.0)

	// Вычисляем расстояние между точками
	distance := Distance(point1, point2)

	// Выводим результат
	fmt.Printf("Расстояние между точками (%.2f, %.2f) и (%.2f, %.2f) = %.2f\n",
		point1.GetX(), point1.GetY(), point2.GetX(), point2.GetY(), distance)
}
