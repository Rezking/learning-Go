package main

import "fmt"

func half(x int) (int, bool) {
	if x%2 == 0 {
		return 1, true
	} else {
		return 0, false
	}
}

func largest(args ...int) int {
	max := 0
	for _, value := range args {
		if max < value {
			max = value
		} else {
			continue
		}
	}
	return max
}

func fibonacci(x int) int {
	if x == 0 {
		return 0
	} else if x == 1 {
		return 1
	} else {
		return fibonacci(x-1) + fibonacci(x-2)
	}
}

func fib(x int) []int {
	slice := []int{}
	for i := 0; i <= x; i++ {
		value := fibonacci(i)
		slice = append(slice, value)
	}
	return slice
}

func makeOddgenerator() func() uint {
	i := uint(1)
	return func() (ret uint) {
		ret = i
		i += 2
		return
	}
}

func main() {
	fmt.Println(half(4))
	fmt.Println(largest(2, 3, 4, 12, 5)) // 12
	nextOdd := makeOddgenerator()
	fmt.Println(nextOdd()) // prints 1
	fmt.Println(nextOdd()) //prints 3
	fmt.Println(fib(6))
}
