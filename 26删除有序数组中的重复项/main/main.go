package main

import "fmt"

// 给你一个 非严格递增排列 的数组 nums ，请你 原地 删除重复出现的元素,
// 使每个元素 只出现一次 ，返回删除后数组的新长度。元素的 相对顺序 应该保持 一致 。
// 然后返回 nums 中唯一元素的个数。
//
// 考虑 nums 的唯一元素的数量为 k ，
// 你需要做以下事情确保你的题解可以被通过：
//
// 更改数组 nums ，使 nums 的前 k 个元素包含唯一元素，
// 并按照它们最初在 nums 中出现的顺序排列。nums 的其余元素与 nums 的大小不重要。
// 返回 k 。
// 输入：nums = [1,1,2]
// 输出：2, nums = [1,2,_]
// 解释：函数应该返回新的长度 2 ，
// 并且原数组 nums 的前两个元素被修改为 1, 2 。
// 不需要考虑数组中超出新长度后面的元素。
func main() {
	//nums := []int{1, 1, 2}
	nums2 := []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}
	//输入：nums = [1,1,2]
	//输出：2, nums = [1,2,_]
	//result1 := removeDuplicates(nums)
	//fmt.Println(result1)
	result2 := removeDuplicates(nums2)
	fmt.Println(result2)
}

// 函数应该返回新的长度 2 ，
func removeDuplicates(nums []int) int {
	//k := 0
	//for i := 0; i > len(nums); i++ {
	//	if nums[i] == nums[i+1] {
	//
	//	}
	//}
	lenth := len(nums)
	if lenth == 0 {
		return 0
	}
	K := 1
	for fast := 1; fast < lenth; fast++ {
		// 如果快的不等于自己的fast-1 就 将 nums[K] = nums[fast]
		if nums[fast] != nums[fast-1] {
			nums[K] = nums[fast]
			K++
		}
	}
	return K
}
