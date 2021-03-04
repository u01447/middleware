/**
 * Author: Wang P
 * Version: 1.0.0
 * Date: 2021/3/2 下午8:13
 * Description:
 **/

package animal

import (
	"fmt"
)

func ExampleQueue_Pop() {
	q := Queue{}
	_ = q.Push("asd")
	_ = q.Push(123)
	if v, ok := q.Pop(); ok {
		fmt.Println(v)
	}
	if v, ok := q.Pop(); ok {
		fmt.Println(v)
	}

	//Output:
	//asd
	//123
}
