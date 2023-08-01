package main

import (
	"errors"
	"fmt"
	"os"
	"strconv"
)

func main() {
	if len(os.Args) < 4 {
		fmt.Println("must specify operation and numbers")
		os.Exit(1)
	}

	fmt.Println(os.Args)

	var op Operation

	switch os.Args[1] {
	case "add":
		op = &Adder{}
	case "sub":
		op = &Subber{}
	default:
		fmt.Println("valid operations are add and sub")
		os.Exit(1)
	}

	ai, err := convertArgs(os.Args[2:])
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	err = op.Args(ai)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Println(op.Calculate())
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

type Operation interface {
	Args([]int) error
	Calculate() int
}

// Adder is for basic addition
type Adder struct {
	num1 int // first number to add
	num2 int // second number to add
}

func NewAdder(a, b int) *Adder {
	return &Adder{num1: a, num2: b}
}

func (a *Adder) Args(numbers []int) error {
	if len(numbers) > 2 {
		return errors.New("we only support adding 2 numbers")
	}
	a.num1 = numbers[0]
	a.num2 = numbers[1]

	return nil
}

func (a *Adder) Calculate() int {
	return a.num1 + a.num2
}

// Subber is for basic subtraction
type Subber struct {
	nums []int // numbers to subtract
}

func NewSubber(numbers []int) *Subber {
	return &Subber{nums: numbers}
}

func (s *Subber) Args(numbers []int) error {
	s.nums = numbers
	return nil
}

func (s *Subber) Calculate() int {
	var ans int
	for i, n := range s.nums {
		if i == 1 {
			ans = n
			continue
		}
		ans -= n
	}

	return ans
}
