package cicalculator

import (
	"fmt"
	"math"
)

func Calculate() {

	var investmentAmount float64
	var returnRate float64
	var years int
	fmt.Println("-------------------------------------")
	fmt.Print("Amount Deposited: ")
	fmt.Scan(&investmentAmount)
	fmt.Print("Interest Rate: ")
	fmt.Scan(&returnRate)
	fmt.Print("Time Period: ")
	fmt.Scan(&years)

	var futureValue = investmentAmount * math.Pow((1+returnRate/100), float64(years))
	fmt.Printf("Amount for compound interest: %.2f", futureValue)

}
