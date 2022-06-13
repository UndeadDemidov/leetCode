package leetcode

import (
	"reflect"
	"testing"
)

func Test_twoSum(t *testing.T) {
	type args struct {
		nums   []int
		target int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "example 1",
			args: args{
				nums:   []int{2, 7, 11, 15},
				target: 9,
			},
			want: []int{0, 1},
		},
		{
			name: "example 2",
			args: args{
				nums:   []int{3, 2, 4},
				target: 6,
			},
			want: []int{1, 2},
		},
		{
			name: "example 3",
			args: args{
				nums:   []int{3, 3},
				target: 6,
			},
			want: []int{0, 1},
		},
		{
			name: "example 4",
			args: args{
				nums:   []int{2, 7, 11, 15},
				target: 22,
			},
			want: []int{1, 3},
		},
		{
			name: "example 5",
			args: args{
				nums:   []int{2, 7, 11, 15},
				target: 26,
			},
			want: []int{2, 3},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := twoSum(tt.args.nums, tt.args.target); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("twoSum() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_isPalindrome(t *testing.T) {
	type args struct {
		x int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "valid 5132315",
			args: args{x: 5132315},
			want: true,
		},
		{
			name: "valid 5002005",
			args: args{x: 5002005},
			want: true,
		},
		{
			name: "valid 13231",
			args: args{x: 13231},
			want: true,
		},
		{
			name: "valid 121",
			args: args{x: 121},
			want: true,
		},
		{
			name: "valid 11",
			args: args{x: 11},
			want: true,
		},
		{
			name: "valid 0",
			args: args{x: 0},
			want: true,
		},
		{
			name: "invalid -121",
			args: args{x: -121},
			want: false,
		},
		{
			name: "invalid 10",
			args: args{x: 10},
			want: false,
		},
		{
			name: "invalid 1000021",
			args: args{x: 1000021},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isPalindrome(tt.args.x); got != tt.want {
				t.Errorf("isPalindrome() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_isPalindromeNoString(t *testing.T) {
	type args struct {
		x int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "valid 5132315",
			args: args{x: 5132315},
			want: true,
		},
		{
			name: "valid 5002005",
			args: args{x: 5002005},
			want: true,
		},
		{
			name: "valid 13231",
			args: args{x: 13231},
			want: true,
		},
		{
			name: "valid 121",
			args: args{x: 121},
			want: true,
		},
		{
			name: "valid 11",
			args: args{x: 11},
			want: true,
		},
		{
			name: "valid 0",
			args: args{x: 0},
			want: true,
		},
		{
			name: "invalid -121",
			args: args{x: -121},
			want: false,
		},
		{
			name: "invalid 10",
			args: args{x: 10},
			want: false,
		},
		{
			name: "invalid 1000021",
			args: args{x: 1000021},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isPalindromeNoString(tt.args.x); got != tt.want {
				t.Errorf("isPalindromeNoString() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_romanToInt(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "I",
			args: args{s: "I"},
			want: 1,
		},
		{
			name: "III",
			args: args{s: "III"},
			want: 3,
		},
		{
			name: "IV",
			args: args{s: "IV"},
			want: 4,
		},
		{
			name: "V",
			args: args{s: "V"},
			want: 5,
		},
		{
			name: "VI",
			args: args{s: "VI"},
			want: 6,
		},
		{
			name: "VIII",
			args: args{s: "VIII"},
			want: 8,
		},
		{
			name: "IX",
			args: args{s: "IX"},
			want: 9,
		},
		{
			name: "LVIII",
			args: args{s: "LVIII"},
			want: 58,
		},
		{
			name: "MCMXCIV",
			args: args{s: "MCMXCIV"},
			want: 1994,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := romanToInt(tt.args.s); got != tt.want {
				t.Errorf("romanToInt() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_longestCommonPrefix(t *testing.T) {
	type args struct {
		strs []string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "strs = [\"flower\",\"flow\",\"flight\"]",
			args: args{strs: []string{"flower", "flow", "flight"}},
			want: "fl",
		},
		{
			name: "strs = [\"dog\",\"racecar\",\"car\"]",
			args: args{strs: []string{"dog", "racecar", "car"}},
			want: "",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := longestCommonPrefix(tt.args.strs); got != tt.want {
				t.Errorf("longestCommonPrefix() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_isValid(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "()",
			args: args{s: "()"},
			want: true,
		},
		{
			name: "()[]{}",
			args: args{s: "()[]{}"},
			want: true,
		},
		{
			name: "([]{})",
			args: args{s: "([]{})"},
			want: true,
		},
		{
			name: "(]",
			args: args{s: "(]"},
			want: false,
		},
		{
			name: "[",
			args: args{s: "["},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isValid(tt.args.s); got != tt.want {
				t.Errorf("isValid() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_mergeTwoLists(t *testing.T) {
	type args struct {
		list1 *ListNode
		list2 *ListNode
	}
	tests := []struct {
		name string
		args args
		want *ListNode
	}{
		{
			name: "two simple lists",
			args: args{
				list1: &ListNode{
					Val: 1,
					Next: &ListNode{
						Val: 2,
						Next: &ListNode{
							Val:  4,
							Next: nil,
						},
					},
				},
				list2: &ListNode{
					Val: 1,
					Next: &ListNode{
						Val: 3,
						Next: &ListNode{
							Val:  4,
							Next: nil,
						},
					},
				},
			},
			want: &ListNode{
				Val: 1,
				Next: &ListNode{
					Val: 1,
					Next: &ListNode{
						Val: 2,
						Next: &ListNode{
							Val: 3,
							Next: &ListNode{
								Val: 4,
								Next: &ListNode{
									Val:  4,
									Next: nil,
								},
							},
						},
					},
				},
			},
		},
		{
			name: "two nil lists",
			args: args{
				list1: nil,
				list2: nil,
			},
			want: nil,
		},
		{
			name: "first nil list",
			args: args{
				list1: nil,
				list2: &ListNode{
					Val: 1,
					Next: &ListNode{
						Val: 3,
						Next: &ListNode{
							Val:  4,
							Next: nil,
						},
					},
				},
			},
			want: &ListNode{
				Val: 1,
				Next: &ListNode{
					Val: 3,
					Next: &ListNode{
						Val:  4,
						Next: nil,
					},
				},
			},
		},
		{
			name: "second nil list",
			args: args{
				list1: &ListNode{
					Val: 1,
					Next: &ListNode{
						Val: 2,
						Next: &ListNode{
							Val:  4,
							Next: nil,
						},
					},
				},
				list2: nil,
			},
			want: &ListNode{
				Val: 1,
				Next: &ListNode{
					Val: 2,
					Next: &ListNode{
						Val:  4,
						Next: nil,
					},
				},
			},
		},
		{
			name: "short first list",
			args: args{
				list1: &ListNode{
					Val:  3,
					Next: nil,
				},
				list2: &ListNode{
					Val: 1,
					Next: &ListNode{
						Val: 3,
						Next: &ListNode{
							Val:  4,
							Next: nil,
						},
					},
				},
			},
			want: &ListNode{
				Val: 1,
				Next: &ListNode{
					Val: 3,
					Next: &ListNode{
						Val: 3,
						Next: &ListNode{
							Val:  4,
							Next: nil,
						},
					},
				},
			},
		},
		{
			name: "second short list",
			args: args{
				list1: &ListNode{
					Val:  2,
					Next: nil,
				},
				list2: &ListNode{
					Val: 1,
					Next: &ListNode{
						Val: 3,
						Next: &ListNode{
							Val:  4,
							Next: nil,
						},
					},
				},
			},
			want: &ListNode{
				Val: 1,
				Next: &ListNode{
					Val: 2,
					Next: &ListNode{
						Val: 3,
						Next: &ListNode{
							Val:  4,
							Next: nil,
						},
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := mergeTwoLists(tt.args.list1, tt.args.list2); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("mergeTwoLists() = %v, want %v", got, tt.want)
			}
		})
	}
}
