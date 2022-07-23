// author: kango2gler@gmail.com
// date: 2018-03-20

// 线程包.
//
// 线程操作类
// func:
//	ThreadsType.Threads() (map[int]*Thread) //获取当前线程集合
//	ThreadsType.addThread(*Thread) //添加线程到线程集合
//	ThreadsType.removeThread(*Thread) //从线程集合移除线程
//	ThreadsType.getNextThreadNum() //从线程集合中获取新的线程id
package thread

import (
	"errors"
	"fmt"
	"log"
	"sync"
	"time"

	"hummingbird/libs/trycatch"
)

//
// 线程操作类
// var t =  &Thread{Name:'',Async:true,Run:func f(){},Callback func fun(){}}
// t.Start()
//

//线程集合
type threadsCollector struct {
	mutex       *sync.Mutex //同步锁
	threadCount int         //当前线程计数器
	threads     map[int]*Thread
}

//
// 线程集合
var ThreadsCollector *threadsCollector

//
// 初始化线程集合
//
func init() {
	//线程集合
	ThreadsCollector = &threadsCollector{new(sync.Mutex), 0, make(map[int]*Thread)}
}

//
// 添加线程
//
func (tt *threadsCollector) addThread(t *Thread) *Thread {
	tt.mutex.Lock()
	tt.threads[t.id] = t
	t.state = THREAD_RUNING
	tt.mutex.Unlock()
	return t
}

//
// 删除线程
//
func (tt *threadsCollector) removeThread(id int) {
	tt.mutex.Lock()
	delete(tt.threads, id)
	tt.mutex.Unlock()
}

//
// 获取线程id,自增
//
func (tt *threadsCollector) getNextThreadNum() int {
	tt.mutex.Lock()
	tt.threadCount++
	tt.mutex.Unlock()
	return tt.threadCount
}

//
// 等待所有线程结束
//
func (tt *threadsCollector) WaitThreadsEnd() {
	log.Println(len(Threads()))
	for len(Threads()) > 0 {
	}
}

//
// 获取当前线程集合
//
func Threads() map[int]*Thread {
	return ThreadsCollector.threads
}

//
// 定义可以执行的线程接口
//
type ThreadInterface interface {
	Run()
}

//
// 线程类结构体
//
type Thread struct {
	id         int                    //自定义的线程id,不可赋值,只能获取,创建线程时自动生产序列
	localParam map[string]interface{} //本地变量
	state      byte                   //线程状态值, THREAD_RUNING,THREAD_INTERRUPTED
	Name       string                 //线程名称
	Async      bool                   //是否是异步
	Run        func(t *Thread)        // 线程执行的方法,第一个参数为当前线程对象,可设置当前线程中的localParam,
	// Run方法中可使用 for{} 来保证线程一直执行
	// 并且可以在for{} 中通过判断 t.getState() == THREAD_INTERRUPTED 来终止线程运行
	Callback func(t *Thread) //Run执行完的回调函数,第一个参数为当前线程对象,可设置当前线程中的localParam
}

//
// 线程常量
//
const (
	//状态: 运行中
	THREAD_RUNING = byte(1)
	//状态: 被截断
	THREAD_INTERRUPTED = byte(2)
)

// 获取当前线程id
func (t *Thread) getId() int {
	return t.id
}

//
//获取当前线程状态
//状态值有 THREAD_RUNING,THREAD_INTERRUPTED
//可在Run(t *Thread) 中通过判断 t.getState() == THREAD_INTERRUPTED 来终止线程运行
//
func (t *Thread) getState() byte {
	return t.state
}

//
// 线程方法执行
//
func (t *Thread) Start() error {
	if t.Run == nil {
		return errors.New("thread start failed, Run mathod must be implement")
	}
	trycatch.Trycatch(func() {
		if t.Async != true {
			waitingChain := make(chan int)
			go syncThreadRun(t, t.Run, waitingChain, t.Callback)
			<-waitingChain
		} else {
			go asyncThreadRun(t, t.Run, t.Callback)
		}
	}, func(e interface{}) {
		log.Fatalln(e)
	})
	return nil
}

//
// 同步线程,含callback
//
func syncThreadRun(t *Thread, f func(t *Thread), waitingChain chan int, callback func(t *Thread)) {
	t.id = ThreadsCollector.getNextThreadNum()
	ThreadsCollector.addThread(t)
	f(t)
	if callback != nil {
		callback(t)
	}
	close(waitingChain)
	ThreadsCollector.removeThread(t.id)
}

//
// 异步线程,含callback
//
func asyncThreadRun(t *Thread, f func(t *Thread), callback func(t *Thread)) {
	t.id = ThreadsCollector.getNextThreadNum()
	ThreadsCollector.addThread(t)
	f(t)
	if callback != nil {
		callback(t)
	}
	ThreadsCollector.removeThread(t.id)
}

func run() {
	fmt.Println(time.Now())
}
