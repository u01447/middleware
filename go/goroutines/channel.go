/**
 * Author: Wang P
 * Version: 1.0.0
 * Date: 2021/2/5 下午1:57
 * Description:
 **/

package goroutines

import (
	"fmt"
	"time"
)

func works(id int, c chan int) {
	// 不断的从channel取
	for {
		fmt.Printf("channel %d receive %c\n", id, <-c)
	}
}

func workIfNotClose(id int, c chan int) {
	// 接收方判断channel中有数据就不断的从channel取
	for {
		n, ok := <-c
		if !ok {
			break
		}
		fmt.Printf("channel %d receive %c\n", id, n)
	}
}

func workIfNotCloseSimple(id int, c chan int) {
	// 接收方判断channel中有数据就不断的从channel取
	for n := range c {
		fmt.Printf("channel %d receive %c\n", id, n)
	}
}

func createWorks(id int) chan<- int {
	c := make(chan int)
	go works(id, c)
	return c
}

func first() {
	// 创建channel用例
	var chans [10]chan int
	for i := 0; i < 10; i++ {
		chans[i] = make(chan int)
		go works(i, chans[i])
	}
	// 往channel放
	for i := 0; i < 10; i++ {
		chans[i] <- 'a' + i
	}
	// 往channel放
	for i := 0; i < 10; i++ {
		chans[i] <- 'A' + i
	}
}

func second() {
	// 只能发数据  chan<- int 将channel 作为函数返回参数  即 channel也是一等公民
	var chann [10]chan<- int
	for i := 0; i < 10; i++ {
		chann[i] = createWorks(i)
	}

	for i := 0; i < 10; i++ {
		chann[i] <- 'k' + i
	}

	for i := 0; i < 10; i++ {
		chann[i] <- 'K' + i
	}
}

func thirdBufferedChannel() {
	// channel缓冲区   可以放入不大于缓冲区的大小时，可以不用取
	c := make(chan int, 4)
	go works(0, c)
	c <- 'u'
	c <- 'v'
	c <- 'w'
	c <- 'x'
	c <- 'y'
}

func channelClose() {
	c := make(chan int, 4)
	go workIfNotCloseSimple(0, c)
	c <- 'U'
	c <- 'V'
	c <- 'W'
	close(c)
}

func chanDemo() {

	//  channel也是一等公民
	first()
	second()
	fmt.Println("-----------------")
	thirdBufferedChannel()
	channelClose()
	time.Sleep(time.Second)
}
