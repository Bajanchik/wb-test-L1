package main

import (
	"fmt"
	"strconv"
	"sync"
)

// Структура, содержащая отображение для хранения данных и мьютекс для обеспечения безопасности доступа к отображению из разных горутин
type Store struct {
	data map[string]string
	mux  sync.Mutex
}

// Устанавливаем значения элементов отображения
func (s *Store) Set(key, value string) {
	s.mux.Lock()
	defer s.mux.Unlock()

	s.data[key] = value
}

// Получаем значения элементов отображения
func (s *Store) Get(key string) (string, bool) {
	s.mux.Lock()
	defer s.mux.Unlock()

	value, ok := s.data[key]
	return value, ok
}

func main() {
	// Создаём экземпляр отображения
	Store := &Store{
		data: make(map[string]string),
	}

	// Создаём группу ожидания
	var wg sync.WaitGroup

	// Создаём массив ключей
	keys := []string{"key1", "key2", "key3"}

	// Запускаем несколько горутин для записи значений в отображение
	for i, key := range keys {
		wg.Add(1)
		go func(k string, index int) {
			defer wg.Done()
			value := "value" + strconv.Itoa(index+1)
			Store.Set(k, value)
			fmt.Println("Data set...")
		}(key, i)
	}

	// Ждем завершения всех горутин
	wg.Wait()

	// Получаем значения и выводим их на экран
	for _, key := range keys {
		value, ok := Store.Get(key)
		if ok {
			fmt.Printf("Get key: %s, Value: %s\n", key, value)
		} else {
			fmt.Printf("Key: %s not found\n", key)
		}
	}
}
