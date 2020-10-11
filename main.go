package main

import "fmt"

func main() {
	// Variables.
	var (
		s []int
		res int
		value int
		length int
	)

	// Length of slices
	_, _ = fmt.Scan(&length)

	// Insert data.
	for i := 0; i < length; i++ {
		_, _ = fmt.Scan(&value)
		s = append(s, value)
	}

	// Sum.
	for _, val := range s {
		res += val
	}

	// Result
	fmt.Println(res)

	// Print array.
	// fmt.Println(s)

	// Print length and capacity.
	// fmt.Println(len(s), cap(s))
}