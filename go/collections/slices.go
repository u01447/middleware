/**
 * Author: Wang P
 * Version: 1.0.0
 * Date: 2021/2/2 下午2:29
 * Description:
 **/

package collections

import (
	"fmt"
)

func SliceAppend() {
	arr := [...]int{1, 2, 3, 4, 5, 6, 7, 8}
	s1 := arr[2:6]
	fmt.Println(s1, cap(s1), len(s1))
	s2 := s1[3:5]
	fmt.Println(s2)
	s3 := append(s2, 10)
	s4 := append(s3, 11)
	s5 := append(s4, 12)
	fmt.Println("s3, s4, s5=", s3, s4, s5)
	fmt.Println("cap(s3), cap(s4), cap(s5) =", cap(s3), cap(s4), cap(s5))
	fmt.Println(arr)

	s6 := append(s1, 10)
	s7 := append(s6, 11)
	s8 := append(s7, 12)
	fmt.Println("s6, s7, s8=", s6, s7, s8)
	fmt.Println("cap(s6), cap(s7), cap(s8) =", cap(s6), cap(s7), cap(s8))
	fmt.Println(arr)
}

func SliceDefine() {
	var s []int
	for i := 0; i < 10; i++ {
		fmt.Printf("%v, cap(s) = %d, len(s) = %d\n", s, cap(s), len(s))
		s = append(s, 2*i+1)
	}

	s1 := []int{2, 3, 4}
	fmt.Printf("%v, cap(s1) = %d, len(s1) = %d\n", s1, cap(s1), len(s1))

	s2 := make([]int, 8)
	fmt.Printf("%v, cap(s2) = %d, len(s2) = %d\n", s2, cap(s2), len(s2))

	s3 := make([]int, 8, 32)
	fmt.Printf("%v, cap(s3) = %d, len(s3) = %d\n", s3, cap(s3), len(s3))

	// 拷贝
	copy(s2, s1)
	fmt.Printf("%v, cap(s2) = %d, len(s2) = %d\n", s2, cap(s2), len(s2))

	// 删除 没有内建函数，只能通过截取+append
	s4 := append(s2[:2], s2[3:]...)
	fmt.Printf("%v, cap(s4) = %d, len(s4) = %d\n", s4, cap(s4), len(s4))
}
