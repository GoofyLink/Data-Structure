package main

import "fmt"

// 输入：prices = [7,1,5,3,6,4]
// 输出：7
// 解释：在第 2 天（股票价格 = 1）的时候买入，在第 3 天（股票价格 = 5）的时候卖出, 这笔交易所能获得利润 = 5 - 1 = 4 。
// 随后，在第 4 天（股票价格 = 3）的时候买入，在第 5 天（股票价格 = 6）的时候卖出, 这笔交易所能获得利润 = 6 - 3 = 3 。
// 总利润为 4 + 3 = 7 。

// 动态规划 就是 给定一个问题我们把它拆成一个个子问题直到子问题可以直接解决 然后把子问题的答案保存起来
// 以减少重复计算再根据问题答案反推 得出原问题的一种方法
// 记忆化搜索

// 动态规划

// 1.状态设计
// 每天都有 持有和不持有两种情况
// 设计函数 f(i,j) 表示最大利润
// i为对应天数
// j = 0 表示不持有股票  j = 1表示持有股票
// 则 f[i][0] 表示第 i 天后不持有股票的最大利润
// f[i][1]   表示第i天后持有股票的最大利润

// 2.考虑状态转移的过程
// 第 i天 不持有股票   前一天就不持有 或 前一天持有今天卖
// f[i][0] = f[i-1][0]   f[i-1][1] + prices[i]    加上当天股价

// 第 i 天持有股票  前一天就持有或 前一天不持有今天买
// 只能买一次 因此买入的时候当天没有利润价格为负的
// f[i][1]= max(f[i-1][1] ,   -prices[i])
func main() {

	prices := []int{7, 1, 5, 3, 6, 4}
	result := maxProfit(prices)
	fmt.Println(result)
	result1 := maxProfit2(prices)
	fmt.Println(result1)
}

// 动态规划
func maxProfit(prices []int) int {
	n := len(prices)

	// 创建一个二维数组对应天数
	dp := make([][2]int, n)
	// [0][0]
	// [0][0]
	//fmt.Println(dp)
	dp[0][0] = 0          // 第一天不买 最大利润
	dp[0][1] = -prices[0] // 第一天买入股票 所以最大利润为-prices[0]
	//fmt.Println(dp)
	for i := 1; i < n; i++ {
		// i = 1
		// dp[1][0] = Mmax(dp[1-1][0],dp[1-1][1]+prices[1])

		// 根据状态转移方程计算每天获取的最大利润
		dp[i][0] = Mmax(dp[i-1][0], dp[i-1][1]+prices[i])
		dp[i][1] = Mmax(dp[i-1][1], dp[i-1][0]-prices[i])
	}
	//最后一天所持的最大利润
	return dp[n-1][0]
}

// 再次进行状态设计
// 将二维数组降为一维数组
func maxProfit2(prices []int) int {
	n := len(prices)
	maxProfit := 0
	//dp := make([]int, n)
	minPurchasePrice := prices[0]
	for i := 1; i < n; i++ {
		maxProfit = Mmax(maxProfit, prices[i]-minPurchasePrice)
		minPurchasePrice = Mmin(minPurchasePrice, prices[i])
	}
	return maxProfit
}

func Mmax(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func Mmin(a, b int) int {
	if a < b {
		return a
	}
	return b
}
