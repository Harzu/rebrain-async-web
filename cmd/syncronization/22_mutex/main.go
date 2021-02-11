package main

import (
	"fmt"
	"sync"
	"time"
)

type bills struct {
	mu      *sync.Mutex
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

func main() {
	userBill := &bills{
		mu:      &sync.Mutex{},
		amounts: map[int]int{1: 10000},
	}

	for i := 0; i < 1000; i++ {
		go userBill.subAmount(2000)
	}

	time.Sleep(time.Second)
	fmt.Println(userBill.amounts[1])
}
