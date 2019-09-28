package coderetreat

import (
	"fmt"
	"testing"
)

func TestNotPalindrome(t *testing.T) {
	if r := isPalindrome("notracecar"); r {
		t.Errorf("isPalindrome(\"notracecar\"): want: false, got: %t", r)
	}
}

func TestIsPalindrome(t *testing.T) {
	if r := isPalindrome("kayak"); !r {
		t.Errorf("isPalindrome(\"kayak\"): want: true, got: %t", r)
	}
}

func TestFindMinEmpty(t *testing.T) {
	if _, err := findMin([]int{}); err == nil {
		t.Errorf("findMin([]int{}): want: nil, got: %v", fmt.Errorf("slice is empty"))
	}
}

func TestFindMin(t *testing.T) {
	if a, err := findMin([]int{9, -5, 1, 4}); err != nil || a != -5 {
		t.Errorf("findMin([]int{9, -5, 1, 4}): want: (5, nil), got: (%d, %v)", a, err)
	}
}
