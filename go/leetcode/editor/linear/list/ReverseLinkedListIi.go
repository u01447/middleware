package leetcode

/**
 * Author: Wang P
 * Version: 1.0.0
 * Date: 2021-03-05 19:07:45
 * Description: //Given the head of a singly linked list and two integers left and right where left <= right, reverse
                //the nodes of the list from position left to position right, and return the reversed list.
				//
				// Example 1:
				//Input: head = [1,2,3,4,5], left = 2, right = 4
				//Output: [1,4,3,2,5]
				//
				// Example 2:
				//Input: head = [5], left = 1, right = 1
				//Output: [5]
				//
				// Constraints:
				// The number of nodes in the list is n.
				// 1 <= n <= 500
				// -500 <= Node.val <= 500
				// 1 <= left <= right <= n
				//Follow up: Could you do it in one pass? Related Topics 链表
				// 👍 709 👎 0

 **/
//leetcode submit region begin(Prohibit modification and deletion)
/**
 * Definition for singly-linked list.
 */
type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseBetween(head *ListNode, left int, right int) *ListNode {
	hair := &ListNode{Next: head}
	cur := hair
	n := right - left + 1
	for left > 1 {
		cur = cur.Next
		left--
	}
	cur.Next = reverse(cur.Next, n)
	return hair.Next
}

func reverse(head *ListNode, n int) *ListNode {
	pre, next := &ListNode{}, &ListNode{}
	cur := head

	for n > 0 {
		next = cur.Next
		cur.Next = pre.Next
		pre.Next = cur
		cur = next
		n--
	}
	head.Next = cur
	return pre.Next
}

//leetcode submit region end(Prohibit modification and deletion)
