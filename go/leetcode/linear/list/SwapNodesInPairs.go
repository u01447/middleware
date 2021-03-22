package list

/**
 * Author: Wang P
 * Version: 1.0.0
 * Date: 2021-03-06 20:26:56
 * Description:  Given a linked list, swap every two adjacent nodes and return its head.

				 Example 1:
				Input: head = [1,2,3,4]
				Output: [2,1,4,3]

				 Example 2:
				Input: head = []
				Output: []

				 Example 3:
				Input: head = [1]
				Output: [1]

				 Constraints:
				 The number of nodes in the list is in the range [0, 100].
				 0 <= Node.val <= 100

				Follow up: Can you solve the problem without modifying the values in the list'
				s nodes? (i.e., Only nodes themselves may be changed.) Related Topics 递归 链表
				 👍 859 👎 0
 **/
//leetcode submit region begin(Prohibit modification and deletion)
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func swapPairs(head *ListNode) *ListNode {
	hair := &ListNode{Next: head}
	pre, tail := hair, hair
	for head != nil {
		for i := 0; i < 2; i++ {
			tail = tail.Next
			if tail == nil {
				return hair.Next
			}
		}

		head, tail = reverseN(head, tail)
		pre.Next = head
		pre = tail
		head = pre.Next
	}
	return hair.Next
}

func reverseN(head, tail *ListNode) (*ListNode, *ListNode) {

	pre, cur, next := tail.Next, head, &ListNode{}
	for pre != tail {
		next = cur.Next
		cur.Next = pre
		pre = cur
		cur = next
	}

	return tail, head
}

//leetcode submit region end(Prohibit modification and deletion)
