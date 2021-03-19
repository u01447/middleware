/*
@Time : 2020/8/10 10:40
@Author : Wang P
@File : Solution
*/
package main

/**
  判断一个整数是否是回文数。回文数是指正序（从左向右）和倒序（从右向左）读都是一样的整数。

   示例 1:
   输入: 121
   输出: true

   示例 2:
   输入: -121
   输出: false
   解释: 从左向右读, 为 -121 。 从右向左读, 为 121- 。因此它不是一个回文数。

   示例 3:
   输入: 10
   输出: false
   解释: 从右向左读, 为 01 。因此它不是一个回文数。

   进阶:
   你能不将整数转为字符串来解决这个问题吗？
   Related Topics 数学
   👍 1174 👎 0
*/
func isPalindrome(x int) bool {
	if x < 0 || x%10 == 0 {
		return false
	}
	r := 0
	for x > r {
		r = r*10 + x%10
		x = x / 10
	}
	return x == r || x == r/10 || x == 0
}

func main() {

}
