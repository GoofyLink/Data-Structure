package main

import (
	"math/rand"
)

//实现RandomizedSet 类：
//
//RandomizedSet() 初始化 RandomizedSet 对象
//bool insert(int val) 当元素 val 不存在时，向集合中插入该项，并返回 true ；否则，返回 false 。
//bool remove(int val) 当元素 val 存在时，从集合中移除该项，并返回 true ；否则，返回 false 。
//int getRandom() 随机返回现有集合中的一项（测试用例保证调用此方法时集合中至少存在一个元素）。
//每个元素应该有 相同的概率 被返回。
//你必须实现类的所有函数，并满足每个函数的 平均 时间复杂度为 O(1) 。

//
//输入
//["RandomizedSet", "insert", "remove", "insert", "getRandom", "remove", "insert", "getRandom"]
//[[], [1], [2], [2], [], [1], [2], []]
//输出
//[null, true, false, true, 2, true, false, 2]

func main() {
	randomString := Constructor()
	//strArr := []string{"RandomizedSet", "insert", "remove", "insert", "getRandom", "remove", "insert", "getRandom"}
	//fmt.Println(randomString)

}

type RandomizedSet struct {
	idices map[int]int
	nums   []int
}

func Constructor() RandomizedSet {
	var set RandomizedSet
	set.idices = map[int]int{}
	set.nums = []int{}
	return set
}

func (this *RandomizedSet) Insert(val int) bool {
	if _, ok := this.idices[val]; ok {
		return false
	}
	this.idices[val] = len(this.nums)
	this.nums = append(this.nums, val)
	return true
}

func (this *RandomizedSet) Remove(val int) bool {
	idx, ok := this.idices[val]
	if !ok {
		return false
	}
	// 将最后一个元素移到 val的位置idx处
	length := len(this.nums)
	endVal := this.nums[length-1]
	this.nums[idx] = endVal
	this.idices[endVal] = idx
	// 删除索引中val的值
	delete(this.idices, val)
	// 删除最后一个元素
	this.nums = this.nums[:length-1]
	return true
}

func (this *RandomizedSet) GetRandom() int {
	idx := rand.Intn(len(this.nums))
	return this.nums[idx]
}
