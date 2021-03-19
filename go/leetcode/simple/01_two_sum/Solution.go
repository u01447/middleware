/*
@Time : 2020/8/6 14:48
@Author : Wang P
@File : Solution
*/
package main

func twoSum(nums []int, target, flag int) []int {
	/*
	   给定一个整数数组 nums 和一个目标值 target，请你在该数组中找出和为目标值的那 两个 整数，并返回他们的数组下标。
	   你可以假设每种输入只会对应一个答案。但是，数组中同一个元素不能使用两遍。
	     示例:
	       给定 nums = [2, 7, 11, 15], target = 9
	       因为 nums[0] + nums[1] = 2 + 7 = 9
	       所以返回 [0, 1]
	   Related Topics 数组 哈希表
	   👍 8826 👎 0
	*/
	if nums == nil || len(nums) <= 1 {
		panic("Illegal nums")
	}
	switch flag {
	case 1:
		for i := 0; i < len(nums)-1; i++ {
			for j := i + 1; j < len(nums); j++ {
				if nums[i]+nums[j] == target {
					return []int{nums[i], nums[j]}
				}
			}
		}
	default:
		numMap := make(map[int]int)
		numMap[nums[0]] = 0
		temp := nums[0]
		for i, num := range nums {
			if i > 0 {
				temp = target - num
				if v, ok := numMap[temp]; ok {
					return []int{v, i}
				} else {
					numMap[num] = i
				}
			}
		}
	}
	panic("No two sum solution")
}
