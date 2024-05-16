package main

import (
	"fmt"
)

func main() {
	var msg string
	var cardType string

	// test Amex card
	cardNum := "30569309025904"

	// check if the card is valid and set the message as valid or invalid
	if checkLuhn(cardNum) {
		msg = "valid"
	} else {
		msg = "invalid"
	}

	// check which type of card is being used
	switch cardNum[0] - '0' {
	case 3:
		cardType = "Amex card"
	case 4:
		cardType = "Visa card"
	case 2, 5:
		cardType = "Mastercard"
	case 6:
		cardType = "Discover card"
	case 7:
		cardType = "Gas card"
	case 8:
		cardType = "Healthcare and telecommunications card"
	case 9:
		cardType = "Governmental card"
	}

	fmt.Printf("The provided %s is %s", cardType, msg)
}

func checkLuhn(s string) bool {
	var sum int
	var isSecond bool

	// loop through the string
	for i := len(s) - 1; i >= 0; i-- {

		// represent the number string into it's integer value
		d := s[i] - '0'

		// double every second number
		if isSecond {
			d *= 2
		}

		// add two digits in the cases that make two digits after doubling the number
		sum += int(d) / 10
		sum += int(d) % 10

		// Updating the is second value to true then false until the end of the loop
		isSecond = !isSecond

	}

	// return true or false
	return (sum%10 == 0)
}
