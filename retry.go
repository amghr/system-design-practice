// Implement a retry library. It must support following retry policy mechanism:
// 1) Fibonacci
// 2) EvenNumber
// 3) OddNumber
//
// I should be able to retry any function

package main

import (
	"errors"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"time"
)

// NumberGenerator is an interface for generating the next number of any sequence
type NumberGenerator interface {
	Next() int
}

// EvenNumberGenerator generates even numbers, starting from a given value
type EvenNumberGenerator struct {
	curr int
}

// NewEvenNumberGenerator initializes the EvenNumberGenerator starting at 'start'
func NewEvenNumberGenerator(start int) *EvenNumberGenerator {
	return &EvenNumberGenerator{curr: start}
}

// Next function here returns the current number and increments the pointer to next even number
func (e *EvenNumberGenerator) Next() int {
	val := e.curr
	e.curr += 2
	return val
}

// OddNumberGenerator generates odd numbers, starting from a given value
type OddNumberGenerator struct {
	curr int
}

// NewOddNumberGenerator initializes the OddNumberGenerator starting at 'start'
func NewOddNumberGenerator(start int) *OddNumberGenerator {
	return &OddNumberGenerator{curr: start}
}

// Next function here returns the current number and increments the pointer to next odd number
func (o *OddNumberGenerator) Next() int {
	val := o.curr
	o.curr += 2
	return val
}

// FibonacciGenerator generates numbers in the Fibonacci sequence
type FibonacciGenerator struct {
	a, b int
}

// NewFibonacciGenerator initializes the FibonacciGenerator with a and b
func NewFibonacciGenerator(a, b int) *FibonacciGenerator {
	return &FibonacciGenerator{a: a, b: b}
}

// Next function here returns the next number in the fibonacci sequence
func (f *FibonacciGenerator) Next() int {
	res := f.a
	f.a, f.b = f.b, f.a+f.b
	return res
}

// RetryDecorator used to wrap the original function with different retry mechanisms
type RetryDecorator struct {
	fn         func() error
	policy     NumberGenerator
	maxRetries int
}

// NewRetryDecorator initializes the retry mechanism
func NewRetryDecorator(fn func() error, policy NumberGenerator, maxRetries int) *RetryDecorator {
	return &RetryDecorator{
		fn:         fn,
		policy:     policy,
		maxRetries: maxRetries,
	}
}

// Retry retries the given function with given retry mechanism for max retries
func (r *RetryDecorator) Retry() error {
	for attempt := 0; attempt <= r.maxRetries; attempt++ {
		err := r.fn()
		if err == nil {
			fmt.Printf("Attempt %d: Success\n", attempt+1)
			return nil
		}

		if attempt < r.maxRetries {
			wait := r.policy.Next()
			fmt.Printf("Attempt %d failed: %v. Retrying in %dms...\n", attempt+1, err, wait)
			time.Sleep(time.Duration(wait) * time.Millisecond)
		} else {
			return fmt.Errorf("all %d attempts failed: %w", r.maxRetries+1, err)
		}
	}
	return nil
}

// sampleFunction simulates a random failure. This could be replaced with http or db calls
func sampleFunction() error {
	if rand.Intn(3) != 0 {
		return errors.New("simulated failure")
	}
	return nil
}

// getRetryPolicy switches between different retry policies
func getRetryPolicy(policyName string) NumberGenerator {
	switch policyName {
	case "even":
		return NewEvenNumberGenerator(0)
	case "odd":
		return NewOddNumberGenerator(1)
	case "fib":
		return NewFibonacciGenerator(0, 1)
	default:
		fmt.Println("Invalid policy name. Use 'even', 'odd', or 'fib'")
		os.Exit(1)
	}
	return nil
}

func main() {
	if len(os.Args) < 3 {
		fmt.Println("Usage: go run main.go <even|odd|fib> <maxRetries>")
		return
	}

	policyName := os.Args[1]
	maxRetries, err := strconv.Atoi(os.Args[2])
	if err != nil || maxRetries < 1 {
		fmt.Println("Invalid maxRetries value. Must be a positive integer.")
		return
	}

	retryPolicy := getRetryPolicy(policyName)
	decorator := NewRetryDecorator(sampleFunction, retryPolicy, maxRetries)

	if err := decorator.Retry(); err != nil {
		fmt.Println("Final result:", err)
	} else {
		fmt.Println("Operation succeeded within retry limit.")
	}
}
