package main

import "fmt"

func isPrimeNo(num int) bool {
	if num <= 1 {
		return false
	}
	for i := 2; i*i <= num; i++ {
		if num%i == 0 {
			return false
		}
	}
	return true
}

func main() {
	sum := 0
	fmt.Println("The Prime numbers between 1 and 10:")

	for i := 1; i <= 10; i++ {
		if isPrimeNo(i) {
			fmt.Println(i)
			sum += i
		}
	}

	fmt.Printf("The Sum of the prime numbers between 1 and 10 is : %d\n", sum)
}
