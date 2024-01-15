package main

import "fmt"

// Структура Human
type Human struct {
	name string
	age  int
}

// Метод вывода информации о человеке
func (h Human) Info() {
	fmt.Printf("Name: %s\n", h.name)
	fmt.Printf("Age: %d\n", h.age)
}

// Структура Action, встраивающая структуру Human
type Action struct {
	Human
}

// Дополнительный метод для структуры Action
func (a Action) Walk() {
	fmt.Printf("%s is %d years old", a.name, a.age)
}

// Пример использования
func main() {
	// Создаем объект структуры Action
	action := Action{
		Human: Human{
			name: "Ivan",
			age:  19,
		},
	}

	// Вызываем методы
	action.Info()
	action.Walk()
}
