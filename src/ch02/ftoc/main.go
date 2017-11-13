package main

import "fmt"

func main() {
	const freezingF, boilingF = 32.0, 212.0
	fmt.Printf("%g°F = %g°C\n", freezingF, f2c(freezingF))
	fmt.Printf("%g°F = %g°C\n", boilingF, f2c(boilingF))
}

func f2c(f float64) float64 {
	return (f - 32) * 5 / 9
}
