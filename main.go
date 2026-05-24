package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Printf("main function started \n")
	CalcStoreTotal(Products)
	time.Sleep(time.Second * 5)
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

/* after add goroutine in CalcStoreTotal function
main function started
Total: $0.00
main function complete
*/

/* wait goroutine in main function
main function started
Total: $0.00
Running subtotal: $79595.00
Soccer subtotal: $54.45
Watersports subtotal: $323.95
main function complete
*/
