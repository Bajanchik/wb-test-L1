package main

import (
	"fmt"
	"strings"
)

func isUnique(s string) bool {
	// Приводим строку к нижнему регистру
	s = strings.ToLower(s)

	// Создаем отображение для отслеживания символов
	charMap := make(map[rune]bool)

	// Перебираем каждый символ в строке
	for _, ch := range s {
		// Если символ уже присутствует в отображении, то строка не уникальна
		if charMap[ch] {
			return false
		}
		// Добавляем символ в отображение
		charMap[ch] = true
	}

	// Если прошли все символы и не было повторений, то строка уникальна
	return true
}

func main() {
	var str string
	fmt.Print("Введите строку: ")
	fmt.Scan(&str)

	if isUnique(str) {
		fmt.Println("Все символы в строке уникальны.")
	} else {
		fmt.Println("В строке есть повторяющиеся символы.")
	}
}
