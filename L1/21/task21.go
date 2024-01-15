package main

import "fmt"

// Интерфейс, определяющий требуемые методы клиента
type Target interface {
	Request() string
}

// Адаптируемый класс, с которым необходимо работать
type Adaptee struct{}

// Метод, относящийся к классу Adaptee
func (a *Adaptee) AdapteeRequest() string {
	return "Адаптируемый запрос"
}

// Адаптер, реализующий интерфейс Target и использующий адаптируемый класс
type Adapter struct {
	Adaptee *Adaptee
}

// Адаптация метода AdapteeRequest() для интерфейса Target
func (a *Adapter) Request() string {
	return "Адаптированный запрос: " + a.Adaptee.AdapteeRequest()
}

// Метод, принимающий в качестве параметра экземпляр интерфейса Target и вызвающий метод Request()
func ClientCode(target Target) {
	fmt.Println(target.Request())
}

func main() {
	// Создаём указатели на адаптируемый класс и на адаптер
	adaptee := &Adaptee{}
	adapter := &Adapter{Adaptee: adaptee}

	ClientCode(adapter)
}
