package main

import "fmt"

func GuessMyNumber(game Game) string {
	index := 1
	arr := [1000]int{}

	for {
		for i := 0; i < 1000; i++ {
			arr[i] = index
			index++
		}
		res := BinarySearch(arr, game)
		if res == -1 {
			continue
		}
		return fmt.Sprintf("Your Number was %d", arr[res])
	}
}

func BinarySearch(a [1000]int, g Game) int {
	r := -1 // not found
	start := 0
	end := len(a) - 1
	for start <= end {
		mid := (start + end) / 2
		res := g.CheckNumber(a[mid])
		if res == "CORRECT" {
			r = mid // found
			break
		} else if res == "My Number is Greater" {
			start = mid + 1
		} else if res == "My Number is Lower" {
			end = mid - 1
		}
	}
	return r
}
