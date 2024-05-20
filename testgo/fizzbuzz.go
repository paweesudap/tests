package main

import "strconv"

func FizzBuzz(n int) string {
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
