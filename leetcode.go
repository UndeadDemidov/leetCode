package leetcode

import (
	"strconv"
)

// https://leetcode.com/problems/two-sum
func twoSum(nums []int, target int) []int {
	if nums == nil {
		return nil
	}

	l := len(nums)
	if l < 2 {
		return nil
	}

	for k, fnum := range nums[:l-1] {
		for i, snum := range nums[k+1:] {
			if fnum+snum == target {
				return []int{k, i + k + 1}
			}
		}
	}
	return nil
}

// https://leetcode.com/problems/palindrome-number
func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}
	if x < 10 {
		return true
	}

	bytes := []byte(strconv.Itoa(x))
	l := len(bytes)
	for i := 0; i < l/2; i++ {
		if bytes[i] != bytes[l-i-1] {
			return false
		}
	}
	return true
}

// https://leetcode.com/problems/palindrome-number
func isPalindromeNoString(x int) bool {
	var reversedNum int
	tmp := x
	for tmp > 0 {
		reversedNum = reversedNum*10 + tmp%10
		tmp = tmp / 10
	}
	return x == reversedNum
}

// https://leetcode.com/problems/roman-to-integer
func romanToInt(s string) int {
	r2d := map[rune]int{
		'I': 1,
		'V': 5,
		'X': 10,
		'L': 50,
		'C': 100,
		'D': 500,
		'M': 1000,
	}

	prev := r2d[rune(s[0])]
	result := 0
	dec := 0
	for _, r := range s {
		dec = r2d[r]
		result += dec
		if prev < dec {
			result -= 2 * prev
		}
		prev = dec
	}
	return result
}

// https://leetcode.com/problems/longest-common-prefix
func longestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}
	if len(strs) == 1 {
		return strs[0]
	}
	ln := len(strs[0])
	for _, str := range strs {
		l := len(str)
		if ln > l {
			ln = l
		}
	}

	if ln == 0 {
		return ""
	}

	prefix := ""
	for i := 0; i < ln; i++ {
		char := string(strs[0][i])
		for _, str := range strs[1:] {
			if char != string(str[i]) {
				return prefix
			}
		}
		prefix += char
	}
	return prefix
}

// https://leetcode.com/problems/valid-parentheses/
func isValid(s string) bool {
	stack := make([]rune, 0)
	for _, r := range s {
		if r == '{' || r == '[' || r == '(' {
			stack = append(stack, r)
			continue
		}
		if len(stack) == 0 {
			return false
		}
		if r == '}' && stack[len(stack)-1] == '{' || r == ']' && stack[len(stack)-1] == '[' || r == ')' && stack[len(stack)-1] == '(' {
			stack = stack[:len(stack)-1]
			continue
		}
		return false
	}
	return len(stack) == 0
}

// Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

// https://leetcode.com/problems/merge-two-sorted-lists/
func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	return nil
}
