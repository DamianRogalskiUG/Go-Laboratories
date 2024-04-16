package main

import (
	"fmt"
	"math/big"
	"strings"
)

func createNick(name string, lastName string) []int {
	var nick string = name[:3] + lastName[:3]
	var lowerNick string = strings.ToLower(nick)
	var listOfCodes []int
	fmt.Println(lowerNick)
	for _, char := range lowerNick {
		listOfCodes = append(listOfCodes, int(char))
	}
	fmt.Println(listOfCodes)
	return []int{1}
}

func factorial(n int64) *big.Int {
	if n < 0 {
		return big.NewInt(0)
	}
	if n == 0 {
		return big.NewInt(1)
	}
	result := big.NewInt(n)
	return result.Mul(result, factorial(n-1))
}

func checkForAsciiCode(code string)

func findStrongNumber(name string, lastName string) *big.Int {
	result := big.NewInt(1)
	nick := createNick(name, lastName)
	fmt.Println(nick)
	for i := range 101 {

		// fmt.Println(i)
	}
	return result
}

func main() {
	var name string = "Damian"
	var lastName string = "Rogalski"
	// createNick(name, lastName)
	// fmt.Println(factorial(300))
	findStrongNumber(name, lastName)
}
