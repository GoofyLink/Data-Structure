package main

import "fmt"

//给你一个非负整数数组 nums ，你最初位于数组的 第一个下标 。
//数组中的每个元素代表你在该位置可以跳跃的最大长度。
//判断你是否能够到达最后一个下标，如果可以，返回 true ；否则，返回 false 。

//输入：nums = [2,3,1,1,4]
//输出：true
//解释：可以先跳 1 步，从下标 0 到达下标 1, 然后再从下标 1 跳 3 步到达最后一个下标。

func main() {
	nums := []int{2, 3, 1, 1, 4}
	result := canJump(nums)
	fmt.Println(result)
}

func canJump(nums []int) bool {
	n := len(nums) - 1
	cover := 0
	for i := 0; i <= cover; i++ {
		cover = Mmax(i+nums[i], cover)
		if cover >= n {
			return true
		}
	}
	return false
}

func Mmax(a, b int) int {
	if a > b {
		return a
	}
	return b
}
