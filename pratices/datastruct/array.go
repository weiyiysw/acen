package datastruct

import (
	"strings"
)

// 数组问题，双指针
// 双指针迭代
// 一个从头开始，一个从尾开始，在中间结束
// 通常在已排好序的数组中经常使用
func reverse(s string) string {
	i, j := 0, len(s)-1
	tmp := []rune(s)
	for i < j {
		tmp[i], tmp[j] = tmp[j], tmp[i]
		i++
		j--
	}
	return string(tmp)
}

// 双指针：不同的步长，解决问题
// 1快 1慢
// 两个指针的移动策略
func findMaxConsecutiveOnes(nums []int) int {
	if nums == nil || len(nums) == 0 {
		return 0
	}
	slow, fast := 0, 0
	max, size := 0, len(nums)
	for slow < size {
		// init slow 、fast
		for slow < size && nums[slow] == 0 {
			slow++
		}
		// 此处 fast 不能为 slow+1，因为不能确定 nums[slow] == 1
		// 即 当 slow的循环退出时，有两种情况，1：nums[slow] = 1, 2：slow++后，已经大于 size了。
		// 例如：[0]
		fast = slow
		for fast < size && nums[fast] == 1 {
			fast++
		}
		tmp := fast - slow
		if tmp > max {
			max = tmp
		}
		slow = fast + 1
	}
	return max
}

func minSubArrayLen(s int, nums []int) int {
	slow, fast := 0, 0
	sum, size := 0, len(nums)
	min := size

	// 当sum < s 时，并且fast已到了最后，此时即可退出循环
	for slow < size {
		// 计算连续子数组之和
		for sum < s {
			if fast == size {
				break
			}
			sum += nums[fast]
			fast++
		}

		// fast = size，循环会退出，此时sum可能不大于s
		for sum >= s {
			// 记录最小的连续数组长度
			tmp := fast - slow
			if tmp < min {
				min = tmp
			}
			// 重置 slow
			sum -= nums[slow]
			slow++
		}
		// 当sum < s 时，并且fast已到了最后，此时即可退出循环
		if sum < s && fast == size {
			// 判断是否数组之和均小于s
			if (fast - slow) == size {
				return 0
			}
			break
		}
	}
	return min
}

func reverseWords(s string) string {
	strs := strings.Split(s, " ")
	// fmt.Printf("strs: %v %d\n", strs, len(strs))
	// 去除空格
	tmp := make([]string, 0)
	for i := 0; i < len(strs); i++ {
		// fmt.Printf("i: %d, val：%s\n", i, strs[i])
		if !strings.EqualFold(strs[i], "") {
			tmp = append(tmp, strs[i])
		}
	}
	// 翻转tmp
	i, j := 0, len(tmp)-1
	for i < j {
		tmp[i], tmp[j] = tmp[j], tmp[i]
		i++
		j--
	}
	if len(tmp) == 0 {
		return ""
	}
	// 拼接字符
	var res string
	for i = 0; i < len(tmp); i++ {
		res += tmp[i] + " "
	}
	return res[0 : len(res)-1]
}

func reverseWords2(s string) string {
	strs := strings.Split(s, " ")
	tmp := make([]string, len(strs))
	for i := 0; i < len(strs); i++ {
		// reverse string
		rArr := []rune(strs[i])
		low, high := 0, len(rArr)-1
		for low < high {
			rArr[low], rArr[high] = rArr[high], rArr[low]
			low++
			high--
		}
		tmp[i] = string(rArr)
	}
	if len(tmp) == 0 {
		return ""
	}
	var res string
	for i := 0; i < len(tmp); i++ {
		res += tmp[i] + " "
	}
	return res[0 : len(res)-1]
}
