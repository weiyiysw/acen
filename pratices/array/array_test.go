package array

import (
	"testing"
)

func Test_reverse(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"case1", args{"abcd"}, "dcba"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := reverse(tt.args.s); got != tt.want {
				t.Errorf("reverse() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_findMaxConsecutiveOnes(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"case0", args{[]int{0}}, 0},
		{"case1", args{[]int{1, 1, 0, 1, 1, 1}}, 3},
		{"case2", args{[]int{0, 0, 1, 0, 1, 1, 1}}, 3},
		{"case3", args{[]int{0, 0, 1, 0, 1, 1, 1, 0, 0, 1, 1, 1, 1, 1, 0}}, 5},
		{"case4", args{[]int{1}}, 1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findMaxConsecutiveOnes(tt.args.nums); got != tt.want {
				t.Errorf("findMaxConsecutiveOnes() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_minSubArrayLen(t *testing.T) {
	type args struct {
		s    int
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"case1", args{7, []int{2, 3, 1, 2, 4, 3}}, 2},
		{"case2", args{70, []int{2, 3, 1, 2, 4, 3}}, 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minSubArrayLen(tt.args.s, tt.args.nums); got != tt.want {
				t.Errorf("minSubArrayLen() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_reverseWords(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"case1", args{"  hello world!  "}, "world! hello"},
		{"case2", args{""}, ""},
		{"case3", args{"   "}, ""},
		{"case4", args{"  a"}, "a"},
		{"case5", args{"  a  b"}, "b a"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := reverseWords(tt.args.s); got != tt.want {
				t.Errorf("reverseWords() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_reverseWords2(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"case1", args{s: "Let's take LeetCode contest"}, "s'teL ekat edoCteeL tsetnoc"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := reverseWords2(tt.args.s); got != tt.want {
				t.Errorf("reverseWords2() = %v, want %v", got, tt.want)
			}
		})
	}
}
