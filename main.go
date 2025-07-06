package main

import (
	"fmt"
	"sync"
)

func main() {
	intChannel := make(chan int) //создаем небуфферизированный канал инт
	var wg sync.WaitGroup

	wg.Add(100) //добавим счетчик на выполнение 100 операций
	for i := 1; i <= 100; i++ {
		go func(val int) {
			defer wg.Done()
			intChannel <- val
		}(i)
	}

	go func() { //закрываем канал по завершению
		wg.Wait()
		close(intChannel)
	}()

	for value := range intChannel { //выводим данные из канала
		fmt.Println(value)
	}

}
