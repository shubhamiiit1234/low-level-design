package main

const (
	Sugar  string = "SUGAR"
	Coffee string = "COFFEE"
	Milk   string = "MILK"
	Water  string = "WATER"
)

type Ingredient struct {
	Name     string
	Quantity int
}
