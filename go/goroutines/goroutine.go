/**
 * Author: Wang P
 * Version: 1.0.0
 * Date: 2021/2/5 下午1:43
 * Description:
 **/

package goroutines

import (
	"fmt"
	"time"
)

func run() {
	for i := 0; i < 100; i++ {
		go func(j int) {
			for {
				fmt.Println("go run from", j)
			}
		}(i)
	}
	time.Sleep(time.Minute)
}
