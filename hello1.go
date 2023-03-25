package main

import (
	"fmt"
	"math"
)

// arrays, slices, maps, control flow, functions
type Circle struct {
	x, y, r float64
}

type Shapes interface {
	area() float64
}

type Rectangle struct {
	l, w float64
}

func square(x *int) {
	*x = *x * *x
}

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

func swap(x *int, y *int) {
	new := *x
	var ptrx *int = x
	var ptry *int = y
	*ptrx = *y
	*ptry = new
}
func makeOddgenerator() func() uint {
	i := uint(1)
	return func() (ret uint) {
		ret = i
		i += 2
		return
	}
}

// working with methods
func (c *Circle) area() float64 {
	return math.Pi * c.r * c.r
}

func (r *Rectangle) area() float64 {
	return r.l * r.w
}

func totalArea(shapes ...Shapes) float64 {
	var totalarea float64
	for _, s := range shapes {
		totalarea += s.area()
	}
	return totalarea
}

func main() {
	// x := 2
	// y := 6
	// swap(&x, &y)
	// fmt.Println(x)
	// fmt.Println(y)
	c := Circle{0, 0, 5}
	fmt.Println(c.area())
	fmt.Println(c.x)
	// square(&x)
	// fmt.Println(x)
	// fmt.Println(half(4))
	// fmt.Println(largest(2, 3, 4, 12, 5)) // 12
	// nextOdd := makeOddgenerator()
	// fmt.Println(nextOdd()) // prints 1
	// fmt.Println(nextOdd()) //prints 3
	// fmt.Println(fib(6))

}
