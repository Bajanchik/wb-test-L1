package main

import "fmt"

func main() {
	slice := []int{1, 2, 3, 4, 5}
	fmt.Println("Исходный срез:", slice)

	var index int
	fmt.Print("Введите индекс элемента, который нужно удалить: ")
	fmt.Scan(&index)

	// Проверяем, что индекс находится в пределах среза
	if index >= 0 && index < len(slice) {
		// Удаляем i-й элемент из среза
		slice = append(slice[:index], slice[index+1:]...)
		fmt.Println("Срез после удаления элемента:", slice)
	} else {
		fmt.Println("Некорректный индекс")
	}
}
