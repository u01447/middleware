/**
 * Author: Wang P
 * Version: 1.0.0
 * Date: 2021/2/5 下午5:05
 * Description:
 **/

package goroutines

import (
	"fmt"
)

type Worker struct {
	in   chan int
	done chan bool
}

func doWork(id int, w Worker) {
	// 接收方判断channel中有数据就不断的从channel取
	for n := range w.in {
		fmt.Printf("channel %d receive %c\n", id, n)
		w.done <- true
	}
}

func createWorkers(id int) Worker {
	w := Worker{
		in:   make(chan int),
		done: make(chan bool),
	}
	go doWork(id, w)
	return w
}

func communicate() {
	// 只能发数据  chan<- int 将channel 作为函数返回参数  即 channel也是一等公民
	var works [10]Worker
	for i := 0; i < 10; i++ {
		works[i] = createWorkers(i)
	}

	for i, worker := range works {
		worker.in <- 'a' + i
	}

	for _, worker := range works {
		<-worker.done
	}

	for i, worker := range works {
		worker.in <- 'A' + i
	}

	for _, worker := range works {
		<-worker.done
	}
}

func doWork2(id int, w Worker) {
	// 接收方判断channel中有数据就不断的从channel取
	for n := range w.in {
		fmt.Printf("channel %d receive %c\n", id, n)
		go func() {
			w.done <- true
		}()
	}
}

func createWorkers2(id int) Worker {
	w := Worker{
		in:   make(chan int),
		done: make(chan bool),
	}
	go doWork2(id, w)
	return w
}

func communicateDoneAll() {
	// 只能发数据  chan<- int 将channel 作为函数返回参数  即 channel也是一等公民
	var worksAll [10]Worker

	for i := 0; i < 10; i++ {
		worksAll[i] = createWorkers2(i)
	}

	for i, worker := range worksAll {
		worker.in <- 'j' + i
	}

	for i, worker := range worksAll {
		worker.in <- 'J' + i
	}

	// 等所有任务都取完 再结束
	for _, worker := range worksAll {
		<-worker.done
		<-worker.done
	}
}

func communicateDemo() {
	communicate()
	fmt.Println("----------------")
	communicateDoneAll()
}
