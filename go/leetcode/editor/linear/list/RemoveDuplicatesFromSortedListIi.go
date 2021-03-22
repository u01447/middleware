package leetcode

/**
 * Author: Wang P
 * Version: 1.0.0
 * Date: 2021-03-06 13:50:38
 * Description: //Given the head of a sorted linked list, delete all nodes that have duplicate n
				//umbers, leaving only distinct numbers from the original list. Return the linked
				//list sorted as well.
				// Example 1:
				//Input: head = [1,2,3,3,4,4,5]
				//Output: [1,2,5]
				//
				// Example 2:
				//Input: head = [1,1,1,2,3]
				//Output: [2,3]
				//
				// Constraints:
				// The number of nodes in the list is in the range [0, 300].
				// -100 <= Node.val <= 100
				// The list is guaranteed to be sorted in ascending order.
				//
				// Related Topics 链表
				// 👍 469 👎 0
 **/

//leetcode submit region begin(Prohibit modification and deletion)
/**
 * Definition for singly-linked list.
 */
//type ListNode struct {
//	Val int
//	Next *ListNode
//}

func deleteDuplicates2(head *ListNode) *ListNode {
	hair := &ListNode{Next: head}
	fast, slow := head, hair
	for fast != nil {
		for fast.Next != nil && fast.Val == fast.Next.Val {
			fast = fast.Next
		}
		if slow.Next == fast {
			slow = slow.Next
		} else {
			slow.Next = fast.Next
		}
		fast = fast.Next
	}
	return hair.Next
}

//leetcode submit region end(Prohibit modification and deletion)
