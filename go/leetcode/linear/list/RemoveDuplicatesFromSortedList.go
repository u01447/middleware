package list

/**
 * Author: Wang P
 * Version: 1.0.0
 * Date: 2021-03-06 13:50:35
 * Description: //Given the head of a sorted linked list, delete all duplicates such that each e
				//lement appears only once. Return the linked list sorted as well.
				//
				// Example 1:
				//Input: head = [1,1,2]
				//Output: [1,2]
				//
				// Example 2:
				//Input: head = [1,1,2,3,3]
				//Output: [1,2,3]
				//
				// Constraints:
				// The number of nodes in the list is in the range [0, 300].
				// -100 <= Node.val <= 100
				// The list is guaranteed to be sorted in ascending order.
				//
				// Related Topics 链表
				// 👍 488 👎 0
 **/

//leetcode submit region begin(Prohibit modification and deletion)
/**
 * Definition for singly-linked list.
 */
//type ListNode struct {
//	Val int
//	Next *ListNode
//}

func deleteDuplicates1(head *ListNode) *ListNode {
	cur := head
	for cur != nil && cur.Next != nil {
		if cur.Val == cur.Next.Val {
			cur.Next = cur.Next.Next
		} else {
			cur = cur.Next
		}
	}
	return head
}

//leetcode submit region end(Prohibit modification and deletion)
