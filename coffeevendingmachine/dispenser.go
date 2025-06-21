package main

import "fmt"

type Dispenser struct {
	CoffeeObj CoffeeType
}

func (d *Dispenser) DispenseCoffee(coffee CoffeeType) bool {
	fmt.Println("Enjoy your coffee!!")
	return true
}
