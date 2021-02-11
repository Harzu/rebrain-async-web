package main

import (
	"fmt"
	"time"
)

type bills struct {
	amounts map[int]int
}

func (b *bills) subAmount(amount int) {
	if b.amounts[1] > 0 {
		time.Sleep(100 * time.Millisecond)
		b.amounts[1] -= amount
	}
}

func main() {
	userBill := &bills{
		amounts: map[int]int{1: 10000},
	}

	for i := 0; i < 10000; i++ {
		go userBill.subAmount(i)
	}

	time.Sleep(time.Second)
	fmt.Println(userBill.amounts[1])
}
