package main

import (
	"fmt"
	"time"
)

var (
	StopChan = make(chan struct{})
	name     string
)

func GorutineStart() {
	go func() {

		for {
			select {
			default:
				message(name)
			case <-StopChan:
				// Останавливаем горутину
				return
			}
			time.Sleep(5 * time.Second)
		}
	}()
}

func message(msg string) {
	fmt.Println(time.Now(), msg)
}

func main() {
	go GorutineStart()
	for {
		fmt.Print("Введите имя: ")
		fmt.Scan(&name)
		time.Sleep(11 * time.Second)
	}
}
