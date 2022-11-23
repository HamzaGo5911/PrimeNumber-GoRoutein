package main

import (
	"fmt"
	"time"
)

func main() {

	startTime := time.Now()
	fmt.Println(startTime)

	fmt.Println("Prime numbers from a range are:")
	PrimeNumbersRange(3, 100)

	endTime := time.Now()

	timeDuration := endTime.Sub(startTime)
	fmt.Println(timeDuration)

	generated := SumOfPrime(100)
	sum := 0
	for _, val := range generated {
		sum += val
	}
	fmt.Printf("Sum of prime number 1 - 100:  %v \n", sum)
}

func PrimeNumbersRange(num1, num2 int) *int {
	if num1 < 2 || num2 < 2 {

		fmt.Println(" Numbers greater than 2")
		return nil
	}

	var count int
	for num1 <= num2 {
		IsPrime := true

		for i := 2; i <= 2; i++ {
			if IsPrime == true {
				count++
			}
			if num1%i == 0 {
				IsPrime = false
				break
			}
		}
		if IsPrime {
			fmt.Printf("%d ", num1)
		}
		num1++
	}
	return &count
}
func SumOfPrime(n int) []int {

	integers := make([]bool, n+1)

	for i := 2; i < n+1; i++ {
		integers[i] = true
	}

	for p := 2; p*p <= n; p++ {
		if integers[p] == true {

			for i := p * 2; i <= n; i += p {
				integers[i] = false
			}
		}
	}

	var primes []int
	for p := 2; p <= n; p++ {
		if integers[p] == true {
			primes = append(primes, p)
		}
	}

	return primes
}
