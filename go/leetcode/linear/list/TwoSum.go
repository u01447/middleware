package list

/**
 * Author: Wang P
 * Version: 1.0.0
 * Date: 2021-03-05 17:56:08
 * Description: 两数求和
 **/

//Given an array of integers nums and an integer target, return indices of the t
//wo numbers such that they add up to target.
//
// You may assume that each input would have exactly one solution, and you may n
//ot use the same element twice.
//
// You can return the answer in any order.
//
//
// Example 1:
//
//
//Input: nums = [2,7,11,15], target = 9
//Output: [0,1]
//Output: Because nums[0] + nums[1] == 9, we return [0, 1].
//
//
// Example 2:
//
//
//Input: nums = [3,2,4], target = 6
//Output: [1,2]
//
//
// Example 3:
//
//
//Input: nums = [3,3], target = 6
//Output: [0,1]
//
//
//
// Constraints:
//
//
// 2 <= nums.length <= 103
// -109 <= nums[i] <= 109
// -109 <= target <= 109
// Only one valid answer exists.
//
// Related Topics 数组 哈希表
// 👍 10439 👎 0

//leetcode submit region begin(Prohibit modification and deletion)
func twoSum(nums []int, target int) []int {
	if nums == nil || len(nums) <= 1 {
		panic("Illegal nums")
	}
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
	panic("No two sum solution")
}

//leetcode submit region end(Prohibit modification and deletion)
