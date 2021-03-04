/**
* Author: Wang P
* Version: 1.0.0
* Date: 2021/2/5 上午10:37
* Description: selectNotBlock  select非阻塞调度简单用例
               channelSelectBlock select阻塞调度简单用例
               channelSelectWork 多channel协助调度用例
**/

package goroutines

import (
	"fmt"
	"math/rand"
	"time"
)

type SelectWorker struct {
	id string
	c  chan int
}

func (w *SelectWorker) setValue(id string) {
	w.id = id
}

func selectNotBlock(c1, c2 chan int) {
	for {
		select {
		case n := <-c1:
			fmt.Printf("receice %d from c1\n", n)
		case n := <-c2:
			fmt.Printf("receice %d from c2\n", n)
		default:
			fmt.Println("not receive anything")
		}
	}
}

func selectBlock(c1, c2 chan int) {
	for {
		select {
		case n := <-c1:
			fmt.Printf("receice %d from c1\n", n)
		case n := <-c2:
			fmt.Printf("receice %d from c2\n", n)

		}
	}
}

func selectWorkConditionBlock(rw, sw1, sw2 SelectWorker) {
	var values []int
	id := ""
	endTime := time.After(time.Second * 10)
	tick := time.Tick(time.Second)
	for {
		activeWorker := SelectWorker{
			c: nil,
		}
		var activeValue int
		if len(values) > 0 {
			rw.setValue(id)
			activeWorker = rw
			activeValue = values[0]
		}
		select {
		case n := <-sw1.c:
			id = sw1.id
			values = append(values, n)
		case n := <-sw2.c:
			id = sw2.id
			values = append(values, n)
		case activeWorker.c <- activeValue:
			values = values[1:]
		case <-time.After(time.Millisecond * 600): // 500毫秒未产生数据
			fmt.Println("程序超时")
		case <-tick:
			fmt.Println("queue len = ", len(values))
		case <-endTime: // 程序执行到endTime时结束
			fmt.Println("程序执行结束")
			return
		}
	}
}

func generateChannel() chan int {
	c := make(chan int)
	go func() {
		i := 0
		for {
			time.Sleep(time.Duration(rand.Intn(1000)) * time.Millisecond)
			c <- i
			i++
		}
	}()
	return c
}

func generateWorkChannel(id string) SelectWorker {
	w := SelectWorker{
		id: id,
		c:  make(chan int),
	}
	go func() {
		i := 0
		for {
			time.Sleep(time.Duration(rand.Intn(1000)) * time.Millisecond)
			w.c <- i
			i++
		}
	}()
	return w
}

func selectWorks(id string, rw SelectWorker) {
	// 不断的从channel取
	for {
		fmt.Printf("channel %s from %s receive %d\n", id, rw.id, <-rw.c)
		time.Sleep(time.Second)
	}
}

func createSelectWorks(id string) SelectWorker {
	rw := SelectWorker{
		c: make(chan int),
	}
	go selectWorks(id, rw)
	return rw
}

func channelSelectBlock() {
	var c1, c2 = generateChannel(), generateChannel()
	selectBlock(c1, c2)
}

func channelSelectWork() {
	var sw1, sw2 = generateWorkChannel("send worker1"), generateWorkChannel("send worker2")
	rw := createSelectWorks("receive work1")
	selectWorkConditionBlock(rw, sw1, sw2)
}

func selectDemo() {

	//var c1, c2 chan int
	//selectNotBlock(c1, c2)

	//channelSelectBlock()

	channelSelectWork()
}
