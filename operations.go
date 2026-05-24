package main

import "fmt"

func CalcStoreTotal(data ProductData) {
	var storeTotal float64
	var channel = make(chan float64)
	for category, group := range data {
		go group.TotalPrice(category, channel)
	}
	for range data {
		storeTotal += <-channel
	}
	fmt.Println("Total:", ToCurrency(storeTotal))
}

func (group ProductGroup) TotalPrice(category string, resultChanel chan float64) {
	var total float64
	for _, p := range group {
		total += p.Price
	}
	fmt.Println(category, "subtotal:", ToCurrency(total))
	resultChanel <- total
}
