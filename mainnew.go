package main

import (
	"fmt"
	"time"
)

func produce(numbers chan int) {
	for num := range numbers { // Чтение данных из канала
		fmt.Println("Получено:", num)
	}
	close(numbers) // Закрытие канала
}

func main() {
	numbers := make(chan int)

	go produce(numbers) // Запуск производителя в горутине

	for i := 0; i < 5; i++ {
		numbers <- i
		time.Sleep(2 * time.Second)
	}
	fmt.Println("test")
	for i := 0; i < 5; i++ {
		numbers <- i
		time.Sleep(2 * time.Second)
	}
}
