package main

import (
	"fmt"
	"sort"
)

//给你一个整数数组 citations ，其中 citations[i] 表示研究者的第 i
//篇论文被引用的次数。计算并返回该研究者的 h 指数。
//根据维基百科上 h 指数的定义：h 代表“高引用次数” ，一名科研人员的 h 指数
//是指他（她）至少发表了 h 篇论文，并且 至少 有 h 篇论文被引用次数大于等于 h 。
//如果 h 有多种可能的值，h 指数 是其中最大的那个。

//输入：citations = [3,0,6,1,5]
//输出：3
//解释：给定数组表示研究者总共有 5 篇论文，每篇论文相应的被引用了 3, 0, 6, 1, 5 次。
//由于研究者有 3 篇论文每篇 至少 被引用了 3 次，其余两篇论文每篇被引用 不多于 3 次，所以她的 h 指数是 3。

func main() {
	citations := []int{3, 0, 6, 1, 5}
	citations2 := []int{1, 3, 1}
	result := hIndex(citations)
	fmt.Println(result)
	result2 := hIndex(citations2)
	fmt.Println(result2)
}

func hIndex(citations []int) (h int) {
	sort.Ints(citations) // 先排序
	for i := len(citations) - 1; i >= 0 && citations[i] > h; i-- {
		h++
	}
	return
}

// 二分搜索
func hIndex2(citations []int) int {
	// 答案最多只能到数组长度
	left, right := 0, len(citations)
	var mid int
	for left < right {
		// +1 防止死循环
		mid = (left + right + 1) >> 1
		cnt := 0
		for _, v := range citations {
			if v >= mid {
				cnt++
			}
		}
		if cnt >= mid {
			// 要找的答案在 [mid,right] 区间内
			left = mid
		} else {
			// 要找的答案在 [0,mid) 区间内
			right = mid - 1
		}
	}
	return left
}
