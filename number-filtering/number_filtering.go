package main

import (
	"fmt"
	"math"
)

func isPrime(n int) bool {
	for i := 2; i < int(math.Sqrt(float64(n))); i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}

func FilterEvenNumbers(arr []int) []int {
	var result []int
	for i := 0; i < len(arr); i++ {
		if arr[i]%2 == 0 {
			result = append(result, arr[i])
		}
	}
	return result
}

func FilterOddNumbers(arr []int) []int {
	var result []int
	for _, v := range arr {
		if v%2 != 0 {
			result = append(result, v)
		}
	}
	return result
}

func FilterPrimeNumbers(arr []int) []int {
	var result []int
	for _, v := range arr {
		if v < 2 {
			continue
		}
		isPrime := true
		for i := 2; i <= int(math.Sqrt(float64(v))); i++ {
			if v%i == 0 {
				isPrime = false
				break
			}
		}
		if isPrime {
			result = append(result, v)
		}
	}
	return result
}

func FilterOddPrimeNumbers(arr []int) []int {
	var result []int
	for _, v := range arr {
		if v < 3 {
			continue
		}
		isPrime := true
		for i := 2; i <= int(math.Sqrt(float64(v))); i++ {
			if v%i == 0 {
				isPrime = false
				break
			}
		}
		if isPrime {
			result = append(result, v)
		}
	}
	return result
}

func EvenMultiplesOfFive(arr []int) []int {
	var result []int
	for _, v := range arr {
		if v%2 == 0 && v%5 == 0 {
			result = append(result, v)
		}
	}
	return result
}

func OddMultipleOf3GT10(arr []int) []int {
	var result []int
	for _, v := range arr {
		if v%3 == 0 && v%2 != 0 && v > 10 {
			result = append(result, v)
		}
	}
	return result
}

type Condition func(n int) bool

func FilterConditions(arr []int, conditions ...Condition) []int {
	var result []int
	for _, v := range arr {
		matchingCondition := true
		for _, cond := range conditions {
			if !cond(v) {
				matchingCondition = false
			}
		}
		if matchingCondition {
			result = append(result, v)
		}
	}
	return result
}

func FilterConditionsAny(arr []int, conditions ...Condition) []int {
	var result []int
	for _, v := range arr {
		matchingCondition := false
		for _, cond := range conditions {
			if cond(v) {
				matchingCondition = true
			}
		}
		if matchingCondition {
			result = append(result, v)
		}
	}
	return result
}

func main() {
	story1Input := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	fmt.Println(FilterEvenNumbers(story1Input))
	fmt.Println(FilterOddNumbers(story1Input))
	fmt.Println(FilterPrimeNumbers(story1Input))
	fmt.Println(FilterOddPrimeNumbers(story1Input))
	story5Input := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20}
	fmt.Println(EvenMultiplesOfFive(story5Input))
	fmt.Println(OddMultipleOf3GT10(story5Input))
}
