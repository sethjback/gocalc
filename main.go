package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	if len(os.Args) < 4 {
		fmt.Println("must specify operation and numbers")
		os.Exit(1)
	}

	switch os.Args[1] {
	case "add":
		ai, err := convertArgs(os.Args[2:])
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		a := NewAdder(ai[0], ai[1])

		fmt.Println(a.Calculate())
	case "sub":
		ai, err := convertArgs(os.Args[2:])
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		s := NewSubber(ai[0], ai[1])

		fmt.Println(s.Calculate())
	default:
		fmt.Println("valid operations are add and sub")
		os.Exit(1)
	}

}

func convertArgs(args []string) ([]int, error) {
	ai := []int{}

	for _, a := range args {
		c, err := strconv.Atoi(a)
		if err != nil {
			return nil, err
		}

		ai = append(ai, c)
	}

	return ai, nil
}

// Adder is for basic addition
type Adder struct {
	num1 int // first number to add
	num2 int // second number to add
}

func NewAdder(a, b int) *Adder {
	return &Adder{num1: a, num2: b}
}

func (a *Adder) Calculate() int {
	return a.num1 + a.num2
}

// Subber is for basic subtraction
type Subber struct {
	num1 int // first number to add
	num2 int // second number to add
}

func NewSubber(a, b int) *Subber {
	return &Subber{num1: a, num2: b}
}

func (a *Subber) Calculate() int {
	return a.num1 - a.num2
}
