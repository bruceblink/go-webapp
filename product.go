package main

import "strconv"

type Product struct {
	Name, Category string
	Price          float64
}

var products = []*Product{
	{"Kayak", "Watersports", 275},
	{"Lifejacket", "Watersports", 48.95},
	{"Soccer Ball", "Soccer", 19.50},
	{"Corner Flags", "Soccer", 34.95},
	{"Running Shoes", "Running", 95.00},
	{"Stadium", "Running", 79500.00},
}

type ProductGroup []*Product

type ProductData = map[string]ProductGroup

var Products = make(ProductData)

func ToCurrency(val float64) string {
	return "$" + strconv.FormatFloat(val, 'f', 2, 64)
}

func init() {
	for _, p := range products {
		Products[p.Category] = append(Products[p.Category], p)

	}
}
