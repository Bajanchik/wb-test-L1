package main

import (
	"fmt"
)

func main() {
	// Принимаем значение переменной int64
	var num int64
	fmt.Print("Введите исходное число: ")
	fmt.Scan(&num)

	// Принимаем номер бита
	var i int
	fmt.Print("Введите номер бита, который нужно установить: ")
	fmt.Scan(&i)

	// Принимаем значение бита
	var bitValue int
	fmt.Print("Введите значение, которое будет установлено (может быть 0 или 1): ")
	fmt.Scan(&bitValue)

	// Проверяем и устанавливаем значение бита
	if bitValue == 1 {
		// Устанавливаем i-й бит в 1
		num |= 1 << i
	} else if bitValue == 0 {
		// Устанавливаем i-й бит в 0
		num &= ^(1 << i)
	} else {
		fmt.Print("Значение бита может быть лишь 0 или 1")
		return
	}

	// Выводим результат
	fmt.Printf("Полученное число: %d", num)
}
