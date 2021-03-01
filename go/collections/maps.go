/**
 * Author: Wang P
 * Version: 1.0.0
 * Date: 2021/2/2 下午4:14
 * Description:
 **/

package collections

import "fmt"

func MapDefine() {
	m := map[int]string{
		1: "aa",
		2: "bb",
	}
	if v, ok := m[3]; ok {
		fmt.Println(v)
	} else {
		panic("key not exists")
	}
}
