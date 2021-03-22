/*
 * Author: Wang P
 * Version: 1.0.0
 * Date: 2021-03-05 17:56:08
 * Description:
 **/
package com.weitrue.leetcode.linear.linkedList;

public class ReverseInteger{
    public static void main(String[] args){
        Solution s = new ReverseInteger().new Solution();
    }

    //Given a signed 32-bit integer x, return x with its digits reversed. If reversi
    //ng x causes the value to go outside the signed 32-bit integer range [-231, 231 -
    // 1], then return 0.
    //
    // Assume the environment does not allow you to store 64-bit integers (signed or
    // unsigned).
    //
    //
    // Example 1:
    // Input: x = 123
    //Output: 321
    // Example 2:
    // Input: x = -123
    //Output: -321
    // Example 3:
    // Input: x = 120
    //Output: 21
    // Example 4:
    // Input: x = 0
    //Output: 0
    //
    //
    // Constraints:
    //
    //
    // -231 <= x <= 231 - 1
    //
    // Related Topics 数学
    // 👍 2571 👎 0

    //leetcode submit region begin(Prohibit modification and deletion)
    class Solution {
        public int reverse(int x) {
            int max = 0x7fffffff, min = 0x80000000;  //int的最大值 最小值
            long res = 0;
            while (x != 0) {
                res = res * 10 + x % 10;
                x = x / 10;
            }
            return res<= max && res>= min? (int)res:0;
        }
    }
    //leetcode submit region end(Prohibit modification and deletion)

}
