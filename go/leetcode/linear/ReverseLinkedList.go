package leetcode

/**
 * Author: Wang P
 * Version: 1.0.0
 * Date: 2021-03-06 17:46:07
 * Description: //Given the head of a singly linked list, reverse the list, and return the reversed list.
				// Example 1:
				//Input: head = [1,2,3,4,5]
				//Output: [5,4,3,2,1]
				//
				// Example 2:
				//Input: head = [1,2]
				//Output: [2,1]
				//
				// Example 3:
				//Input: head = []
				//Output: []
				//
				// Constraints:
				// The number of nodes in the list is the range [0, 5000].
				// -5000 <= Node.val <= 5000
				//
				// Follow up: A linked list can be reversed either iteratively or recursively. C
				//ould you implement both?
				// Related Topics 链表
				// 👍 1568 👎 0
 **/
//leetcode submit region begin(Prohibit modification and deletion)
/**
 * Definition for singly-linked list.
 */
//type ListNode struct {
//	Val int
//	Next *ListNode
//}

func reverseList(head *ListNode) *ListNode {
	var pre, cur, next *ListNode = nil, head, nil
	for cur != nil {
		next = cur.Next
		cur.Next = pre
		pre = cur
		cur = next
	}
	return pre
}

//leetcode submit region end(Prohibit modification and deletion)
