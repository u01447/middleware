/**
 * Author: Wang P
 * Version: 1.0.0
 * Date: 2021/2/5 下午6:02
 * Description: travelsWithFunc函数式编程遍历树
				travelWithChannel 利用channel遍历树
 **/

package goroutines

import "fmt"

type Node struct {
	Value       int
	Left, Right *Node
}

func createNode(value int) *Node {
	// 创建节点
	return &Node{Value: value}
}

func (node *Node) setValue(value int) {
	// 接收者使用指针才可以改变结构内容
	if node == nil {
		fmt.Println("Setting Value to nil node. Ignored.")
		return
	}
	node.Value = value
}

func (node *Node) print() {
	// 节点数据打印
	if node == nil {
		return
	}
	fmt.Print(node.Value, " ")
}

func (node *Node) travelsWithFunc(f func(*Node)) {
	// 函数式编程中序遍历树
	if node == nil {
		return
	}
	node.Left.travelsWithFunc(f)
	f(node)
	node.Right.travelsWithFunc(f)
}

func (node *Node) travelWithChannel() chan *Node {
	// 利用channel遍历树
	c := make(chan *Node)
	go func() {
		node.travelsWithFunc(func(node *Node) {
			c <- node
		})
		close(c)
	}()
	return c
}

func initTree() *Node {
	// 初始化树
	var root Node
	root = Node{Value: 3}
	root.Left = &Node{}
	root.Right = &Node{5, nil, nil}
	root.Right.Left = new(Node)
	root.Left.Right = createNode(2)
	root.Right.Left.setValue(4)
	return &root
}

func travelsTree() {
	t := initTree()

	// 计数
	nodeCount := 0
	t.travelsWithFunc(func(node *Node) {
		nodeCount++
		fmt.Print(node.Value, " ")
	})
	fmt.Println("Node counts", nodeCount)

	// 树中最大值
	c := t.travelWithChannel()
	maxNode := 0
	// 从channel中取
	for n := range c {
		if maxNode < n.Value {
			maxNode = n.Value
		}
	}
	fmt.Println("Max node:", maxNode)
}
