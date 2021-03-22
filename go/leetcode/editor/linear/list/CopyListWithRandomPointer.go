package leetcode

/**
 * Author: Wang P
 * Version: 1.0.0
 * Date: 2021-03-06 20:37:49
 * Description: A linked list of length n is given such that each node contains an additional random pointer, which could point to any node in the list, or null.

				Construct a deep copy of the list. The deep copy should consist of exactly n brand new nodes, where each new node has its value set to the value of its corre
				sponding original node. Both the next and random pointer of the new nodes should point to new nodes in the copied list such that the pointers in the original li
				st and copied list represent the same list state. None of the pointers in the new list should point to nodes in the original list.

				For example, if there are two nodes X and Y in the original list, where X.random --> Y, then for the corresponding two nodes x and y in the copied list, x.random --> y.

				Return the head of the copied linked list.

				The linked list is represented in the input/output as a list of n nodes. Each node is represented as a pair of [val, random_index] where:

				val: an integer representing Node.val
				random_index: the index of the node (range from 0 to n-1) that the random pointer points to, or null if it does not point to any node.

				Your code will only be given the head of the original linked list.

				Example 1:
				Input: head = [[7,null],[13,0],[11,4],[10,2],[1,0]]
				Output: [[7,null],[13,0],[11,4],[10,2],[1,0]]

				Example 2:
				Input: head = [[1,1],[2,1]]
				Output: [[1,1],[2,1]]

				Example 3:
				Input: head = [[3,null],[3,0],[3,null]]
				Output: [[3,null],[3,0],[3,null]]

				Example 4:
				Input: head = []
				Output: []
				Explanation: The given linked list is empty (null pointer), so return null.

				Constraints:
				0 <= n <= 1000
				-10000 <= Node.val <= 10000
				Node.random is null or is pointing to some node in the linked list.

				Related Topics 哈希表 链表
				👍 519 👎 0
 **/

//leetcode submit region begin(Prohibit modification and deletion)
/**
 * Definition for a Node.
 */
type Node struct {
	Val    int
	Next   *Node
	Random *Node
}

func copyRandomList(head *Node) *Node {
	if head == nil {
		return nil
	}
	cur := head
	for cur != nil {
		n := &Node{Val: cur.Val, Next: cur.Next}
		cur.Next = n
		cur = n.Next
	}
	cur = head
	for cur != nil && cur.Next != nil {
		if cur.Random != nil {
			cur.Next.Random = cur.Random.Next
		} else {
			cur.Next.Random = nil
		}
		cur = cur.Next.Next
	}
	cur = head
	newNode, next := head.Next, head.Next
	for cur != nil && cur.Next != nil {
		cur.Next = cur.Next.Next
		if next.Next != nil {
			next.Next = next.Next.Next
		} else {
			next.Next = nil
		}
		cur = cur.Next
		next = next.Next
	}
	return newNode
}

//leetcode submit region end(Prohibit modification and deletion)
