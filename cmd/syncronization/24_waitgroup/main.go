package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

type bills struct {
	mu      *sync.RWMutex
	amounts map[int]int
}

func (b *bills) subAmount(amount int) {
	b.mu.Lock()
	defer b.mu.Unlock()

	if b.amounts[1] > 0 {
		time.Sleep(100 * time.Millisecond)
		b.amounts[1] -= amount
	}
}

func (b *bills) getAmount() int {
	b.mu.RLock()
	defer b.mu.RUnlock()

	return b.amounts[1]
}

func main() {
	userBill := &bills{
		mu:      &sync.RWMutex{},
		amounts: map[int]int{1: 10000},
	}

	wg := &sync.WaitGroup{}

	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func(amount int) {
			defer wg.Done()
			userBill.subAmount(1000)
		}(i)
	}

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(gCount int) {
			defer wg.Done()

			time.Sleep(time.Duration(rand.Intn(10)) * time.Second)
			amount := userBill.getAmount()
			fmt.Println(gCount, amount)
		}(i)
	}

	wg.Wait()
}
