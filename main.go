package main

import "fmt"

func fibonacci() func() int {
	x, y := 0, 1
	return func() int {
		x, y = y, x+y
		return x
	}
}

func main() {
	f := fibonacci()
	max := 4000000
	var sum int
	for {
		f := f()
		if f > max {
			fmt.Println(sum)
			break
		}
		if f%2 == 0 {
			sum += f
		}
	}
}
