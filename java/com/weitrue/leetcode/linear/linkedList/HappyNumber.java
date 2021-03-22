/*
 * Author: Wang P
 * Version: 1.0.0
 * Date: 2021-03-05 17:56:08
 * Description:
 **/
package com.weitrue.leetcode.linear.linkedList;

public class HappyNumber{
    public static void main(String[] args){
        Solution s = new HappyNumber().new Solution();
        int sum = 0;
        for (int i=0; i<=100000; i++) {
            if (s.isHappy(i)) {
                sum++;
            }
        }
        System.out.println(sum);
    }

    //Write an algorithm to determine if a number n is happy.
    //
    // A happy number is a number defined by the following process:
    //
    // Starting with any positive integer, replace the number by the sum of the squares of its digits.
    // Repeat the process until the number equals 1 (where it will stay), or it loops endlessly in a cycle which does not include 1.
    // Those numbers for which this process ends in 1 are happy.
    //
    // Return true if n is a happy number, and false if not.
    //
    // Example 1:
    //Input: n = 19
    //Output: true
    //Explanation:
    //12 + 92 = 82
    //82 + 22 = 68
    //62 + 82 = 100
    //12 + 02 + 02 = 1
    //
    // Example 2:
    //Input: n = 2
    //Output: false
    //
    // Constraints:
    //
    // 1 <= n <= 231 - 1
    //
    // Related Topics 哈希表 数学
    // 👍 542 👎 0

    //leetcode submit region begin(Prohibit modification and deletion)
    class Solution {
            public int getNext(int n) {
                int x = 0;
                while(n > 0) {
                    x += (n % 10) * (n % 10);
                    n = n / 10;
                }
                return x;
            }

            public boolean isHappy(int n) {
                int p = n, q = n;
                do{
                    p = getNext(p);
                    q = getNext(getNext(q));
                } while (p != q && q!= 1);
                return q == 1;
            }
    }
    //leetcode submit region end(Prohibit modification and deletion)

}
