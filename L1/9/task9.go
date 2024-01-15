package main

import (
	"fmt"
)

func main() {
	// Создаем каналы для входных и выходных данных
	input := make(chan int)
	output := make(chan int)

	// Массив с исходными числами
	numbers := []int{1, 2, 3, 4, 5}

	// Запускаем горутину для чтения данных из входного канала
	go readNumbers(input, numbers)

	// Запускаем горутину для обработки данных
	go processNumbers(input, output)

	// Выводим все полученные числа
	fmt.Print("Все числа массива умножены на 2: ")
	for num := range output {
		fmt.Print(num, " ")
	}
}

// Отправляет числа из массива во входной канал
func readNumbers(input chan<- int, numbers []int) {
	// По очереди берём каждое число и отправляем во входной канал
	for _, num := range numbers {
		input <- num
	}

	// Закрываем входной канал, чтобы сигнализировать о завершении передачи данных
	close(input)
}

func processNumbers(input <-chan int, output chan<- int) {
	// По очереди умножаем каждое число из входного канала на два и отправляем в выходной канал
	for num := range input {
		result := num * 2
		output <- result
	}

	// Закрываем выходной канал, чтобы сигнализировать о завершении передачи данных
	close(output)
}
