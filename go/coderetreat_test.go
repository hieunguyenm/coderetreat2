package coderetreat

import (
	"fmt"
	"testing"
)

func TestNotPalindrome(t *testing.T) {
	if isPalindrome("notracecar") {
		t.Errorf("isPalindrome(\"notracecar\"): got: true, want: false")
	}
}

func TestIsPalindrome(t *testing.T) {
	if !isPalindrome("kayak") {
		t.Errorf("isPalindrome(\"kayak\"): got: false, want: true")
	}
}

func TestFindMinEmpty(t *testing.T) {
	if _, err := findMin([]int{}); err == nil {
		t.Errorf("findMin([]int{}): got: nil, want: %v", fmt.Errorf("slice is empty"))
	}
}

func TestFindMin(t *testing.T) {
	if a, err := findMin([]int{9, -5, 1, 4}); err != nil || a != -5 {
		t.Errorf("findMin([]int{9, -5, 1, 4}): got: %d, %v, want: %d, %v", a, err, -5, nil)
	}
}
