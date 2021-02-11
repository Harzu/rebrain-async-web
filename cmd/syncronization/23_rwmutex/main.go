package main

import (
	"fmt"
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

	for i := 0; i < 1000; i++ {
		go userBill.subAmount(i)
	}

	for i := 0; i < 10; i++ {
		go func(gCount int) {
			for {
				amount := userBill.getAmount()
				fmt.Println(gCount, amount)
			}
		}(i)
	}

	time.Sleep(time.Second)
}
