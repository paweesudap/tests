package main

import "testing"

func TestFizzBuzzShouldReturn1WhenInput1(t *testing.T) {

	input := 1

	got := FizzBuzz3(input)

	want := "1"
	if got != want {
		t.Errorf("got %q but want %q", got, want)
	}
}
func TestFizzBuzzShouldReturn1WhenInput2(t *testing.T) {

	input := 2

	got := FizzBuzz3(input)

	want := "2"
	if got != want {
		t.Errorf("got %q but want %q", got, want)
	}
}
func TestFizzBuzzShouldReturn1WhenInput3(t *testing.T) {

	input := 3

	got := FizzBuzz3(input)

	want := "fizz"
	if got != want {
		t.Errorf("got %q but want %q", got, want)
	}
}
func TestFizzBuzzShouldReturn1WhenInput4(t *testing.T) {

	input := 4

	got := FizzBuzz3(input)

	want := "4"
	if got != want {
		t.Errorf("got %q but want %q", got, want)
	}
}
func TestFizzBuzzShouldReturn1WhenInput5(t *testing.T) {

	input := 5

	got := FizzBuzz3(input)

	want := "buzz"
	if got != want {
		t.Errorf("got %q but want %q", got, want)
	}
}
func TestFizzBuzzShouldReturn1WhenInput6(t *testing.T) {

	input := 6

	got := FizzBuzz3(input)

	want := "fizz"
	if got != want {
		t.Errorf("got %q but want %q", got, want)
	}
}
func TestFizzBuzzShouldReturn1WhenInput7(t *testing.T) {

	input := 7

	got := FizzBuzz3(input)

	want := "7"
	if got != want {
		t.Errorf("got %q but want %q", got, want)
	}
}
