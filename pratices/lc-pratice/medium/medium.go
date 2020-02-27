package medium

import "container/list"

import "fmt"

// ValidString checkValidString方法的结构体
type ValidString struct {
	Character rune
	Index     int
}

// 678. Valid Parenthesis String
func checkValidString(s string) bool {
	left := list.New()
	star := list.New()

	for k, c := range s {
		val := ValidString{c, k}
		if c == '(' {
			left.PushFront(val)
		}
		if c == '*' {
			star.PushFront(val)
		}
		if c == ')' {
			if left.Len() > 0 {
				left.Remove(left.Front())
				continue
			}
			if star.Len() > 0 {
				star.Remove(star.Front())
				continue
			}
			return false
		}
	}
	// check star and left stack
	for left.Len() > 0 && star.Len() > 0 {
		le := left.Front()
		se := star.Front()
		leftVal, _ := le.Value.(ValidString)
		starVal, _ := se.Value.(ValidString)
		if leftVal.Index > starVal.Index {
			return false
		}
		left.Remove(le)
		star.Remove(se)
	}
	return left.Len() == 0
}

// 848 shiftingLetters
// Note:
// 1 <= S.length = shifts.length <= 20000
// 0 <= shifts[i] <= 10 ^ 9
func shiftingLetters(S string, shifts []int) string {
	rs := []rune(S)
	size := len(shifts)
	var sum int
	for i := size - 1; i >= 0; i-- {
		sum += shifts[i]
		rs[i] = shift(rs[i], sum%26)
	}
	return string(rs)
}

// 定义的shift方法本质是循环移位
// [a ~ z] ascii [97, 122]
func shift(x rune, n int) rune {
	mv := ((int(x) - 97) + n) % 26
	fmt.Printf("rune x: %v, rune res: %v\n", string(x), string(rune(97+mv)))
	return rune(97 + mv)
}
