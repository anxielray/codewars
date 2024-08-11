package main

import "math"

func main() {
	// println(FindAll())
}

//prints wrong output
func FindAll(sumDig, digs int) []int {
	results := []int{}
	lowest := math.MinInt32
	highest := math.MaxInt32

	// Helper function to generate combinations
	var generate func(start int, current []int, sumDig int)
	generate = func(start int, current []int, sumDig int) {
		if len(current) == digs {
			if sumDig == 0 {
				// Convert current combination to a number
				num := 0
				for _, digit := range current {
					num = num*10 + digit
				}
				results = append(results, num)
				if num < lowest {
					lowest = num
				}
				if num > highest {
					highest = num
				}
			}
			return
		}
		for i := start; i <= 9; i++ {
			if sumDig >= i {
				generate(i, append(current, i), sumDig-i)
			}
		}
	}

	// Generate all valid combinations
	generate(0, []int{}, sumDig)

	if len(results) > 0 {
		return []int{len(results), lowest, highest}
	}

	return nil
}
