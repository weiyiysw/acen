package hard

import (
	"encoding/hex"
	"testing"
)

func Test_longestValidParentheses(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"case1", args{"(()()"}, 4},
		{"case2", args{"((())()"}, 6},
		{"case3", args{")((()"}, 2},
		{"case4", args{"()((()"}, 2},
		{"case5", args{")()())()()("}, 4},
		{"case6", args{"))))())()()(()"}, 4},
		{"case7", args{"))())(()))()(((()())(()(((()))))((()(())()((((()))())))())))()(()(()))))())(((())(()()))((())()())((()))(()(())(())((())((((()())()))((()(())()))()(()))))))()))(()))))()())()())()()()()()()()))()(((()()((()(())((()())))(()())))))))(()()(())())(()))))))()()())((((()()()())))))((())(())()()(()((()()))()()())(()())()))()(()(()())))))())()(())(()))(())()(())()((())()((((()()))())(((((())))())())(()((())((()()((((((())))(((())))))))(()()((((((()(((())()(()))(()())((()(((()((()(())())()())(((()))()(((()))))(())))(())()())()(((()))))((())())))())()()))((((()))(())()())()(((())(())(()()((())()())()()())())))((()())(()((()()()(()())(()))(()())((((()(()(((()(((())()((()(()))())()())))))))))))()())()(()(((())()))(((()))((((()())())(()())((()())(()()((()((((()())))()(())(())()))))(())())))))(((((((())(((((()))()))(()()()()))))))(()(()(()(()()(((()()))((()))())((())())()())()))()()(((())))()(())()()(())))(((()))))))))(())((()((()((()))))()()()((())((((((((((()(())))(())((()(()())())(((((((()()()()))())(((()())()(()()))))(()()))))(((()()((()()()(((()))))(()()())()()()(()))))()(())))))))()((((((((()((())))))))(()))()((()())())("}, 490},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := badSulotionLongestValidParentheses(tt.args.s); got != tt.want {
				t.Errorf("badSulotionlongestValidParentheses() = %v, want %v", got, tt.want)
			}
			// if got := longestValidParentheses(tt.args.s); got != tt.want {
			// 	t.Errorf("longestValidParentheses() = %v, want %v", got, tt.want)
			// }
		})
	}
}

func Test_firstMissingPositive(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
		{"case1", args{[]int{3, 1, 0}}, 2},
		{"case2", args{[]int{7, 8, 9, 10, 11, 12}}, 1},
		{"case3", args{[]int{3, 1, 0, -1}}, 2},
		{"case4", args{[]int{2, 1, 0}}, 3},
		{"case5", args{[]int{0}}, 1},
		{"case5", args{[]int{}}, 1},
		{"case6", args{[]int{1}}, 2},
		{"case7", args{[]int{1, 2}}, 3},
		{"case8", args{[]int{-1, -2}}, 1},
		{"case9", args{[]int{1, 1}}, 2},
		{"case9", args{[]int{1, 1, -1, -1, -3, 2, 2, 4, 5, 6}}, 3},
		{"case9", args{[]int{2, 1}}, 3},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := firstMissingPositive(tt.args.nums); got != tt.want {
				t.Errorf("firstMissingPositive() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_Byte(t *testing.T) {
	tmp := make([]byte, 2)
	for i := 0; i < 16; i++ {
		if i != 3 {
			tmp[i/8] |= byte(1 << (i % 8))
		}
	}
	t.Logf("%v \n", tmp)
	t.Logf("%s \n", hex.EncodeToString(tmp))
}

func Test_Solution(t *testing.T) {
	s := Constructor(3, []int{0, 2})
	for i := 0; i < 10; i++ {
		t.Logf("i: %d, %d \n", i, s.Pick())
	}
}
