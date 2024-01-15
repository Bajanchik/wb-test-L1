package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	// Задаём строку, которую будем переворачивать
	var input string
	fmt.Print("Введите строку: ")
	in := bufio.NewScanner(os.Stdin)
	in.Scan()
	input = in.Text()

	// Превращаем строку в срез символов
	runes := []rune(input)

	// Берём символы с начала и конца среза и меняем их местами после чего "сдвигаем" значения индексов ближе к середине среза и так пока не достигнем середины
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}

	fmt.Println("Перевернутая строка:", string(runes))
}
