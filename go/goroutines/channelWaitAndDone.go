/**
 * Author: Wang P
 * Version: 1.0.0
 * Date: 2021/2/5 下午5:37
 * Description: 基于WaitGroup实现通信来共享内存
 **/

package goroutines

import (
	"fmt"
	"sync"
)

type WorkerWG struct {
	in chan int
	wg *sync.WaitGroup
}

func doWorkWaitDone(id int, w WorkerWG) {
	// 接收方判断channel中有数据就不断的从channel取
	for n := range w.in {
		fmt.Printf("channel %d receive %c\n", id, n)
		w.wg.Done()
	}
}

func createWaitDoneWorkers(id int, wg *sync.WaitGroup) WorkerWG {
	w := WorkerWG{
		in: make(chan int),
		wg: wg,
	}
	go doWorkWaitDone(id, w)
	return w
}

func communicateWaitDone() {
	// 只能发数据  chan<- int 将channel 作为函数返回参数  即 channel也是一等公民
	var works [10]WorkerWG
	var wg sync.WaitGroup
	for i := 0; i < 10; i++ {
		works[i] = createWaitDoneWorkers(i, &wg)
	}

	//wg.Add(20)

	for i, worker := range works {
		worker.in <- 'a' + i
		wg.Add(1)
	}

	for i, worker := range works {
		worker.in <- 'A' + i
		wg.Add(1)
	}

	wg.Wait()
}

type WorkerWG2 struct {
	in   chan int
	done func() // 函数式编程
}

// WorkerWG封装
func doWorkWaitDone2(id int, w WorkerWG2) {
	// 接收方判断channel中有数据就不断的从channel取
	for n := range w.in {
		fmt.Printf("channel %d receive %c\n", id, n)
		w.done()
	}
}

func createWaitDoneWorkers2(id int, wg *sync.WaitGroup) WorkerWG2 {
	w := WorkerWG2{
		in: make(chan int),
		done: func() {
			wg.Done()
		},
	}
	go doWorkWaitDone2(id, w)
	return w
}

func communicateWaitDone2() {
	// 只能发数据  chan<- int 将channel 作为函数返回参数  即 channel也是一等公民
	var works [10]WorkerWG2
	var wg sync.WaitGroup
	for i := 0; i < 10; i++ {
		works[i] = createWaitDoneWorkers2(i, &wg)
	}

	wg.Add(20)
	for i, worker := range works {
		worker.in <- 'a' + i
		//wg.Add(1)
	}
	for i, worker := range works {
		worker.in <- 'A' + i
		//wg.Add(1)
	}
	wg.Wait()
}
