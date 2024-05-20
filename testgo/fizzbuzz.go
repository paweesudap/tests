package main

import "strconv"

func FizzBuzz(n int) any {
	if n%15 == 0 {
		return "fizzbuzz"
	} else if n%3 == 0 {
		return "fizz"
	} else if n%5 == 0 {
		return "buzz"
	}
	str := strconv.Itoa(n)
	return str
}

func FizzBuzz2(n int) string {
	fizzBuzz := make(map[int]string)
	fizzBuzz[3] = "fizz"
	fizzBuzz[5] = "buzz"
	fizzBuzz[15] = "fizzbuzz"

	result := fizzBuzz[n]
	if result == "" {
		return strconv.Itoa(n)
	}
	return fizzBuzz[n]
}
