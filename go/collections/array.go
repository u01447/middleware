/**
 * Author: Wang P
 * Version: 1.0.0
 * Date: 2021/2/2 下午1:46
 * Description:
 **/

package collections

import (
	"fmt"
)

func Define() {
	var arr1 [5]int
	arr2 := [3]int{1, 3, 5}
	arr3 := [...]int{2, 4, 6, 8, 10}

	var grid [4][5]bool
	fmt.Println(arr1, arr2, arr3)
	fmt.Println(grid)
}

func printArr(arr [5]int) {
	arr[0] = 100
	for i, v := range arr {
		fmt.Println(i, v)
	}
}

func printArrR(arr *[5]int) {
	arr[0] = 100
	for i, v := range arr {
		fmt.Println(i, v)
	}
}

func ArrTest() {
	var arr1 [5]int
	arr2 := [...]int{2, 4, 6, 8, 10}
	printArr(arr1)
	fmt.Println()
	printArr(arr2)
	fmt.Println()
	fmt.Println(arr1)
	fmt.Println()
	printArrR(&arr1)
	fmt.Println()
	fmt.Println(arr1)
}
