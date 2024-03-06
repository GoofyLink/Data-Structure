package main

import "fmt"

// 169 多数元素
// 给定一个大小为 n 的数组 nums ，返回其中的多数元素。
// 多数元素是指在数组中出现次数 大于 ⌊ n/2 ⌋ 的元素。
// 你可以假设数组是非空的，并且给定的数组总是存在多数元素。

// 用摩尔投票法 也被称为多数投票法
// 算法解决的问题是：如何在任意多的候选人中（选票无序），选出获得票数最多的那个。
// 算法可以分为两个阶段：
//
// 1.对抗阶段：分属两个候选人的票数进行两两对抗抵消
// 2.计数阶段：计算对抗结果中最后留下的候选人票数是否有效
// 输入：nums = [2,2,1,1,1,2,2]
// 输出：2
func main() {
	nums := []int{2, 2, 1, 1, 1, 2, 2}
	result := majorityElement(nums)
	fmt.Println("候选人最多的是:", result)
}

func majorityElement(nums []int) int {
	major := 0 // 候选人持票
	count := 0 // 记录持票数量
	for _, num := range nums {
		// 没有数量的时候 将候选人选中吧num赋值给候选人
		if count == 0 {
			major = num
		}
		// 如果相等就把票数+1
		if major == num {
			count += 1
		} else {
			//如果major 不等于num 票数-1
			count -= 1
		}
	}
	// 最后返回票最多的候选人
	return major
}
