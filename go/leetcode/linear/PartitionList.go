package leetcode

/**
 * Author: Wang P
 * Version: 1.0.0
 * Date: 2021-03-06 20:37:43
 * Description: Given the head of a linked list and a value x, partition it such that all nodes less than x come before nodes greater than or equal to x.
				You should preserve the original relative order of the nodes in each of the two partitions.
				Example 1:
				Input: head = [1,4,3,2,5,2], x = 3
				Output: [1,2,2,4,3,5]

				Example 2:
				Input: head = [2,1], x = 2
				Output: [1,2]

				Constraints:
				 The number of nodes in the list is in the range [0, 200].
				 -100 <= Node.val <= 100
				 -200 <= x <= 200

				 Related Topics 链表 双指针
				 👍 378 👎 0
 **/
//leetcode submit region begin(Prohibit modification and deletion)
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func partition(head *ListNode, x int) *ListNode {
	lesser, greater := &ListNode{Next: head}, &ListNode{Next: head}
	less, great := lesser, greater
	for head != nil {
		if head.Val < x {
			less.Next = head
			less = less.Next
		} else {
			great.Next = head
			great = great.Next
		}
		head = head.Next
	}
	great.Next = nil
	less.Next = greater.Next

	return lesser.Next
}

//leetcode submit region end(Prohibit modification and deletion)
