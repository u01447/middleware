/*
 * Author: Wang P
 * Version: 1.0.0
 * Date: 2021-03-05 18:07:46
 * Description:
 **/
 
package com.weitrue.leetcode.editor.cn;

import java.util.HashMap;
import java.util.Map;

public class TwoSum{
    public static void main(String[] args){
        Solution s = new TwoSum().new Solution();
    }
    
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
class Solution {
    public int[] twoSum(int[] nums, int target) {
        if(nums == null || nums.length <= 1){
            throw new IllegalArgumentException("nums illagal");
        }
        Map<Integer, Integer> map = new HashMap<Integer, Integer>();
        int temp = nums[0];
        map.put(nums[0], 0);
        for (int i=1; i < nums.length; i++){
            temp = target - nums[i];
            if (map.containsKey(temp)){
                return new int[]{map.get(temp), i};
            }else{
                map.put(nums[i], i);
            }
        }
        throw new IllegalArgumentException("No two sunm solution");
    }
}
//leetcode submit region end(Prohibit modification and deletion)

}
