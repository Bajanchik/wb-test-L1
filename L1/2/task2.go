package main

import (
	"fmt"
	"sync"
)

func main() {
	// Массив чисел
	numbers := []int{2, 4, 6, 8, 10}

	// Создаем waitGroup, чтобы дождаться завершения всех горутин
	var wg sync.WaitGroup

	// Для каждого числа в массиве запускаем горутину
	for _, num := range numbers {
		// Добавляем горутину в группу
		wg.Add(1)

		// Вызов горутины
		go func(n int) {
			// Отложенно удаляем горутину из группы по ее завершению
			defer wg.Done()

			// Вычисляем и выводим квадрат числа
			square := n * n
			fmt.Println(square)
		}(num)
	}

	// Ждем, пока все горутины завершатся
	wg.Wait()
}
