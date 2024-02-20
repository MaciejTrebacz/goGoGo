package main

import (
	fo "dabbling/fileops"
	"fmt"
	"github.com/Pallinder/go-randomdata"
)

const accountBalanceFile = "balance.txt"

func main() {
	var accountBalance, err = fo.GetFloatFromFile(accountBalanceFile)

	presentOption()

	if err != nil {
		fmt.Println("ERROR")
		fmt.Println(err)
		panic(err)
	}

	// creating loop
	for i := 0; i < 10; i++ {

	}
	const white = 100
	for { // infinite loop

		var choice int
		fmt.Scan(&choice) // number user enters

		// in go switch works that only one statement can be true!
		switch choice {
		case 1:
			fmt.Println("Your balance is  : ", accountBalance)
			fmt.Println(randomdata.PhoneNumber()) // using random data
		case 2:
			fmt.Print("Your deposit: ")
			var depositAmount float64
			fmt.Scan(&depositAmount)
			accountBalance += depositAmount
			fo.WriteFloatToFile(accountBalanceFile, accountBalance)
		case 3:
			fmt.Print("How much to withdraw: ")
			var withdrawAmount float64
			fmt.Scan(&withdrawAmount)
			accountBalance -= withdrawAmount
			fo.WriteFloatToFile(accountBalanceFile, accountBalance)

		case 4:
			fmt.Print("Exiting")
			return // we need to use return as break doesnt work in switch
		default:
			fmt.Print("Goodbye!!!")
		}

		fmt.Println("Your choice: ", choice)

	}
}
