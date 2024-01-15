package main

import (
	"fmt"
	"math/big"
)

func main() {
	// Инициализируем и вводим переменные
	var aInput, bInput string
	fmt.Print("Введите значение a: ")
	fmt.Scan(&aInput)
	fmt.Print("Введите значение b: ")
	fmt.Scan(&bInput)

	// Переводим введённые строки в числа
	a, _ := new(big.Int).SetString(aInput, 10)
	b, _ := new(big.Int).SetString(bInput, 10)

	// Умножение
	mul := new(big.Int).Mul(a, b)
	fmt.Println("a * b =", mul)

	// Деление
	div := new(big.Float).Quo(new(big.Float).SetInt(a), new(big.Float).SetInt(b))
	fmt.Println("a / b =", div)

	// Сложение
	add := new(big.Int).Add(a, b)
	fmt.Println("a + b =", add)

	// Вычитание
	sub := new(big.Int).Sub(a, b)
	fmt.Println("a - b =", sub)
}
