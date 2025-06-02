package main

import "fmt"

type CashPayment struct {
}

func (c *CashPayment) ProcessPayment(coffee CoffeeType, insertedAmount int) bool {
	if insertedAmount < int(coffee.GetPrice()) {
		fmt.Println("Insufficient money inserted!!")
		return false
	}

	fmt.Println("Payment successfully processed!!")
	return true
}

// Change/Refund
func (c *CashPayment) ReturnMoney(coffee CoffeeType, money int) int {
	fmt.Println("Can't Process, Take your money!!")
	return money
}
