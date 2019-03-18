package thread

import (
	"fmt"
	"reflect"
	"runtime"
	"strconv"
	"testing"
	"time"
)

//
// author: kango2gler@gmail.com
// date: 2018-03-20
//

func TestRun(t *testing.T) {
	fmt.Println("TestRun start:", 1)
	//run()
	//newThreadTest(t)
	new1000ThreadTest(t)
}

//创建线程测试
// t1 异步 2s 延时
// t2 异步 1s 延时
// t3 异步 无 延时
// t4 同步 无 延时
// fmt.Println("it main thread") 正常主线程中输出
// 结果:
//TestRun start: 1
//t4: 1 58 []	//同步 无 延时
//it main thread //主线程执行,在t4后执行
//t3: 2		//异步 无 延时
//t2: 3		//1s 延时
//t1: 4		//2s 延时最后执行
//end
//
func newThreadTest(t *testing.T) {
	var i = 0
	t1 := &Thread{Name: "t1", Async: true, Run: func(t *Thread) {
		time.Sleep(2 * time.Second)
		i++
		fmt.Println("t1:", i)
	}}

	t2 := &Thread{Name: "t2", Async: true, Run: func(t *Thread) {
		time.Sleep(1 * time.Second)
		i++
		fmt.Println("t2:", i)
	}}

	t3 := &Thread{Name: "t3", Async: true, Run: func(t *Thread) {
		i++
		fmt.Println("t3:", i)
	}}

	t4 := &Thread{Name: "t4", Async: false, Run: func(t *Thread) {
		i++
		fmt.Println("t4:", i, ':', reflect.ValueOf(t))
	}}
	t1.Start()
	t2.Start()
	t3.Start()
	t4.Start()
	fmt.Println("it main thread")
	fmt.Println(Threads())

	(&Thread{Name: "t4", Async: false, Run: func(t *Thread) {
		for i := 0; i < 100; i++ {
			fmt.Println("t5:", i)
		}
	}}).Start()

	time.Sleep(3 * time.Second)
	fmt.Println("end")
	fmt.Println(Threads())
}

//创建1000个线程测试
func new1000ThreadTest(t1 *testing.T) {
	k := make([]int, 5)
	for i := 0; i < 10; i++ {
		(&Thread{Name: ("t" + strconv.Itoa(i)), Async: false, Run: func(t *Thread) {
			//for j:=0;j<100;j++{
			//	fmt.Println("t",i,j,":",i)
			//}
			for len(k) < 5 {
				k[len(k)] = 1
				fmt.Println("k:", k)
			}
		}}).Start()
		fmt.Println("t", i)
	}
	go func() {
		k[1] = 1
	}()
	fmt.Println("线程数:", runtime.NumGoroutine())
	//等待所有线程结束
	ThreadsCollector.WaitThreadsEnd()
	fmt.Println("end")

	// for {
	// 	//time.Sleep(1000*time.Hour)
	// }

}
