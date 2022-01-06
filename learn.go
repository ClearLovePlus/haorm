package main

import (
	"fmt"
	"sync"
	"time"

	"github.com/ClearLovePlus/haorm/dialect"
)

func chanTest() {
	ch := make(chan string, 3)
	ch <- "chen"
	ch <- "hao"
	ch <- "3"
	fmt.Printf("\n容量是:%d", cap(ch))
	fmt.Printf("\n长度是:%d", len(ch))
	fmt.Println("\n读取数据", <-ch)
	fmt.Println("新的长度是", len(ch))
}

/**
    类似于java中的countDownlatch
**/
func waitGroupTest(i int, wg *sync.WaitGroup, result chan int) {
	fmt.Println("started Goroutine", i)
	time.Sleep(2 * time.Second)
	fmt.Printf("Goroutine %d ended\n", i)
	result <- i
	wg.Done()
}

type Test struct {
	Name string `haorm:"PRIMARY KEY"`
	Age  int
}

var TestDialect, _ = dialect.GetDialect("mysql")

func main() {
	// 	fmt.Printf("\nprod1 start")
	// 	chanTest()
	// 	no := 3
	// 	var1 := make(chan int)
	// 	var2 := make(chan int)
	// 	var3 := make(chan int)
	// 	var wg sync.WaitGroup
	// 	for i := 0; i < no; i++ {
	// 		wg.Add(1)
	// 		switch i {
	// 		case 0:
	// 			go waitGroupTest(i, &wg, var1)

	// 		case 1:
	// 			go waitGroupTest(i, &wg, var2)

	// 		case 2:
	// 			go waitGroupTest(i, &wg, var3)

	// 		default:
	// 			go waitGroupTest(i, &wg, var1)
	// 		}

	// 	}

	// 	var4, var5, var6 := <-var1, <-var2, <-var3
	// 	wg.Wait()
	// 	fmt.Println("三个协程同时计算完成结果为:", var4+var5+var6)
	// 	fmt.Println("其他协程都已经计算结束")
	// 	//线程池test
	// 	// startTime := time.Now()
	// 	// noOfJobs := 100
	// 	// go component.Allocate(noOfJobs)
	// 	// done := make(chan int)
	// 	// go component.GetResult(done)
	// 	// noOfWorkers := 10
	// 	// go component.CreateWorkerPool(noOfWorkers)
	// 	// fmt.Println("the result is", <-done)
	// 	// endTime := time.Now()
	// 	// diff := endTime.Sub(startTime)
	// 	// fmt.Println("total time taken ", diff.Seconds(), "seconds")

	// 	//select Test
	// 	// component.SelectTest()

	// 	fmt.Println("===========这里是没加锁的========")
	// 	//with no mutex test
	// 	for i := 0; i < 5; i++ {
	// 		component.NoMutex()
	// 	}
	// 	fmt.Println("=========加完锁的区别============")
	// 	//with mutex test
	// 	for i := 0; i < 5; i++ {
	// 		component.WithMutex()
	// 	}
	// 	fmt.Println("========信道并发锁===========")
	// 	for i := 0; i < 5; i++ {
	// 		component.ChannelParallel()
	// 	}

	// 	//go interface test
	// 	var com component.CommonInterface

	// 	job := component.Job{}
	// 	com = job
	// 	result := com.PrintSome("122222")
	// 	fmt.Println(result)

	// 	employee := component.New("chenhao3", "20190001", 1)
	// 	employee1 := component.New("chenhao3", "20190001", 2)
	// 	employee2 := component.New("chenhao3", "20190001", 23)
	// 	e := []component.Employee{employee, employee1, employee2}
	// 	fmt.Println("\n" + employee.ToString())
	// 	component.CombineTest()
	// 	component.AAndB()
	// 	f := func(a, b int) int {
	// 		return a * b
	// 	}
	// 	component.Simple(f)
	// 	f1 := component.Simple1()
	// 	fmt.Println("返回的函数计算为:", f1(6, 7))

	// 	a1 := component.AppenStr()
	// 	b1 := component.AppenStr()
	// 	fmt.Println(a1("World"))
	// 	fmt.Println(b1("everyone"))
	// 	fmt.Println(a1("Gopher"))
	// 	fmt.Println(b1("!"))
	// 	p1 := 0
	// 	p2 := 1
	// 	gen := component.MakeFibGen(&p1, &p2)
	// 	for i := 0; i < 10; i++ {
	// 		fmt.Print(component.MakeFibGen1(p2, i))
	// 		fmt.Print(" ")
	// 	}

	// 	for i := 0; i < 10; i++ {
	// 		fmt.Print(gen())
	// 		fmt.Print(" ")
	// 	}
	// 	es := component.Filter(e, func(e component.Employee) bool {
	// 		return e.Id > 1
	// 	})

	// 	fmt.Println(es)
	// 	component.CompleteCreateQuery(employee)
	// 	component.HandlerTxt()
	// 	component.BfReader()
	// 	bubble := []int{6, 4, 7, 9, 1}
	// 	component.BubbleSort(bubble)
	DbTest()
	// schema := schema.Parse(&Test{}, TestDialect)
	// if schema.Name != "Test" || len(schema.Fields) != 2 {
	// 	log.Println("parse error")
	// }
	// if schema.GetField("Name").Tag != "PRIMARY KEY" {
	// 	log.Fatal("failed to parse primary key")
	// }
}
