package main

import (
	"fmt"
	"math"
)

const (
	epsilon = 0.001
)

// sqrt calculates a square root by guessing
// sqrt is defined as y^2 = x; y = x / y {y != 0}
// This means that if abs(y - x/y) is pretty close then y is pretty close to the square root of x
func sqrt(x float64) float64 {
	return try(x, 1)
}

// If the guess is close enough to the actual value, we're done.
// Otherwise come up with a better guess and try again.
func try(number, guess float64) float64 {
	if closeEnough(number/guess, guess) {
		return guess
	}
	return try(number, betterGuess(number, guess))
}

// Are two numbers within some epsilon?
func closeEnough(a, b float64) bool {
	return math.Abs(a-b) < epsilon
}

// Using these facts:
//   if guess > actual square root of number:
//       then number / guess < actual
//   if guess < actual
//       then number / guess > actual
// Which means we want a value between the two which is why we divide by two.
func betterGuess(number, guess float64) float64 {
	return (guess + number/guess) / 2
}

func main() {
	fmt.Println(sqrt(9))
}
