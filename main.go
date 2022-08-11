package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	fmt.Printf("Запуск: %v\n", time.Now())
	ctx1, cancel := context.WithCancel(context.Background())

	go func(ctx context.Context) {
		ctx2, cancel2 := context.WithTimeout(ctx, 3*time.Second) //попробуй поставить тут 3 или 10 секунд
		defer cancel2()

		select {
		case <-ctx2.Done():
			fmt.Printf("Контекст отменен: %v\n", time.Now())
		}

	}(ctx1)

	time.Sleep(5 * time.Second)
	fmt.Printf("Отмена: %v\n", time.Now())
	cancel()

	time.Sleep(7)
}
