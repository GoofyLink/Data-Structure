package main

import "fmt"

//给你两个字符串 haystack 和 needle ，请你在 haystack
//字符串中找出 needle 字符串的第一个匹配项的下标（下标从 0 开始）。
//如果 needle 不是 haystack 的一部分，则返回  -1 。

//输入：haystack = "sadbutsad", needle = "sad"
//输出：0
//解释："sad" 在下标 0 和 6 处匹配。
//第一个匹配项的下标是 0 ，所以返回 0 。

func main() {
	haystack := "sadbutsad"
	needle := "sad"
	result := strStr(haystack, needle)
	fmt.Println(result)
	haystack2 := "leetcode"
	needle2 := "leeto"
	r := strStr(haystack2, needle2)
	fmt.Println(r)
}

func strStr(haystack string, needle string) int {
	h := len(haystack)
	n := len(needle)
	for i := 0; i+n <= h; i++ {
		for j := range needle {
			if haystack[i+j] != needle[j] {
				continue
			}
		}
		return i
	}
	return -1
}

// kmp算法
// 根据子串字符串构造一个next部分匹配表
// 遍历主串字符串,当遍历失败时,查询next部分匹配表,定位子串接着与主串比较的位置
func strStr2(haystack, needle string) int {
	n, m := len(haystack), len(needle)
	if m == 0 {
		return 0
	}
	pi := make([]int, m)
	for i, j := 1, 0; i < m; i++ {
		for j > 0 && needle[i] != needle[j] {
			j = pi[j-1]
		}
		if needle[i] == needle[j] {
			j++
		}
		pi[i] = j
	}
	for i, j := 0, 0; i < n; i++ {
		for j > 0 && haystack[i] != needle[j] {
			j = pi[j-1]
		}
		if haystack[i] == needle[j] {
			j++
		}
		if j == m {
			return i - m + 1
		}
	}
	return -1
}
