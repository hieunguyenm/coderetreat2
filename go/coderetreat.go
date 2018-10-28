package coderetreat

import (
	"fmt"
	"sort"
)

func isPalindrome(a string) bool {
	r := []rune(a)
	for i := 0; i < len(r)/2; i++ {
		tmp := r[i]
		r[i] = r[len(r)-i-1]
		r[len(r)-i-1] = tmp
	}
	return a == string(r)
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
