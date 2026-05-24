package main

import "fmt"

func main() {
	fmt.Printf("main function started \n")
	CalcStoreTotal(Products)
	fmt.Printf("main function complete")
}

// output
/*
main function started
Soccer subtotal: $54.45
Running subtotal: $79595.00
Watersports subtotal: $323.95
Total: $79973.40
main function complete
*/
