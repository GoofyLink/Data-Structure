package main

import "fmt"

//给你一个数组 nums 和一个值 val，你需要 原地 移除所有数值等于 val 的元素，并返回移除后数组的新长度。
//不要使用额外的数组空间，你必须仅使用 O(1) 额外空间并 原地 修改输入数组。
//元素的顺序可以改变。你不需要考虑数组中超出新长度后面的元素。

// 输入：nums = [3,2,2,3], val = 3
// 输出：2, nums = [2,2]
// 解释：函数应该返回新的长度 2,
// 并且 nums 中的前两个元素均为 2。
// 你不需要考虑数组中超出新长度后面的元素。例如，函数返回的新长度为 2 ，
// 而 nums = [2,2,3,3] 或 nums = [2,2,0,0]，也会被视作正确答案。
func main() {
	nums := []int{3, 2, 2, 3}
	val := 3
	// 移除nums数组中的元素3
	// 必须使用空间为O(1)复杂度
	removeElement(nums, val)
	fmt.Println(nums)
}

func removeElement(nums []int, val int) int {
	length := len(nums)
	if length == 0 {
		return 0
	}
	i := 0
	j := 0
	for j < length {
		if nums[j] == val {
			// 去找一个不是val的值
			j++
		} else {
			// 复制
			nums[i] = nums[j]
			i++
			j++
		}
	}
	return length - (j - i)
}
