package main

import "fmt"

//给定一个整数数组 nums，
//将数组中的元素向右轮转 k 个位置，其中 k 是非负数。

//输入: nums = [1,2,3,4,5,6,7], k = 3
//输出: [5,6,7,1,2,3,4]
//解释:
//向右轮转 1 步: [7,1,2,3,4,5,6]
//向右轮转 2 步: [6,7,1,2,3,4,5]
//向右轮转 3 步: [5,6,7,1,2,3,4]

func main() {
	nums := []int{1, 2, 3, 4, 5, 6, 7}
	k := 3
	rotate(nums, k)
	fmt.Println("轮转后的数组", nums)
}

func rotate(nums []int, k int) {
	if len(nums) < 2 { //后面 copy 需要 nums 的长度大于等于2，小于2的数组翻转过来还是它自己，所以就直接 return 了
		return
	}
	newNums := make([]int, len(nums))
	for i, v := range nums {
		newNums[(i+k)%len(nums)] = v
	}
	copy(nums, newNums)
}

func rotate2(nums []int, k int) {
	if len(nums) < 2 { //后面 copy 需要 nums 的长度大于等于2，小于2的数组翻转过来还是它自己，所以就直接 return 了
		return
	}
	for i := 1; i <= k; i++ {
		value := nums[len(nums)-1] //每次把数组最后一个取出来，拷贝到数组第一位
		copy(nums[1:], nums[0:])
		nums[0] = value
	}
}
