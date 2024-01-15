package main

import (
	"fmt"
	"math"
)

// Структура, представляющая точку в двумерном пространстве
type Point struct {
	x float64
	y float64
}

// Создание нового экземпляра точки с заданными координатами.
func NewPoint(x, y float64) Point {
	return Point{x: x, y: y}
}

// Вычисление расстояния между двумя точками.
func Distance(p Point, q Point) float64 {
	dx := p.x - q.x
	dy := p.y - q.y
	return math.Sqrt(dx*dx + dy*dy)
}

func main() {
	// Создаем две точки
	var x, y float64
	fmt.Print("Введите координату x для первой точки: ")
	fmt.Scan(&x)
	fmt.Print("Введите координату y для первой точки: ")
	fmt.Scan(&y)
	p := NewPoint(x, y)

	fmt.Print("Введите координату x для второй точки: ")
	fmt.Scan(&x)
	fmt.Print("Введите координату y для второй точки: ")
	fmt.Scan(&y)
	q := NewPoint(x, y)

	// Вычисляем и выводим расстояние между ними
	distance := Distance(p, q)
	fmt.Printf("Расстояние между точками (%g, %g) и (%g, %g) равно %g\n", p.x, p.y, q.x, q.y, distance)
}
