package main

import "fmt"

//
//给定一个长度为 n 的 0 索引整数数组 nums。初始位置为 nums[0]。
//
//每个元素 nums[i] 表示从索引 i 向前跳转的最大长度。换句话说，如果你在 nums[i] 处，你可以跳转到任意 nums[i + j] 处:
//
//0 <= j <= nums[i]
//i + j < n
//返回到达 nums[n - 1] 的最小跳跃次数。生成的测试用例可以到达 nums[n - 1]。

// 输入: nums = [2,3,1,1,4]
// 输出: 2
// 解释: 跳到最后一个位置的最小跳跃数是 2。
// 从下标为 0 跳到下标为 1 的位置，跳 1 步，然后跳 3 步到达数组的最后一个位置。
func main() {
	nums := []int{2, 3, 1, 1, 4}
	result := jump(nums)
	fmt.Println(result)

	r := jump2(nums)
	fmt.Println(r)
}

func jump(nums []int) int {
	position := len(nums) - 1
	steps := 0

	for position > 0 {
		for i := 0; i < position; i++ {
			if i+nums[i] >= position {
				position = i
				steps++
				break
			}
		}
	}
	return steps
}

func Max(x, y int) int {
	if x < y {
		return y
	} else {
		return x
	}
}

func jump2(nums []int) int {
	steps := 0
	end := 0
	maxPos := 0
	for i := 0; i < len(nums)-1; i++ {
		// 找到这一步最大的一步
		maxPos = Max(maxPos, nums[i]+i)
		// 没有跳到就继续往下找
		if i == end {
			end = maxPos
			steps++
		}
	}
	return steps
}
