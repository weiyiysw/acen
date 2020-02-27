package easy

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

func isHappy(n int) bool {
	inLoop := make(map[int]int, 0)
	var remain, sum int
	for {
		sum = 0
		for n > 0 {
			remain = n % 10
			sum += remain * remain
			n = n / 10
		}
		if sum == 1 {
			return true
		}
		// if do not exist, add to map
		// else endlessly in a cycle and break
		if _, ok := inLoop[sum]; !ok {
			inLoop[sum] = sum
			n = sum
		} else {
			break
		}
	}
	return false
}

// Count the number of prime numbers less than a non-negative number, n.
// Sieve of Eratosthenes
func countPrimes(n int) int {
	isPrime := make([]bool, n)
	for i := 2; i < n; i++ {
		isPrime[i] = true
	}

	for i := 2; i*i < n; i++ {
		if !isPrime[i] {
			continue
		}
		for j := i * i; j < n; j = j + i {
			isPrime[j] = false
		}
	}
	var count int
	for i := 2; i < n; i++ {
		if isPrime[i] {
			count++
		}
	}
	return count
}

func isIsomorphic(s string, t string) bool {
	replaceMap := make(map[byte]byte, len(s))
	for i := 0; i < len(s); i++ {
		// key exist, if value not equals, return false
		if v, ok := replaceMap[s[i]]; ok {
			if v == t[i] {
				continue
			} else {
				return false
			}
		} else {
			// key is not exist, if value exist, return false
			for j := 0; j < i; j++ {
				if t[j] == t[i] {
					return false
				}
			}
			replaceMap[s[i]] = t[i]
		}
	}
	return true
}

func containsDuplicate(nums []int) bool {
	countMap := make(map[int]int, 0)

	for i := 0; i < len(nums); i++ {
		if _, ok := countMap[nums[i]]; ok {
			return true
		}
		countMap[nums[i]] = 1
	}
	return false
}

// mqtt calculate length, last byte must be 0x7f
func mqttLength2Bytes(b []byte) int {
	var length int
	multiplier := 1
	for i := 0; i < len(b); i++ {
		length += int(b[i]&0x7f) * multiplier
		multiplier *= 128
	}
	return length
}

// at most k
func containsNearbyDuplicate(nums []int, k int) bool {
	for i := 0; i < len(nums); i++ {
		for j := 1; j < k+1 && i+j < len(nums); j++ {
			if nums[i]^nums[i+j] == 0 {
				return true
			}
		}
	}
	return false
}

func isPowerOfTwo(n int) bool {
	i := 1
	for i < n {
		i = i << 1
	}
	return i == n
}

// ListNode  Int Single Linked struct
type ListNode struct {
	Val  int
	Next *ListNode
}

// bad
// time: o(nlogn), space:
func isPalindrome(head *ListNode) bool {
	var length int
	// cal linked length
	p := head
	for p != nil {
		length++
		p = p.Next
	}
	return recursionLinked2(head, head, 1, length)
}

func recursionLinked(h *ListNode) (int, *ListNode) {
	if h == nil {
		return 0, nil
	}
	fmt.Printf("%d \t", h.Val)
	dep, _ := recursionLinked(h.Next)
	return dep + 1, h
}

func recursionLinked2(h, root *ListNode, dep, length int) bool {
	if h == nil {
		return true
	}
	flag := recursionLinked2(h.Next, root, dep+1, length)
	fmt.Printf("val: %d, dep: %d \n", h.Val, dep)
	for j := 0; j < length-dep; j++ {
		root = root.Next
	}
	return flag && h.Val == root.Val
}

func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}
	sCount := make([]int, 26)
	tCount := make([]int, 26)
	for i := 0; i < len(s); i++ {
		x := s[i] - 97
		y := t[i] - 97
		sCount[x]++
		tCount[y]++
	}
	for i := 0; i < 26; i++ {
		if sCount[i] != tCount[i] {
			return false
		}
	}
	return true
}

// TreeNode Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func binaryTreePaths(root *TreeNode) []string {
	if root == nil {
		return nil
	}
	results := make([]string, 0)
	depAccess(root, "", &results)
	return results
}

