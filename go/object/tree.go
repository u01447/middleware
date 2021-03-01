/**
 * Author: Wang P
 * Version: 1.0.0
 * Date: 2021/2/19 下午12:49
 * Description:
 **/

package object

import "fmt"

type Node struct {
	Value       int
	Left, Right *Node
}

type MyNode struct {
	node *Node
}

func (node Node) Print() {
	if &node == nil {
		fmt.Println("node is nil")
		return
	}
	fmt.Print(node.Value, " ")
}

func (node *Node) SetValue(value int) {
	// 接收者使用指针才可以改变结构内容
	if node == nil {
		fmt.Println("Setting Value to nil node. Ignored.")
		return
	}
	node.Value = value
}

func (node Node) SetValueNotUpdate(value int) {
	// 值传递 node值无法改变
	fmt.Println(node)
	if &node == nil {
		fmt.Println("Setting Value to nil node. Ignored.")
		return
	}
	node.Value = value
}

func CreateNode(value int) *Node {
	return &Node{Value: value}
}
