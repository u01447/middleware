/**
 * Author: Wang P
 * Version: 1.0.0
 * Date: 2021/2/20 下午4:52
 * Description: 函数闭包相关学习笔记
 **/

package functions

import "fmt"

func adder() func(int) int {
	sum := 0
	return func(i int) int {
		sum += i
		return sum
	}
}

// 正统函数式编程
type iAdder func(int) (int, iAdder)

func iAdd(base int) iAdder {
	return func(v int) (int, iAdder) {
		return base + v, iAdd(base + v)
	}
}

func TestAdder() {
	a := adder()
	for i := 0; i < 10; i++ {
		fmt.Printf("0 + ... + %d = %d \n", i, a(i))
	}

	a2 := iAdd(0)
	var s int
	for i := 1; i < 10; i++ {
		s, a2 = a2(i)
		fmt.Printf("0 + ... + %d = %d \n", i, s)
	}
}
