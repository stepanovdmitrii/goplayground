package main

import (
	"log"
	"net/http"
	"sync"
)

var NumberOfWorkers int = 8

func main() {
	var urls = []string{
		"http://ya.ru",
		"http://ya.ru",
		"http://ya.ru",
		"http://ya.ru",
		"https://ozon.ru",
		"https://ozon.ru",
		"https://ozon.ru",
		"https://ozon.ru",
		"undefined site",
	}
	ch := make(chan string)
	wg := startWorkers(ch)
	for _, url := range urls {
		ch <- url
	}
	close(ch)
	wg.Wait()
}

func startWorkers(input <-chan string) *sync.WaitGroup {
	wg := &sync.WaitGroup{}
	for i := 0; i < NumberOfWorkers; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for {
				url, ok := <-input
				if !ok {
					return
				}
				resp, err := http.Get(url)
				if err != nil {
					log.Printf("ошибка запроса \"%s\": %s\n", url, err.Error())
					continue
				}
				if resp.StatusCode == http.StatusOK {
					log.Printf("\"%s\" - ok\n", url)
				} else {
					log.Printf("\"%s\" - not ok\n", url)
				}
			}

		}()
	}
	return wg
}

func merge(cs ...<-chan int) <-chan int {
	result := make(chan int)
	wg := &sync.WaitGroup{}
	for _, ch := range cs {
		ch := ch
		wg.Add(1)
		go func() {
			defer wg.Done()
			for v := range ch {
				result <- v
			}
		}()
	}

	go func() {
		wg.Wait()
		close(result)
	}()

	return result
}
