package main

import "fmt"

func main() {
	age := 32

	agePointer := &age

	fmt.Println("Age: ", age, &age)
	editAgeToAdultYears(&age)
	fmt.Println("Age: ", age, &age)
	fmt.Println(getAdultYears(age))
	fmt.Printf("Age pointer address: %v \nAge pointer value: %v \n", agePointer, *agePointer)
}

func getAdultYears(age int) int {
	return age - 18
}

func editAgeToAdultYears(age *int) {
	*age -= 18
}
