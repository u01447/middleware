/**
 * Author: Wang P
 * Version: 1.0.0
 * Date: 2021/2/18 上午10:39
 * Description: 函数相关学习笔记
 **/

package functions

import (
	"fmt"
	"math"
	"reflect"
	"runtime"
)

func eval(a, b int, op string) (int, error) {
	switch op {
	case "+":
		return a + b, nil
	case "-":
		return a - b, nil
	case "*":
		return a * b, nil
	case "/":
		r, _ := div(a, b)
		return r, nil
	default:
		return 0, fmt.Errorf("unsupported operation")
	}
}

func div(a, b int) (q, r int) {
	return a / b, a % b
}

func apply(op func(int, int) float64, a, b int) float64 {
	//
	p := reflect.ValueOf(op).Pointer()
	opName := runtime.FuncForPC(p).Name()
	fmt.Printf("Calling function %s with params (%d, %d)\n", opName, a, b)
	return op(int(a), int(b))
}

func pow(a, b int) float64 {
	return math.Pow(float64(a), float64(b))
}

func sum(nums ...int) {
	// 函数可变参数列表
	sum := 0
	for _, v := range nums {
		sum += v
	}
	fmt.Println(sum)
}
