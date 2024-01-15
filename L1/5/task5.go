package main

import (
	"fmt"
	"time"
)

func sendValues(ch chan<- int) {
	for i := 1; i <= 10; i++ {
		ch <- i // Отправка значения в канал
		fmt.Println("Send:", i)
		time.Sleep(time.Second) // Задержка перед отправкой следующего значения
	}
	close(ch) // Закрытие канала после отправки всех значений
}

func receiveValues(ch <-chan int) {
	for value := range ch {
		fmt.Println("Received:", value) // Чтение значения из канала
	}
}

func main() {
	ch := make(chan int)

	go sendValues(ch)    // Запуск горутины для отправки значений в канал
	go receiveValues(ch) // Запуск горутины для чтения значений из канала

	time.Sleep(10 * time.Second) // Ожидание N секунд

	fmt.Println("Program completed") // Сообщение о завершении программы
}
