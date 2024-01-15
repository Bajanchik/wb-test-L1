package main

import (
	"fmt"
	"time"
)

func sleep(seconds int) {
	// Создаём канал, который после определенного времени передаст значение, указывающее на завершение временного интервала
	<-time.After(time.Second * time.Duration(seconds))
}

func main() {
	fmt.Println("Начало программы")
	sleep(5)
	fmt.Println("Прошло 5 секунд")
}
