package main

import (
	"fmt"
	"math"
	"math/big"
	"strconv"
	"strings"
	"time"
)

// createNick creates a nickname from the first 3 letters of the name and last name
func createNick(name, lastName string) []int {
	nick := name[:3] + lastName[:3]
	lower := strings.ToLower(nick)
	var codes []int
	for _, char := range lower {
		codes = append(codes, int(char))
	}
	return codes
}

// factorial calculates the factorial of a number
func factorial(n int64) *big.Int {
	if n < 0 {
		return big.NewInt(0)
	} else if n == 0 {
		return big.NewInt(1)
	}
	result := big.NewInt(n)
	return result.Mul(result, factorial(n-1))
}

// containsAllDigits checks if the number contains all digits from the nickname codes
func containsAllDigits(number *big.Int, nickCodes []int) bool {
	numberStr := number.String()
	count := 0
	for _, digit := range nickCodes {
		digitStr := strconv.Itoa(digit)
		if strings.Contains(numberStr, digitStr) {
			numberStr = strings.Replace(numberStr, digitStr, "", 1)
			count++
		}
	}
	return count == len(nickCodes)
}

// findStrongNumber finds the first number for which the factorial contains all digits from the nickname
func findStrongNumber(name, lastName string) int {
	result := 0
	nickCodes := createNick(name, lastName)
	for i := int64(1); i <= 500; i++ {
		num := factorial(i)
		if containsAllDigits(num, nickCodes) {
			result = int(i)
			break
		}
	}
	return result
}

// findWeakNumber finds the number for which the number of calls of the Fibonacci function is the closest to the strong number
func findWeakNumber(fibNumber, strongNumber int) int {
	fmt.Printf("Fibonacci(%d) = %d\n", fibNumber, fibonacci(fibNumber))
	weakNumber := 0
	minDifference := math.MaxInt64

	for key, value := range callCount {
		diff := int(math.Abs(float64(value - strongNumber)))
		fmt.Printf("Calls for n=%d: %d\n", key, value)
		if diff < minDifference {
			minDifference = diff
			weakNumber = key
		}
	}

	return weakNumber
}

// fibonacci calculates the Fibonacci number for a given n
func fibonacci(n int) int {
	callCount[n]++
	if n <= 1 {
		return n
	}
	return fibonacci(n-1) + fibonacci(n-2)
}

var callCount map[int]int

func init() {
	callCount = make(map[int]int)
}

func main() {
	name := "Damian"
	lastName := "Rogalski"

	strongNumber := findStrongNumber(name, lastName)
	fmt.Println("Strong number:", strongNumber)

	fibNumber := 30
	weakNumber := findWeakNumber(fibNumber, strongNumber)
	fmt.Println("Weak number:", weakNumber)

	startTime := time.Now()
	result := fibonacci(30)
	endTime := time.Now()

	fmt.Printf("Wynik: %d\n", result)
	fmt.Printf("Czas wykonania: %v\n", endTime.Sub(startTime))
}
