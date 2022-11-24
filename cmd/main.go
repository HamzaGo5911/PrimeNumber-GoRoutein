package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var waitgroup sync.WaitGroup

	startTime := time.Now()
	waitgroup.Add(3)

	go PrimeNumbersRange(3, 40, &waitgroup)
	go PrimeNumbersRange(41, 70, &waitgroup)
	go PrimeNumbersRange(71, 100, &waitgroup)
	waitgroup.Wait()
	endTime := time.Now()

	timeDuration := endTime.Sub(startTime).Seconds()
	fmt.Println(timeDuration)

}

func PrimeNumbersRange(num1, num2 int, waitgroup *sync.WaitGroup) *int {
	defer waitgroup.Done()
	sum := 0
	count := 0
	if num1 < 2 || num2 < 2 {
		fmt.Println(" Numbers greater than 2")
		return nil
	}

	for num1 <= num2 {
		IsPrime := true
		for i := 2; i <= num1/2; i++ {
			if num1%i == 0 {
				IsPrime = false
				break
			}
		}
		if IsPrime {
			fmt.Printf("%d ", num1)
			sum = sum + num1
			count++
		}
		num1++
	}

	fmt.Println("\nsum", sum)
	fmt.Println("count", count)
	return nil
}