func depAccess(node *TreeNode, path string, paths *[]string) {
	// if node is leaf node, add path and return
	if node.Left == nil && node.Right == nil {
		path += strconv.Itoa(node.Val)
		*paths = append(*paths, path)
		fmt.Printf("%s \n", path)
		return
	}
	path += strconv.Itoa(node.Val) + "->"
	if node.Left != nil {
		depAccess(node.Left, path, paths)
	}
	if node.Right != nil {
		depAccess(node.Right, path, paths)
	}
}

func addDigits(num int) int {
	if num == 0 {
		return 0
	}
	if num%9 == 0 {
		return 9
	}
	return num % 9
}

// 1、3、5、7、9，奇数序列相加即为 n的平方，原理等差数列
func isPerfectSquare(num int) bool {
	i := 1
	for num > 0 {
		num -= i
		i += 2
	}
	return num == 0
}

func getSum(a int, b int) int {
	return a ^ b
}

func xor(arrs []int) int {
	var res int
	for i := 0; i < len(arrs); i++ {
		res ^= arrs[i]
	}
	return res
}

// problem 88
func merge(nums1 []int, m int, nums2 []int, n int) {
	lastIndex := m + n - 1
	start1, start2 := m-1, n-1
	for start1 > -1 && start2 > -1 {
		if nums1[start1] > nums2[start2] {
			nums1[lastIndex] = nums1[start1]
			start1--
		} else {
			nums1[lastIndex] = nums2[start2]
			start2--
		}
		lastIndex--
	}
	// check
	if start1 == -1 {
		for i := 0; i <= lastIndex; i++ {
			nums1[i] = nums2[i]
		}
	}
	fmt.Printf("nums1: %v \n", nums1)
}

// problem 189
func rotate2(nums []int, k int) {
	if nums == nil || len(nums) == 1 {
		return
	}
	k = k % len(nums)
	if k == 0 {
		return
	}
	size := len(nums)

	counter, i := 0, 0
	// 思路，跟踪求余，size奇数时，i值为0，一遍就可完成size数的遍历
	// size为偶数时，需外层循环控制，需增加i的值，直到counter数等于size时
	// size为偶数时，k = size /2时，i的值需增加到size/2
	for ; counter < size; i++ {
		y := (k + i) % size
		tmp := nums[i]
		for y != i {
			nums[y], tmp = tmp, nums[y]
			y = (y + k) % size
			counter++
		}
		nums[y] = tmp
		counter++
	}
	fmt.Printf("%v \n", nums)
}

func rotate(nums []int, k int) {
	k %= len(nums)
	reverse(nums, 0, len(nums)-1)
	reverse(nums, 0, k-1)
	reverse(nums, k, len(nums)-1)
}

func reverse(nums []int, start int, end int) {
	for start < end {
		nums[start], nums[end] = nums[end], nums[start]
		start++
		end--
	}
}

