package main

import (
	"fmt"
	"sort"
)

// 题目描述
// 输入：nums1 = [1,2,3,0,0,0], m = 3, nums2 = [2,5,6], n = 3
// 输出：[1,2,2,3,5,6]
// 解释：需要合并 [1,2,3] 和 [2,5,6] 。
// 合并结果是 [1,2,2,3,5,6] ，其中斜体加粗标注的为 nums1 中的元素。
func main() {
	nums1 := []int{1, 2, 3, 0, 0, 0}
	m := 3
	nums2 := []int{2, 5, 6}
	n := 3

	fmt.Println(nums1, m, nums2, n)
	//输出：[1,2,2,3,5,6]
	//merge1(nums1, m, nums2, n)
	//merge2(nums1, m, nums2, n)
	//merge3(nums1, m, nums2, n)
	merge4(nums1, m, nums2, n)
	fmt.Println("merge合并后的结果:", nums1)
}

// 用切片合并两个数组后进行排序
func merge1(nums1 []int, m int, nums2 []int, n int) {
	copy(nums1[m:], nums2)
	sort.Ints(nums1)
}

// 另，开辟一个数组 双指针法排序两数组，最后再拷贝回num1
func merge2(nums1 []int, m int, nums2 []int, n int) {
	sort := make([]int, 0, m+n)
	i, j := 0, 0
	for {
		if i == m {
			sort = append(sort, nums2[j:]...)
			break
		}
		if j == n {
			sort = append(sort, nums1[i:]...)
			break
		}
		if nums1[i] < nums2[j] {
			sort = append(sort, nums1[i])
			i++
		} else {
			sort = append(sort, nums2[j])
			j++
		}
	}
	copy(nums1, sort)
}

// 逆双指针
func merge3(nums1 []int, m int, nums2 []int, n int) {
	i, j, tail := m-1, n-1, m+n-1
	for i >= 0 || j >= 0 {
		var cur int
		if i == -1 {
			cur = nums2[j]
			j--
		} else if j == -1 {
			cur = nums1[i]
			i--
		} else if nums1[i] < nums2[j] {
			cur = nums2[j]
			j--
		} else {
			cur = nums1[i]
			i--
		}
		nums1[tail] = cur
		tail--
	}
}

// 本题的要求是，把nums1的前m项和nums2的前n项合并，放入nums1中。
func merge4(nums1 []int, m int, nums2 []int, n int) {
	//把nums1复制到temp中
	temp := make([]int, m)
	copy(temp, nums1)

	t, j := 0, 0 //t为temp的索引，j为nums2的索引
	for i := 0; i < len(nums1); i++ {
		//当t大于temp的长度，那就是说temp全部放进去了nums1中，那剩下的就是放nums2剩余的值了
		if t >= len(temp) {
			nums1[i] = nums2[j]
			j++
			continue
		}
		//当j大于nums2的长度的时候，那就是说明nums2全部都放进去了nums1中，那剩下的就是放temp剩余的值了
		if j >= n {
			nums1[i] = temp[t]
			t++
			continue
		}
		//比较nums2与temp对应值的大小，小的那个就放进nums1中
		if nums2[j] <= temp[t] {
			nums1[i] = nums2[j]
			j++
		} else {
			nums1[i] = temp[t]
			t++
		}
	}
}
