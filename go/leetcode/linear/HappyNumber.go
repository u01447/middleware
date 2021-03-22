package leetcode

/**
 * Author: Wang P
 * Version: 1.0.0
 * Date: 2021-03-05 12:56:41
 * Description: //Write an algorithm to determine if a number n is happy.
				//
				// A happy number is a number defined by the following process:
				//
				// Starting with any positive integer, replace the number by the sum of the squa
				//res of its digits.
				// Repeat the process until the number equals 1 (where it will stay), or it loop
				//s endlessly in a cycle which does not include 1.
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
				// 1 <= n <= 231 - 1
				//
				// Related Topics 哈希表 数学
				// 👍 547 👎 0
 **/

//leetcode submit region begin(Prohibit modification and deletion)
func IsHappy(n int) bool {
	p, q := getNext(n), getNext(getNext(n))
	for p != q && q != 1 {
		p = getNext(p)
		q = getNext(getNext(q))
	}
	return q == 1
}

func getNext(n int) int {
	x := 0
	for n > 0 {
		x += (n % 10) * (n % 10)
		n = n / 10
	}
	return x
}

//leetcode submit region end(Prohibit modification and deletion)
