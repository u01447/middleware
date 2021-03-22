package list

/**
 * Author: Wang P
 * Version: 1.0.0
 * Date: 2021-03-06 13:43:33
 * Description:
 **/
//Given the head of a linked list, remove the nth node from the end of the list
//and return its head.
//
// Follow up: Could you do this in one pass?
//
// Example 1:
//Input: head = [1,2,3,4,5], n = 2
//Output: [1,2,3,5]
//
// Example 2:
//Input: head = [1], n = 1
//Output: []
//
// Example 3:
//Input: head = [1,2], n = 1
//Output: [1]

// Constraints:
// The number of nodes in the list is sz.
// 1 <= sz <= 30
// 0 <= Node.val <= 100
// 1 <= n <= sz
//
// Related Topics 链表 双指针
// 👍 1264 👎 0

//leetcode submit region begin(Prohibit modification and deletion)
/**
 * Definition for singly-linked list.
 */
//type ListNode struct {
// 	Val int
// 	Next *ListNode
//}

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	hair := &ListNode{Next: head}
	fast, slow := head, hair
	for n > 0 {
		fast = fast.Next
		n--
	}
	for fast != nil {
		fast = fast.Next
		slow = slow.Next
	}
	slow.Next = slow.Next.Next
	return hair.Next
}

//leetcode submit region end(Prohibit modification and deletion)
