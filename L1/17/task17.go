package main

import (
	"fmt"
	"sort"
)

func main() {
	// Создаем отсортированный массив
	nums := []int{2, 4, 6, 8, 10, 12, 14, 16, 18, 20}

	// Задаём число, которое будем искать
	var n int
	fmt.Print("Введите число: ")
	fmt.Scan(&n)

	// Ищем число с помощью бинарного поиска
	index := sort.SearchInts(nums, n)

	// Проверяем, найдено ли число
	if index < len(nums) && nums[index] == n {
		fmt.Printf("Число %d найдено в массиве.", n)
	} else {
		fmt.Printf("Число %d не найдено в массиве.", n)
	}
}
