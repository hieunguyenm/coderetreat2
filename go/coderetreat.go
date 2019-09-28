package coderetreat

import (
	"fmt"
	"sort"
)

func isPalindrome(a string) bool {
	for i := 0; i < len(a)/2; i++ {
		if a[i] != a[len(a)-i-1] {
			return false
		}
	}
	return true
}

func findMin(a []int) (int, error) {
	if len(a) == 0 {
		return 0, fmt.Errorf("slice is empty")
	}
	sort.Slice(a, func(i, j int) bool {
		return a[i] < a[j]
	})
	return a[0], nil
}
