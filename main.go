package main

import (
	data "Go-Tuto-Go-Rutines/data"
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func main() {
	start := time.Now()
	wg := &sync.WaitGroup{}
	m := &sync.RWMutex{}

	for i := 0; i < 10; i++ {
		wg.Add(1)
		//go showGoRoutine(i, wg)
		go readBook(i, wg, m)
	}

	wg.Wait()
	duration := time.Since(start).Milliseconds()
	fmt.Printf("%dms\n", duration)
}

func readBook(id int, wg *sync.WaitGroup, m *sync.RWMutex) {
	data.FinishBook(id, m)
	delay := rand.Intn(800)
	time.Sleep(time.Millisecond * time.Duration(delay))

	wg.Done()
}

func showGoRoutine(id int, wg *sync.WaitGroup) {
	delay := rand.Intn(500)
	fmt.Printf("GoRutine #%d with %dms\n", id, delay)

	time.Sleep(time.Millisecond * time.Duration(delay))

	wg.Done()
}
