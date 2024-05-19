package main

import "testing"

func TestFizzBuzzShouldReturn1WhenInput1(t *testing.T) {

	input := 1

	got := FizzBuzz(input)

	want := "1"
	if got != want {
		t.Errorf("got %q but want %q", got, want)
	}
}
func TestFizzBuzzShouldReturn1WhenInput2(t *testing.T) {

	input := 2

	got := FizzBuzz(input)

	want := "2"
	if got != want {
		t.Errorf("got %q but want %q", got, want)
	}
}
func TestFizzBuzzShouldReturn1WhenInput3(t *testing.T) {

	input := 3

	got := FizzBuzz(input)

	want := "fizz"
	if got != want {
		t.Errorf("got %q but want %q", got, want)
	}
}
