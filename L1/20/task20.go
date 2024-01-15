package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	// Задаём строку, которую будем переворачивать
	var input string
	fmt.Print("Введите строку: ")
	in := bufio.NewScanner(os.Stdin)
	in.Scan()
	input = in.Text()

	// Разделение строки на слова
	words := strings.Fields(input)

	// Создаём срез, в который будем помещать слова в обратном порядке
	reversedWords := make([]string, len(words))

	// Берём слово с конца строки и помещаем его в начало созданного среза, после чего "сдвигаем" индексы на одну позицию, и так, пока не "перенесём" все слова
	for i, j := 0, len(words)-1; i < len(words); i, j = i+1, j-1 {
		reversedWords[i] = words[j]
	}

	// Формирование и вывод результата
	fmt.Println(strings.Join(reversedWords, " "))
}
