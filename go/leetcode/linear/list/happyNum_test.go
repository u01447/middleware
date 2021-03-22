/**
 * Author: Wang P
 * Version: 1.0.0
 * Date: 2021/3/5 上午10:59
 * Description:
 **/

package list

import (
	"fmt"
	"testing"
)

func TestIsHappy(t *testing.T) {
	sum := 0
	for i := 0; i <= 100000; i++ {
		if IsHappy(i) {
			sum += i
		}
	}
	fmt.Println(sum)
	fmt.Println(IsHappy(19))
}
