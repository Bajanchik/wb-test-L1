package main

import (
	"fmt"
	"math"
)

func main() {
	// Исходные значения температур
	temperatures := []float64{-25.4, -27.0, 13.0, 19.0, 15.5, 24.5, -21.0, 32.5}

	// Отображение, которое будет содержать температурные группы
	groups := make(map[int][]float64)

	for _, temperature := range temperatures {
		var group int
		// Выясняем, в какую группу необходимо добавить температуру
		if temperature < 0 {
			group = int(math.Ceil(temperature/10.0)) * 10
		} else {
			group = int(math.Floor(temperature/10.0)) * 10
		}

		// Добавляем температуру в соответствующую группу
		groups[group] = append(groups[group], temperature)
	}

	// Выводим все группы и содержащиеся в них значения
	for group, values := range groups {
		fmt.Printf("%d: %.1f\n", group, values)
	}
}
