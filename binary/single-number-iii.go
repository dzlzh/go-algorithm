/*
  给定一个整数数组 nums，其中恰好有两个元素只出现一次，其余所有元素均出现两次。
  找出只出现一次的那两个元素。
  输入: [1,2,1,3,2,5]
  输出: [3,5]
*/
package main

func singleNumber(nums []int) []int {
	var x, diff, bitmask int
	for _, num := range nums {
		bitmask ^= num
	}
	diff = bitmask ^ (bitmask & (bitmask - 1))

	for _, num := range nums {
		if diff&num == diff {
			x ^= num
		}
	}
	return []int{x, x ^ bitmask}
}
