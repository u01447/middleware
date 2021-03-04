/**
 * Author: Wang P
 * Version: 1.0.0
 * Date: 2021/2/1 上午11:57
 * Description:
 **/

package main

import "offer/note/goroutines"

func main() {

	// 02 函数相关学习用例
	//fmt.Println(apply(pow, 3, 4))
	//// 匿名函数方式
	//fmt.Println(apply(func(f, f2 int) float64 {
	//	return math.Pow(float64(f), float64(f2))
	//}, 3, 4))
	//sum(1, 2, 3, 4, 5)

	////03 指针相关学习用例
	//a, b := 3, 4
	//swapR(&a, &b)
	//fmt.Println(a, b)

	//// 04 数组相关学习
	//collections.Define()
	//collections.ArrTest()
	//collections.SliceAppend()
	//collections.SliceDefine()

	//// 04 map相关学习
	//collections.MapDefine()

	//// 05 rune使用
	//fmt.Println(collections.MaxNoRepeated("abcadcb"))
	//fmt.Println(collections.MaxNoRepeatedChn("黑化肥挥发发灰会花飞灰化肥挥发发黑会飞花"))
	//collections.StrByte("")

	//node := object.Node{}
	//node.Print()
	//node.SetValueNotUpdate(10)
	//node.Print()
	//node.SetValue(10)
	//node.Print()
	//var pNode *object.Node
	//pNode.SetValue(30)
	//pNode.SetValueNotUpdate(20)

	// 07 接口
	//animal.FindDuck(&mock.Duck{})
	//animal.FindDuck(&mock.Duck{Name: "我是一只假鸭子"})
	//
	//q := animal.Queue{}
	//_ = q.Push("asd")
	//_ = q.Push(123)
	//if v, ok := q.Pop(); ok {
	//	fmt.Println(v)
	//}
	//if v, ok := q.Pop(); ok {
	//	fmt.Println(v)
	//}
	//animal.DuckBehavior(&mock.Duck{})

	//// 08 高阶函数/闭包
	//functions.TestAdder()

	//// 09 错误处理与单元测试在简单web服务中的应用
	//web.Server()

	// 10 goroutine
	//goroutines.Run()

	goroutines.AtomicDemo()
}
