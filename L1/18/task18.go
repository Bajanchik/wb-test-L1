package main

import (
	"fmt"
	"sync"
)

// Структура-счётчик, содержащая значение счетчика и мьютекс для обеспечения безопасности доступа к счетчику в конкурентной среде
type Counter struct {
	value int64
	mu    sync.Mutex
}

// Выводит текущее значение счётчика, а затем прибавляет к нему единицу
func (c *Counter) Increment() {
	c.mu.Lock()
	defer c.mu.Unlock()
	fmt.Printf("Текущее значение счётчика: %v\n", c.value)
	c.value++
}

func main() {
	// Создаём группу ожидания, которая будет использоваться для блокировки основного потока до завершения всех горутин
	var wg sync.WaitGroup

	// Инициализируем экземпляр счётчика
	counter := Counter{}

	// Запускаем 100 горутин, каждая из которых инкрементирует счетчик
	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func() {
			counter.Increment()
			wg.Done()
		}()
	}

	// Ждём окончания работы всех горутин
	wg.Wait()

	fmt.Println("Итоговое значение счетчика:", counter.value)
}
