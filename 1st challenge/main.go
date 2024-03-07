package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	input := bufio.NewScanner(os.Stdin)

	// Read the number of credit card numbers
	input.Scan()
	num, _ := strconv.Atoi(input.Text())

	// Read each credit card number and check it's valid or invalid
	for i := 0; i < num && input.Scan(); i++ {
		card := input.Text()
		if checkCardNumber(card) {
			fmt.Println("Valid")
		} else {
			fmt.Println("Invalid")
		}
	}
}

func checkCardNumber(card string) bool {
	// Check if the card starts with 4, 5, or 6
	if !(strings.HasPrefix(card, "4") || strings.HasPrefix(card, "5") || strings.HasPrefix(card, "6")) {
		return false
	}

	// Remove hyphens and check if the length is 16
	card = strings.ReplaceAll(card, "-", "")
	if len(card) != 16 {
		return false
	}

	// Check if the card contains only digits
	match, _ := regexp.MatchString("^[0-9]+$", card)
	if !match {
		return false
	}

	// Check if there are no more than 3 consecutive repeating digits
	for i := 0; i < len(card)-3; i++ {
		if card[i] == card[i+1] && card[i] == card[i+2] && card[i] == card[i+3] {
			return false
		}
	}

	return true
}
