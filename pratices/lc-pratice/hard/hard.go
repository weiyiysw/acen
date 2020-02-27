package hard

import (
	"math/rand"
	"strings"
	"time"
)

// bad sulotion, time limit Exceeded,
// checkValid use list, passed 219 / 230
// checkValid use sum, passed 227 / 230
func badSulotionLongestValidParentheses(s string) int {
	var max int
	sLen := len(s)
	for i := 0; i < sLen-1; i++ {
		for j := sLen; j > i; j-- {
			if checkValid(s[i:j]) {
				if max < j-i {
					max = j - i
				}
			}
		}
	}
	return max
}

// 求和比list要高一点
func checkValid(s string) bool {
	var sum int
	for i := 0; i < len(s); i++ {
		if strings.EqualFold("(", s[i:i+1]) {
			// insert into left stack
			sum++
		} else {
			sum--
			// 遇到右括号，但是无左括号，无效
			if sum < 0 {
				return false
			}
		}
	}
	return sum == 0
}

// wrong answer
func longestValidParentheses(s string) int {
	if len(s) == 0 {
		return 0
	}
	leftArr := make([]int, 0)
	rightArr := make([]int, 0)
	// init
	for i := 0; i < len(s); i++ {
		if strings.EqualFold("(", s[i:i+1]) {
			leftArr = append(leftArr, i)
		} else {
			rightArr = append(rightArr, i)
		}
	}
	// 全是左括号、或右括号、或最大右括号的Index小于最小左括号的Index
	if len(leftArr) == 0 || len(rightArr) == 0 || rightArr[len(rightArr)-1] < leftArr[0] {
		return 0
	}

	var max, cur int
	cur++ // no use
	// 从最后一个 ) 开始，倒序查找
	for i := len(rightArr) - 1; i > -1; i-- {
		// for i
	}
	return max
}

// func firstMissingPositive(nums []int) int {
// 	if nums == nil || len(nums) == 0 {
// 		return 1
// 	}
// 	numsLen := len(nums)
// 	for i := 0; i < numsLen; i++ {
// 		for nums[i] > 0 {
// 			// swap
// 			if nums[i] > 0 {
// 				if nums[i]-1 == i {
// 					break
// 				}
// 				if nums[i] > numsLen {
// 					nums[i] = 0
// 					break
// 				}
// 				if nums[i] != nums[nums[i]-1] {
// 					nums[i], nums[nums[i]-1] = nums[nums[i]-1], nums[i]
// 				} else {
// 					nums[i] = 0
// 				}
// 			}
// 		}
// 	}
// 	for i := 0; i < numsLen; i++ {
// 		if nums[i] <= 0 {
// 			return i + 1
// 		}
// 	}
// 	return numsLen + 1
// }

func firstMissingPositive(nums []int) int {
	n := 1
	i := 0

	for ; i < len(nums); i++ {
		if nums[i] == n {
			n++
			i = -1
		}
	}

	return n
}

// NoPassSolution 710
// 讨论区解法：使用map，重新映射
// 使用数组：思路是对的，但是内存限制了
// 因为无重复，考虑使用 bit 来表示
// N, [0, N) 则需要 N bit, bit 从 0 开始计数
type NoPassSolution struct {
	Number       int
	BlacklistLen int
	BlackEmpty   bool
	Whitelist    []byte
	Count        []byte // 统计whitelist里byte的1的数量
}

// NoPassConstructor 710
func NoPassConstructor(N int, blacklist []int) NoPassSolution {
	if blacklist == nil || len(blacklist) == 0 {
		return NoPassSolution{
			Number:       N,
			BlacklistLen: 0,
			BlackEmpty:   true,
			Whitelist:    nil,
			Count:        nil,
		}
	}

	size := (N / 8) + 1
	whitelist := make([]byte, size, size)
	count := make([]byte, size, size)
	// blacklist not empty
	for i := 0; i < N; i++ {
		// calculator i bit and set 1 if not in blacklist else set 0
		if notInBlacklist(blacklist, i) {
			// i not in blacklist
			// index i % 8
			whitelist[i/8] |= byte(1 << uint(i%8))
			count[i/8]++
		}
	}
	return NoPassSolution{
		Number:       N,
		BlacklistLen: len(blacklist),
		BlackEmpty:   false,
		Whitelist:    whitelist,
		Count:        count,
	}
}

func notInBlacklist(blacklist []int, i int) bool {
	for _, v := range blacklist {
		if i == v {
			return false
		}
	}
	return true
}

// Pick 710
func (this *NoPassSolution) NoPassPick() int {
	rand.Seed(time.Now().UnixNano())
	if this.BlackEmpty {
		return rand.Intn(this.Number)
	}
	// random set whitelist len, then return whiltelist[i]
	// 随机值 = 0 时，应找到第1个1的位置，因此，index = random + 1
	index := rand.Intn(this.Number-this.BlacklistLen) + 1
	// 找到 whitelist 里第 {index} 个 1 返回其 index
	var i, sum int
	for i = 0; i < len(this.Count); i++ {
		sum += int(this.Count[i])
		if sum >= index {
			break
		}
	}
	// 确定 Whitelist 第 i 个 byte 里的 1的位置
	childIndex := index - (sum - int(this.Count[i]))
	j, b := 0, this.Whitelist[i]
	var count int
	for j = 0; j < 8; j++ {
		if ((b >> uint(j)) & 0x01) == byte(1) {
			count++
		}
		if count == childIndex {
			break
		}
	}
	return i*8 + j
}

// Solution 710 discuss
type Solution struct {
	Number   int
	BlackMap map[int]int
}

func Constructor(N int, blacklist []int) Solution {
	if blacklist == nil || len(blacklist) == 0 {
		return Solution{
			Number:   N,
			BlackMap: nil,
		}
	}
	m := make(map[int]int)
	for _, b := range blacklist {
		m[b] = -1
	}
	num := N - len(blacklist)
	for _, b := range blacklist {
		if b < num {
			for _, ok := m[N-1]; ok; _, ok = m[N-1] {
				N--
			}
			m[b] = N - 1
			N--
		}
	}
	return Solution{
		Number:   num,
		BlackMap: m,
	}
}

func (this *Solution) Pick() int {
	rand.Seed(time.Now().UnixNano())
	if this.BlackMap == nil {
		return rand.Intn(this.Number)
	}
	p := rand.Intn(this.Number)
	if v, ok := this.BlackMap[p]; ok {
		return v
	}
	return p
}
