package main

import (
	"fmt"
	"time"
)

type bill struct {
	amount int
}

func (b *bill) subAmount(amount int) {
	if b.amount > 0 {
		time.Sleep(100 * time.Millisecond)
		b.amount -= amount
	}
}

func main() {
	userBill := &bill{amount: 10}

	go userBill.subAmount(10)
	go userBill.subAmount(10)

	time.Sleep(time.Second)
	fmt.Println(userBill.amount)
}
