package main

import "fmt"

func main() {
	var accountBalance float64 = 1000
	// creating loop
	for i := 0; i < 10; i++ {

	}

	const white = 100
	for { // infinite loop

		fmt.Println("Welcome to GO BANK...")
		fmt.Println("What you want to do...")
		fmt.Println("1. Check balance")
		fmt.Println("2. Deposit money")
		fmt.Println("3. Withdraw money")
		fmt.Println("4. Exit")

		var choice int
		fmt.Scan(&choice) // number user enters

		// in go switch works that only one statement can be true!
		switch choice {
		case 1:
			fmt.Println("Your balance is  : ", accountBalance)
		case 2:
			fmt.Print("Your deposit: ")
			var depositAmount float64
			fmt.Scan(&depositAmount)
			accountBalance += depositAmount
		case 3:
			fmt.Print("How much to withdraw: ")
			var withdrawAmount float64
			fmt.Scan(&withdrawAmount)
			accountBalance -= withdrawAmount
		case 4:
			fmt.Print("Exiting")
			return // we need to use return as break doesnt work in switch
		default:
			fmt.Print("Goodbye!!!")
		}

		fmt.Println("Your choice: ", choice)

	}
}
