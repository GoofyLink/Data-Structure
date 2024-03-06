package main

import "fmt"

//给你一个有序数组 nums ，请你 原地 删除重复出现的元素，
//使得出现次数超过两次的元素只出现两次 ，返回删除后数组的新长度。
//不要使用额外的数组空间，你必须在 原地 修改输入数组 并在使用 O(1) 额外空间的条件下完成。
//输入：nums = [1,1,1,2,2,3]
//输出：5, nums = [1,1,2,2,3]
//解释：函数应返回新长度 length = 5, 并且原数组的前五个元素被修改为 1, 1, 2, 2, 3。
//不需要考虑数组中超出新长度后面的元素。

// 需要检查上上个元素  也就是  nums[slow-2]
// 需要检查当前指针  nums[fast] = nums[slow-2]   就不应该被保留
// 此时必然有 nums[slow-2] = nums[slow-1] = nums[fast]
// slow就是处理好的数组长度
func main() {

	nums := []int{1, 1, 1, 2, 2, 3}
	result := removeDuplicates(nums)
	fmt.Println(result)
	//5, nums = [1,1,2,2,3]
}

// 返回数组的新长度
func removeDuplicates(nums []int) int {
	length := len(nums)
	// 判断数组不为空
	if length <= 2 {
		return length
	}
	// 可以使用双指针
	slow := 2
	fast := 2
	for fast < length {
		//上上个元素  也就是  nums[slow-2]
		//不等于当前元素
		// slow = 2  fast = 2   {1, 1, 1, 2, 2, 3}
		//1. slow:nums[2-2] = nums[0] = 1  fast:nums[2] = 1
		//2. slow:nums[2-2] = nums[0] = 1  fast:nums[3] = 2
		//3. slow++:nums[3-2] = nums[1] = 1  fast:nums[4] = 2
		//4. slow++ nums[4-2]
		//5. slow++ nums[5-2]                 fast:nums[6]
		if nums[slow-2] != nums[fast] {
			nums[slow] = nums[fast]
			slow++
		}
		fast++
	}
	return slow
}
