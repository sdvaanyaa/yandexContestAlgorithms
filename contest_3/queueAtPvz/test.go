package main

import (
	"fmt"
	"os"
)

func main() {
	// Открываем файл для записи
	file, err := os.Create("test_input.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	// Параметры теста
	n := 100000
	b := 100000000 // 10^8
	a := 100000000 // 10^8

	// Пишем количество минут работы и максимальное число обслуживаемых клиентов в файл
	fmt.Fprintf(file, "%d %d\n", n, b)

	// Генерируем строки с количеством клиентов в каждой минуте
	for i := 0; i < n; i++ {
		fmt.Fprintf(file, "%d ", a)
	}

	fmt.Println("Тест сгенерирован в файл test_input.txt")
}
