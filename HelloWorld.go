package main

import (
	"conf2022_the_best_in_the_tests_templates_go/input"
	"os"
)

func main() {
	file, err := os.Open("local.csv")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	validation := input.Validate(file.Name())

	validation.PrintErrors()
}
