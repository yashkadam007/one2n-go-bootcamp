package main

import (
	"math"
	"reflect"
	"testing"
)

func TestEvenNumbers(t *testing.T) {
	input := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	evenNumbers := FilterEvenNumbers(input)
	for _, v := range evenNumbers {
		if v%2 != 0 {
			t.Fail()
		}
	}
}

func TestOddNumbers(t *testing.T) {
	input := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	evenNumbers := FilterOddNumbers(input)
	for _, v := range evenNumbers {
		if v%2 == 0 {
			t.Fail()
		}
	}
}

func TestPrimeNumbers(t *testing.T) {
	input := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	primeNumbers := FilterPrimeNumbers(input)
	for _, v := range primeNumbers {
		for i := 2; i < int(math.Sqrt(float64(v))); i++ {
			if v%i == 0 {
				t.Fail()
			}
		}
	}
}

func TestOddPrimeNumbers(t *testing.T) {
	input := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	primeNumbers := FilterOddPrimeNumbers(input)
	for _, v := range primeNumbers {
		for i := 2; i <= int(math.Sqrt(float64(v))); i++ {
			if v%i == 0 && v != 2 {
				t.Errorf("%d divided by %d is not a odd prime number", v, i)
				t.Fail()
			}
		}
	}
}

func TestEvenMultiplesOfFive(t *testing.T) {
	input := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20}
	evenMultiplesOfFive := EvenMultiplesOfFive(input)
	for _, v := range evenMultiplesOfFive {
		if v%2 != 0 || v%5 != 0 {
			t.Fail()
		}
	}
}

func TestOddMultiplesOf3GT10(t *testing.T) {
	input := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20}
	oddMutliplesOf3GT10 := OddMultipleOf3GT10(input)
	for _, v := range oddMutliplesOf3GT10 {
		if v%2 == 0 || v%3 != 0 || v < 10 {
			t.Fail()
		}
	}
}

func TestFilterCondition(t *testing.T) {
	input := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20}
	want := []int{9, 15}
	isOdd := func(n int) bool { return n%2 != 0 }
	isGreaterthan := func(n int) Condition { return func(m int) bool { return m > n } }
	isGreaterthan5 := isGreaterthan(5)
	isMultipleOf := func(n int) Condition { return func(m int) bool { return m%n == 0 } }
	isMultipleOf3 := isMultipleOf(3)
	output := FilterConditions(input, isOdd, isGreaterthan5, isMultipleOf3)
	if !reflect.DeepEqual(want, output) {
		t.Errorf("wanted: %v got:%v", want, output)
	}
}

func TestFilterConditionAny(t *testing.T) {
	input := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20}
	want := []int{1, 2, 3, 4, 5, 6, 9, 12, 15, 18}
	isLessThan := func(n int) Condition { return func(m int) bool { return m < n } }
	isLessThan6 := isLessThan(6)
	isMultipleOf := func(n int) Condition { return func(m int) bool { return m%n == 0 } }
	isMultipleOf3 := isMultipleOf(3)
	output := FilterConditionsAny(input, isLessThan6, isMultipleOf3)
	if !reflect.DeepEqual(want, output) {
		t.Errorf("wanted: %v got:%v", want, output)
	}
}
