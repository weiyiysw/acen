package easy

import (
	"reflect"
	"testing"
)

func Test_isHappy(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"case1", args{19}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isHappy(tt.args.n); got != tt.want {
				t.Errorf("isHappy() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_countPrimes(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"case1", args{10}, 4},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := countPrimes(tt.args.n); got != tt.want {
				t.Errorf("countPrimes() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_isIsomorphic(t *testing.T) {
	type args struct {
		s string
		t string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"case1", args{"egg", "add"}, true},
		{"case2", args{"foo", "bar"}, false},
		{"case3", args{"paper", "title"}, true},
		{"case4", args{"ab", "aa"}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isIsomorphic(tt.args.s, tt.args.t); got != tt.want {
				t.Errorf("isIsomorphic() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_containsDuplicate(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"case1", args{[]int{1, 1, 3}}, true},
		{"case2", args{[]int{1, 2, 3}}, false},
		{"case2", args{[]int{1, 2, 3, 1}}, true},
		{"empty", args{[]int{}}, false},
		{"nil", args{nil}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := containsDuplicate(tt.args.nums); got != tt.want {
				t.Errorf("containsDuplicate() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_mqttLength2Bytes(t *testing.T) {
	type args struct {
		b []byte
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// {"max 1 byte", args{[]byte{0x7f}}, 127},
		// {"max 2 byte", args{[]byte{0xff, 0x7f}}, 16383},
		// {"max 3 byte", args{[]byte{0xff, 0xff, 0x7f}}, 2097151},
		// {"max 4 byte", args{[]byte{0xff, 0xff, 0xff, 0x7f}}, 268435455},
		// {"min 1 byte", args{[]byte{0x00}}, 0},
		// {"min 2 byte", args{[]byte{0x80, 0x01}}, 128},
		// {"min 3 byte", args{[]byte{0x80, 0x80, 0x01}}, 16384},
		// {"min 4 byte", args{[]byte{0x80, 0x80, 0x80, 0x01}}, 2097152},
		{"321", args{[]byte{0xc1, 0x02}}, 321},
		{"417", args{[]byte{0xA1, 0x03}}, 417},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := mqttLength2Bytes(tt.args.b); got != tt.want {
				t.Errorf("mqttLength2Bytes() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_containsNearbyDuplicate(t *testing.T) {
	type args struct {
		nums []int
		k    int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"case1", args{[]int{1, 2, 3, 1}, 3}, true},
		{"case2", args{[]int{1, 0, 1, 1}, 1}, true},
		{"case3", args{[]int{1, 2, 3, 1, 2, 3}, 2}, false},
		{"case4", args{[]int{99, 99}, 2}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := containsNearbyDuplicate(tt.args.nums, tt.args.k); got != tt.want {
				t.Errorf("containsNearbyDuplicate() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_isPowerOfTwo(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"1", args{1}, true},
		{"16", args{16}, true},
		{"218", args{218}, false},
		{"1024", args{1}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isPowerOfTwo(tt.args.n); got != tt.want {
				t.Errorf("isPowerOfTwo() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_recursionLinked(t *testing.T) {
	n1 := &ListNode{1, nil}
	n2 := &ListNode{2, n1}
	n3 := &ListNode{2, n2}
	root := &ListNode{1, n3}
	// recursionLinked(root)
	flag := recursionLinked2(root, root, 1, 4)
	t.Logf("flag: %v \n", flag)
}

func Test_isAnagram(t *testing.T) {
	type args struct {
		s string
		t string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"case1", args{"abc", "cba"}, true},
		{"case2", args{"abt", "cba"}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isAnagram(tt.args.s, tt.args.t); got != tt.want {
				t.Errorf("isAnagram() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_addDigits(t *testing.T) {
	type args struct {
		num int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"38", args{38}, 2},
		{"189", args{189}, 9},
		{"3", args{3}, 3},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := addDigits(tt.args.num); got != tt.want {
				t.Errorf("addDigits() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_getSum(t *testing.T) {
	type args struct {
		a int
		b int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"case1", args{1, 2}, 3},
		{"case2", args{-1, 2}, 1},
		{"case3", args{100, 2}, 102},
		{"case4", args{79, 23}, 102},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getSum(tt.args.a, tt.args.b); got != tt.want {
				t.Errorf("getSum() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_xor(t *testing.T) {
	type args struct {
		arrs []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"case1", args{[]int{1, 2, 3, 4, 3, 1, 2}}, 4},
		{"case2", args{[]int{1, 2, 3, 4, 3, 1, 2, 4, 6}}, 6},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := xor(tt.args.arrs); got != tt.want {
				t.Errorf("xor() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_merge(t *testing.T) {
	type args struct {
		nums1 []int
		m     int
		nums2 []int
		n     int
	}
	tests := []struct {
		name string
		args args
	}{
		{"case1", args{[]int{1, 2, 3, 0, 0, 0}, 3, []int{2, 5, 6}, 3}},
		{"case2", args{[]int{7, 8, 9, 0, 0, 0}, 3, []int{2, 5, 6}, 3}},
		{"case3", args{[]int{2, 5, 9, 0, 0, 0}, 3, []int{2, 5, 6}, 3}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			merge(tt.args.nums1, tt.args.m, tt.args.nums2, tt.args.n)
		})
	}
}

func Test_rotate(t *testing.T) {
	type args struct {
		nums []int
		k    int
	}
	tests := []struct {
		name string
		args args
	}{
		{"case1", args{[]int{1, 2, 3, 4, 5, 6, 7}, 3}},
		{"case2", args{[]int{1, 2, 3, 4, 5, 6, 7}, 10}},
		{"case3", args{[]int{1, 2, 3, 4, 5, 6, 7}, 1}},
		{"case4", args{[]int{1, 2, 3, 4, 5, 6}, 2}},
		{"case5", args{[]int{1, 2, 3, 4, 5, 6}, 3}},
		{"case6", args{[]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11}, 4}},
		{"case7", args{[]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11}, 7}},
		{"case8", args{[]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12}, 4}},
		{"case9", args{[]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12}, 7}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rotate(tt.args.nums, tt.args.k)
		})
	}
}

func Test_rob(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"case1", args{[]int{1, 2, 3, 1}}, 4},
		{"case2", args{[]int{2, 7, 9, 3, 1}}, 12},
		{"case3", args{[]int{2, 1}}, 2},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := rob(tt.args.nums); got != tt.want {
				t.Errorf("rob() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_isPalindromeString(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"case1", args{"A man, a plan, a canal: Panama"}, true},
		{"case2", args{"race a car"}, false},
		{"case3", args{" "}, true},
		{"case4", args{" a "}, true},
		{"case5", args{" !%$#% "}, true},
		{"case6", args{"0P"}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isPalindromeString(tt.args.s); got != tt.want {
				t.Errorf("isPalindromeString() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_intersection(t *testing.T) {
	type args struct {
		nums1 []int
		nums2 []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{"case1", args{[]int{1, 2, 2, 1}, []int{2, 2}}, []int{2}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := intersection(tt.args.nums1, tt.args.nums2); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("intersection() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_isPowerOfThree(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"case1", args{27}, true},
		{"case2", args{1}, true},
		{"case3", args{0}, false},
		{"case4", args{-27}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isPowerOfThree(tt.args.n); got != tt.want {
				t.Errorf("isPowerOfThree() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_removeLast(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"case1", args{2}, 0},
		{"case2", args{3}, 2},
		{"case2", args{4}, 0},
		{"case2", args{5}, 4},
		{"case2", args{6}, 4},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := removeLast(tt.args.n); got != tt.want {
				t.Errorf("removeLast() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_intersect(t *testing.T) {
	type args struct {
		nums1 []int
		nums2 []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{"case1", args{[]int{1, 2, 2, 1}, []int{2, 2}}, []int{2, 2}},
		{"case2", args{[]int{4, 9, 5}, []int{9, 4, 9, 8, 4}}, []int{9, 4}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := intersect(tt.args.nums1, tt.args.nums2); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("intersect() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_reverseVowels(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"case1", args{"leetcode"}, "leotcede"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := reverseVowels(tt.args.s); got != tt.want {
				t.Errorf("reverseVowels() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_moveZeroes(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
	}{
		{"case1", args{[]int{0, 1, 0, 3, 12}}},
		{"case2", args{[]int{1, 1, 0, 3, 12}}},
		{"case3", args{nil}},
		{"case4", args{[]int{1, 1, 0, 3, 4, 0, 0, 12}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			moveZeroes(tt.args.nums)
		})
	}
}

func Test_findDiagonalOrder(t *testing.T) {
	type args struct {
		matrix [][]int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		// TODO: Add test cases.
		{"case1", args{[][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}}, []int{1, 2, 4, 7, 5, 3, 6, 8, 9}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findDiagonalOrder(tt.args.matrix); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("findDiagonalOrder() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_euclid(t *testing.T) {
	type args struct {
		p int
		q int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"case1", args{105, 24}, 3},
		{"case2", args{1111111, 1234567}, 1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := euclid(tt.args.p, tt.args.q); got != tt.want {
				t.Errorf("euclid() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_findTheDifference(t *testing.T) {
	type args struct {
		s string
		t string
	}
	tests := []struct {
		name string
		args args
		want byte
	}{
		{"case1", args{"abcd", "abcde"}, 0x65},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findTheDifference(tt.args.s, tt.args.t); got != tt.want {
				t.Errorf("findTheDifference() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_isSubsequence(t *testing.T) {
	type args struct {
		s string
		t string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"case1", args{s: "abc", t: "ahbgdc"}, true},
		{"case2", args{s: "acd", t: "ahbgdc"}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isSubsequence(tt.args.s, tt.args.t); got != tt.want {
				t.Errorf("isSubsequence() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_shortestToChar(t *testing.T) {
	// t.Logf("%s \n", string(0x65))
	type args struct {
		S string
		C byte
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{"case1", args{"loveleetcode", 0x65}, []int{3, 2, 1, 0, 1, 0, 0, 1, 2, 2, 1, 0}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := shortestToChar(tt.args.S, tt.args.C); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("shortestToChar() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_maxNumberOfBalloons(t *testing.T) {
	type args struct {
		text string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"1", args{"nlaebolko"}, 1},
		{"2", args{"loonbalxballpoon"}, 2},
		{"3", args{"leetcode"}, 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxNumberOfBalloons(tt.args.text); got != tt.want {
				t.Errorf("maxNumberOfBalloons() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_reverseNum(t *testing.T) {
	type args struct {
		num int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"case1", args{123}, 321},
		{"case2", args{-123}, -321},
		{"case3", args{1534236469}, 9646324351},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := reverseNum(tt.args.num); got != tt.want {
				t.Errorf("reverseNum() = %v, want %v", got, tt.want)
			}
		})
	}
}
