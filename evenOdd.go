// Problem statement: Write a program to print the first N even and odd numbers.
// User of this program should be able to print either even or odd numbers (think of passing even/odd option as a command line argument).
// So, if option selected is 'even' and N = 5, program should print 0, 2, 4, 6, 8.

//Remember it is a design problem. Use any OO programming language (including GoLang).

// ------------- I have used Interface Segration and Single Responsibility SOLID principles to implement this ------------ //

package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
)

// name can be better
type EvenOdd interface {
	EvenOrOdd(n int) []int
}

type Even struct{}

func (e Even) EvenOrOdd(n int) []int {
	nums := make([]int, n)
	for i := 0; i < n; i++ {
		nums[i] = 2 * i
	}

	return nums
}

type Odd struct{}

func (o Odd) EvenOrOdd(n int) []int {
	nums := make([]int, n)
	for i := 0; i < n; i++ {
		nums[i] = 2*i + 1
	}

	return nums
}

type Fib struct{}

func printEvenOrOdd(eo EvenOdd, n int) {
	result := eo.EvenOrOdd(n)
	fmt.Printf("The numbers are: %v\n", result)
}

func main() {
	if len(os.Args) < 3 {
		fmt.Println("Usage: go run evenOdd.go even|odd N")
		os.Exit(1)
	}

	mode := os.Args[1]

	n, err := strconv.Atoi(os.Args[2])
	if err != nil || n < 1 {
		log.Fatalf("Error: '%v' is not a valid number of elements.", os.Args[2])
	}

	var eo EvenOdd
	switch mode {
	case "even":
		eo = Even{}
	case "odd":
		eo = Odd{}
	default:
		log.Fatal("Invalid option. Please specify 'even' or 'odd'.")
	}

	printEvenOrOdd(eo, n)
}

// --------------------- Below is a straight forward approach I wrote first before using any SOLID principles --------------------- //

//package main
//
//import (
//"fmt"
//"log"
//)
//
//func main() {
//	var n, o int
//
//	fmt.Println("Enter the number N.")
//	_, err := fmt.Scan(&n)
//	if err != nil {
//		log.Fatal(err)
//	}
//
//	if n < 0 {
//		log.Fatal("You must enter a positive integer.")
//	}
//
//	fmt.Println("Enter 0 for even and 1 for odd")
//	_, err = fmt.Scan(&o)
//	if err != nil {
//		log.Fatal(err)
//	}
//
//	nums := evenOrOdd(n, o)
//	fmt.Println(nums)
//
//}
//
//func evenOrOdd(n, o int) []int {
//	nums := make([]int, n)
//
//	if o == 0 {
//		for i := 0; i < n; i++ {
//			nums[i] = 2 * i
//		}
//	} else {
//		for i := 0; i < n; i++ {
//			nums[i] = 2*i + 1
//		}
//	}
//
//	return nums
//}
