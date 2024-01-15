package main

import (
	"fmt"
)

func main() {
	// Инициализируем переменную-интерфейс
	var val interface{}

	// Присваиваем переменной значения разных типов и вызываем функцию, определяющую тип переменной
	val = 42
	checkType(val)

	val = "Hello, World!"
	checkType(val)

	val = true
	checkType(val)

	val = make(chan int)
	checkType(val)
}

func checkType(val interface{}) {
	switch val.(type) {
	case int:
		fmt.Println("Тип переменной: int")
	case string:
		fmt.Println("Тип переменной: string")
	case bool:
		fmt.Println("Тип переменной: bool")
	case chan int:
		fmt.Println("Тип переменной: channel")
	default:
		fmt.Println("Неизвестный тип")
	}
}
