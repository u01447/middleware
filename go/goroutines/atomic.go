/**
 * Author: Wang P
 * Version: 1.0.0
 * Date: 2021/2/5 下午2:34
 * Description:
 **/

package goroutines

import (
	"fmt"
	"sync"
	"time"
)

type AtomicInt struct {
	value int
	lock  sync.Mutex
}

func (a *AtomicInt) increase() {
	fmt.Println("safe Increase")
	func() {
		a.lock.Lock()
		defer a.lock.Unlock()
		a.value++
	}()
}

func (a *AtomicInt) get() int {
	a.lock.Lock()
	defer a.lock.Unlock()
	return a.value
}

func AtomicDemo() {
	i := AtomicInt{}
	i.increase()
	go func() {
		i.increase()
	}()
	time.Sleep(time.Millisecond)
	fmt.Println(i.get())
}
