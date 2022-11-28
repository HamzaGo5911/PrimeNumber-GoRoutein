package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	startTime := time.Now()

	wait()
	endTime := time.Now()

	timeDuration := endTime.Sub(startTime).Seconds()
	fmt.Println(timeDuration)

}
func wait() {
	var wg sync.WaitGroup
	routines := 10
	dataSize := 1000
	chunk := dataSize / routines
	start := 0
	end := 0
	for i := 0; i < routines; i++ {
		start = end
		end = start + chunk
		wg.Add(1)
		fmt.Printf("%v -> %v\n", start, end)
		go PrimeNumbersRange(start, end, &wg)
	}
	wg.Wait()
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