// problem 198
func rob(nums []int) int {
	if nums == nil || len(nums) == 0 {
		return 0
	}
	if len(nums) == 1 {
		return nums[0]
	}
	size := len(nums)
	result := make([]int, size)
	// init
	result[0], result[1] = nums[0], max(nums[0], nums[1])
	for i := 2; i < size; i++ {
		result[i] = max(result[i-1], result[i-2]+nums[i])
	}
	return result[size-1]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// problem 125, 仅判断字母和数字
func isPalindromeString(s string) bool {
	if "" == s {
		return true
	}
	s = strings.ToLower(s)
	length := len(s)
	low, high := 0, length-1
	for low <= high {
		for low < length && notAlphanumeric(s[low]) {
			low++
		}
		for high > -1 && notAlphanumeric(s[high]) {
			high--
		}
		// not contains alphanumeric
		if low == length {
			return true
		}
		if s[low] != s[high] {
			return false
		}
		low++
		high--
	}
	return true
}

// must convert to lower
func notAlphanumeric(x byte) bool {
	return (((x > 57 || x < 48) && (x < 97)) || x > 122)
}

// problem 349
func intersection(nums1 []int, nums2 []int) []int {
	if nums1 == nil || nums2 == nil || len(nums1) == 0 || len(nums2) == 0 {
		return nil
	}
	m := make(map[int]bool)
	// init map
	for i := 0; i < len(nums1); i++ {
		m[nums1[i]] = false
	}
	// mark map
	for i := 0; i < len(nums2); i++ {
		if _, ok := m[nums2[i]]; ok {
			m[nums2[i]] = true
		}
	}
	// get true values from map
	res := make([]int, 0)
	for k, v := range m {
		if v {
			res = append(res, k)
		}
	}
	return res
}

// problem 326
func isPowerOfThree(n int) bool {
	// use loop
	if n < 0 {
		return false
	}
	for n > 1 {
		if n%3 == 0 {
			n /= 3
		} else {
			return false
		}
	}
	return n == 1
}

func isPowerOfFour(num int) bool {
	// use loop
	if num < 0 {
		return false
	}
	for num > 1 {
		if num%4 == 0 {
			num /= 4
		} else {
			return false
		}
	}
	return num == 1
}

// problem 292
func canWinNim(n int) bool {
	return false
}

func removeLast(n int) int {
	return n & (n - 1)
}

func intersect(nums1 []int, nums2 []int) []int {
	if nums1 == nil || nums2 == nil || len(nums1) == 0 || len(nums2) == 0 {
		return nil
	}
	m := make(map[int]int)
	// init map
	for i := 0; i < len(nums1); i++ {
		m[nums1[i]]++
	}
	res := make([]int, 0)
	for i := 0; i < len(nums2); i++ {
		if _, ok := m[nums2[i]]; ok && m[nums2[i]] > 0 {
			res = append(res, nums2[i])
			m[nums2[i]]--
		}
	}
	return res
}

func reverseVowels(s string) string {
	if len(s) == 0 {
		return ""
	}
	var tmp string
	for i := 0; i < len(s); i++ {
		if isVoewl(s[i : i+1]) {
			tmp += s[i : i+1]
		}
	}
	size := len(tmp)
	var res string
	for i := 0; i < len(s); i++ {
		if isVoewl(s[i : i+1]) {
			res += tmp[size-1 : size]
			size--
		} else {
			res += s[i : i+1]
		}
	}
	return res
}

func isVoewl(s string) bool {
	return strings.EqualFold("a", s) || strings.EqualFold("e", s) || strings.EqualFold("i", s) || strings.EqualFold("o", s) || strings.EqualFold("u", s)
}

// 283
func moveZeroes(nums []int) {
	if nums == nil || len(nums) == 1 {
		return
	}
	zeroIndex := -1
	for i := 0; i < len(nums); i++ {
		if nums[i] == 0 {
			if zeroIndex == -1 {
				zeroIndex = i
			}
			continue
		}
		// if nums[i] != 0, swap nums[zeroIndex] and nums[i]
		if zeroIndex != -1 {
			nums[zeroIndex], nums[i] = nums[i], nums[zeroIndex]
			zeroIndex++
		}
	}
	fmt.Printf("nums: %v", nums)
}

func findDiagonalOrder(matrix [][]int) []int {
	if matrix == nil || len(matrix) == 0 {
		return nil
	}
	up := true
	row, col := len(matrix), len(matrix[0])
	res := make([]int, row*col)
	i, j := 0, 0
	var k int
	for k < row*col-1 {
		res[k] = matrix[i][j]
		k++
		if up {
			if j < col-1 && i == 0 {
				up = false
				j++
			} else if j == col-1 {
				up = false
				i++
			} else {
				j++
				i--
			}
		} else {
			if i < row-1 && j == 0 {
				up = true
				i++
			} else if i == row-1 {
				up = true
				j++
			} else {
				i++
				j--
			}
		}
	}
	res[k] = matrix[i][j]
	return res
}

func arrayPairSum(nums []int) int {
	nums = sort.IntSlice(nums)
	var sum int
	for i := 0; i < len(nums); i += 2 {
		sum += nums[i]
	}
	return sum
}

// 欧几里得算法，计算非负整数p、q最大公约数
func euclid(p, q int) int {
	fmt.Printf("p: %d, q: %d\n", p, q)
	if q == 0 {
		return p
	}
	r := p % q
	return euclid(q, r)
}

func findTheDifference(s string, t string) byte {
	// m := make(map[string]int)
	// for i := 0; i < len(s); i++ {
	// 	m[s[i:i+1]]++
	// }
	// for i := 0; i < len(t); i++ {
	// 	m[t[i:i+1]]--
	// }
	// var res byte
	// for k, v := range m {
	// 	if v == -1 {
	// 		res = byte(rune(k[0]))
	// 	}
	// }
	// return res
	var res byte
	for i := 0; i < len(s); i++ {
		res ^= byte(rune(s[i : i+1][0]))
		res ^= byte(rune(t[i : i+1][0]))
	}
	return res ^ byte(rune(t[len(t)-1 : len(t)][0]))
}

func isSubsequence(s string, t string) bool {
	if len(t) < len(s) {
		return false
	}
	var i, j int
	for ; i < len(s); i++ {
		for j < len(t) {
			if s[i:i+1] == t[j:j+1] {
				break
			}
			j++
		}
		if j < len(t) {
			j++
		} else {
			break
		}
	}
	return i == len(s)
}

func sumOfLeftLeaves(root *TreeNode) int {
	if root == nil {
		return 0
	}

	left := sumOfLeft(root.Left, 0, true)
	right := sumOfLeft(root.Right, 0, false)
	return left + right
}

func sumOfLeft(root *TreeNode, sum int, isLeft bool) int {
	if root == nil {
		return 0
	}
	// left leaves node
	if isLeft && root.Left == nil && root.Right == nil {
		return sum + root.Val
	}
	return sumOfLeft(root.Left, sum, true) + sumOfLeft(root.Right, sum, false)
}

// problem 821
func shortestToChar(S string, C byte) []int {
	i, size := 0, len(S)
	res := make([]int, size)
	first := strings.IndexByte(S, C)
	index := first
	for i < first+1 {
		// first index
		res[i] = index
		index--
		i++
	}
	first = strings.IndexByte(S[i:], C)
	for first != -1 {
		low := i
		high := first + i
		if low < high {
			i = high + 1
			for low < high {
				res[low] = res[low-1] + 1
				res[high-1] = res[high] + 1
				low++
				high--
			}
		} else {
			i++
		}
		// first may be 0
		first = strings.IndexByte(S[i:], C)
	}
	// deal with lastIndex
	for i < size {
		res[i] = res[i-1] + 1
		i++
	}
	return res
}

func maxNumberOfBalloons(text string) int {
	// a b l n o
	res := make([]int, 5)
	for i := 0; i < len(text); i++ {
		switch text[i] {
		case 'a':
			res[0]++
		case 'b':
			res[1]++
		case 'l':
			res[2]++
		case 'n':
			res[3]++
		case 'o':
			res[4]++
		default:
			continue
		}
	}
	res[2] /= 2
	res[4] /= 2
	min := res[0]
	for i := 1; i < len(res); i++ {
		if min > res[i] {
			min = res[i]
		}
	}
	return min
}

func reverseNum(num int) int {
	var result int
	for num%10 != 0 {
		result = result*10 + num%10
		num = num / 10
	}
	return result
}

// problem 1260
// 二维数组索引转换为一维数组索引 (i, j) ==> index,  index = (i*n + j)
func shiftGrid(grid [][]int, k int) [][]int {
	m, n := len(grid), len(grid[0])
	mul := m * n
	mv := k % mul
	if mv == 0 {
		return grid
	}
	var tarIndex int // 在for里面声明，也变慢
	res := make([]int, mul)
	for index := 0; index < mul; index++ {
		// target index
		tarIndex = index + mv
		if tarIndex >= mul {
			tarIndex = tarIndex - mul
		}
		// current grid[i][j] need move
		res[tarIndex] = grid[index/n][index%n]
	}
	for index := 0; index < mul; index++ {
		i, j := index/n, index%n // 28ms
		// i, j = index/n, index%n // 88ms，为何去掉了 用 = 反而更慢呢？
		grid[i][j] = res[index]
	}
	return grid
}
