/**
 * Author: Wang P
 * Version: 1.0.0
 * Date: 2021/2/2 上午11:27
 * Description: 指针相关学习笔记
 **/

package main

import "fmt"

func pointer() {
	var a int = 2
	var pa *int = &a
	*pa = 3
	fmt.Println(a)
}

func swapV(a, b int) {
	b, a = a, b
}

func swapR(a, b *int) {
	*b, *a = *a, *b
}
