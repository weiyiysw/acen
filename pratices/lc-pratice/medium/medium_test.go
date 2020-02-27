package medium

import (
	"testing"
)

func Test_checkValidString(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"case1", args{"()"}, true},
		{"case2", args{"(*)"}, true},
		{"case3", args{"(*))"}, true},
		{"case4", args{"(*())("}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := checkValidString(tt.args.s); got != tt.want {
				t.Errorf("checkValidString() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_shift(t *testing.T) {
	type args struct {
		x rune
		n int
	}
	tests := []struct {
		name string
		args args
		want rune
	}{
		{"case1", args{rune(97), 1}, rune(98)},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := shift(tt.args.x, tt.args.n); got != tt.want {
				t.Errorf("shift() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_shiftingLetters(t *testing.T) {
	type args struct {
		S      string
		shifts []int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"case1", args{"abc", []int{3, 5, 9}}, "rpl"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := shiftingLetters(tt.args.S, tt.args.shifts); got != tt.want {
				t.Errorf("shiftingLetters() = %v, want %v", got, tt.want)
			}
		})
	}
}
