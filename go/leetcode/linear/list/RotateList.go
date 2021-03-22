package list

/**
 * Author: Wang P
 * Version: 1.0.0
 * Date: 2021-03-06 20:26:17
 * Description:  Given the head of a linked list, rotate the list to the right by k places.

				 Example 1:
				Input: head = [1,2,3,4,5], k = 2
				Output: [4,5,1,2,3]

				 Example 2:
				Input: head = [0,1,2], k = 4
				Output: [2,0,1]

				 Constraints:
				 The number of nodes in the list is in the range [0, 500].
				 -100 <= Node.val <= 100
				 0 <= k <= 2 * 109

				 Related Topics 链表 双指针
				 👍 444 👎 0
 **/
//leetcode submit region begin(Prohibit modification and deletion)
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func rotateRight(head *ListNode, k int) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	oldTail := head
	length := 1
	for oldTail.Next != nil {
		oldTail = oldTail.Next
		length++
	}
	oldTail.Next = head
	newTail := head
	for i := 0; i < length-k%length-1; i++ {
		newTail = newTail.Next
	}
	head = newTail.Next
	newTail.Next = nil
	return head
}

//leetcode submit region end(Prohibit modification and deletion)
