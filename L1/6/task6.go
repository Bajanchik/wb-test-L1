// package main // 1 Способ

// import (
// 	"fmt"
// 	"time"
// )

// func worker(stop <-chan bool) {
// 	for {
// 		select {
// 		case <-stop: // Принимаем сигнал остановки из канала
// 			fmt.Println("Worker stopped")
// 			return
// 		default:
// 			fmt.Println("Working...")
// 			time.Sleep(time.Second)
// 		}
// 	}
// }

// func main() {
// 	stop := make(chan bool)

// 	go worker(stop) // Запуск горутины

// 	time.Sleep(5 * time.Second) // Выполняем работу в течение 5 секунд

// 	stop <- true // Отправляем сигнал остановки в канал

// 	time.Sleep(2 * time.Second) // Даем горутине время на остановку

// 	fmt.Println("Main goroutine stopped")
// }

// package main // 2 Способ

// import (
// 	"context"
// 	"fmt"
// 	"time"
// )

// func worker(ctx context.Context) {
// 	for {
// 		select {
// 		case <-ctx.Done(): // Принимаем сигнал об остановке из контекста
// 			fmt.Println("Worker stopped")
// 			return
// 		default:
// 			fmt.Println("Working...")
// 			time.Sleep(time.Second)
// 		}
// 	}
// }

// func main() {
// 	ctx, cancel := context.WithCancel(context.Background())

// 	go worker(ctx) // Запуск горутины

// 	time.Sleep(5 * time.Second) // Выполняем работу в течение 5 секунд

// 	cancel() // Вызываем функцию cancel для остановки горутины

// 	time.Sleep(2 * time.Second) // Даем горутине время на остановку

// 	fmt.Println("Main goroutine stopped")
// }

package main // 3 sposob

import (
	"fmt"
	"time"
)

func worker(stop *bool) {
	for !*stop { // Проверяем флаг перед каждой итерацией цикла
		fmt.Println("Working...")
		time.Sleep(time.Second)
	}

	fmt.Println("Worker stopped")
}

func main() {
	stop := false

	go worker(&stop) // Запуск горутины

	time.Sleep(5 * time.Second) // Выполняем работу в течение 5 секунд

	stop = true

	time.Sleep(2 * time.Second)

	fmt.Println("Main goroutine stopped")
}
