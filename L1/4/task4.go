package main

import (
	"fmt"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"
)

func worker(id int, input <-chan string) {
	for data := range input {
		fmt.Printf("Worker %d: %s\n", id, data)
	}
}

func main() {
	// Количество воркеров
	numWorkers := 5

	// Создаем канал для передачи данных
	dataChan := make(chan string)

	// Запускаем воркеры
	var wg sync.WaitGroup
	wg.Add(numWorkers)
	for i := 0; i < numWorkers; i++ {
		go func(workerID int) {
			defer wg.Done()
			worker(workerID, dataChan)
		}(i + 1)
	}

	// Обработка сигнала Ctrl+C для завершения программы и воркеров
	go func() {
		sigChan := make(chan os.Signal, 1)
		signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)
		<-sigChan

		// Закрываем канал и ожидаем завершения всех воркеров
		close(dataChan)
		wg.Wait()

		fmt.Println("Программа завершена")
		os.Exit(0)
	}()

	// Запись данных в канал
	for i := 0; ; i++ {
		select {
		case dataChan <- fmt.Sprintf("Данные %d", i):
		case <-time.After(time.Second):
			// Тайм-аут, если канал переполнен
			fmt.Println("Канал переполнен")
		}
	}
}
