/**
 * Author: Wang P
 * Version: 1.0.0
 * Date: 2021/2/1 下午4:28
 * Description: 变量学习笔记
 **/

package main

import (
	"fmt"
	"math"
	"math/cmplx"
)

func euler() {
	// 欧拉公式 e^(i*Pi) + 1 = 0
	fmt.Println(cmplx.Exp(1i*math.Pi) + 1)
	fmt.Println(cmplx.Pow(math.E, 1i*math.Pi) + 1)
}

func triangle() {
	// 强制类型转换
	var a, b int = 3, 4
	var c int
	c = int(math.Sqrt(float64(a*a + b*b)))
	fmt.Println(c)
}

func consts() {
	// 常量未声明类型，其类型是不确定的（数值可以作各种类型使用）
	const a, b = 3, 4
	var c int
	c = int(math.Sqrt(a*a + b*b))
	fmt.Println(c)
}

func enums() {
	// 基于常量声明的枚举
	const (
		golang = iota
		_
		java
		python
	)

	const (
		b = 1 << (10 * iota)
		kb
		mb
		gb
		tb
	)
	fmt.Println(golang, java, python)
	fmt.Println(b, kb, mb, gb, tb)
}
