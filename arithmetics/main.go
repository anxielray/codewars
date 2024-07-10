package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

func main() {
	fmt.Println(SumOfDivided([]int{15, 21, 24, 30, 45}))
}

func SumOfDivided(lst []int) string {
	var (
		final       [][]int
		prime       []int
		ans         string
		finalAnswer string
	)
	for x := 2; x < 10; x++ {
		if isPrime(x) {
			prime = append(prime, x)
		}
	}
	for _, prm := range prime {
		final = append(final, core(prm, lst))
	}
	var str string
	start := "("
	for _, array := range final {
		begin := "("
		str += strconv.Itoa(array[0]) + " " + strconv.Itoa(array[1])
		ans += begin + str + ") "
		str = ""
		continue
	}
	finalAnswer = start + strings.TrimSpace(ans) + ")"
	return finalAnswer
}

func core(prime int, lst []int) []int {
	var (
		first []int
	)
	first = append(first, prime)
	for _, num := range lst {
		if num%prime == 0 {
			for _, restOfNum := range lst {
				if (restOfNum%prime == 0) && (restOfNum != num) {
					num += restOfNum
				}
			}
			first = append(first, num)
			break
		}
	}
	return first
}

func isPrime(n int) bool {
	if n < 2 {
		return false
	}
	for i := 2; i <= int(math.Sqrt(float64(n))); i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}
