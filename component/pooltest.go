package component

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

type Job struct {
	Id, Randomno int
}

type Result struct {
	job         Job
	sumofdigits int
}

type CommonInterface interface {
	PrintSome(var1 string) string
}

func (job Job) PrintSome(var1 string) string {
	return "something"
}

var jobs = make(chan Job, 10)
var results = make(chan Result, 10)

func digits(num int) int {
	sum := 0
	no := num
	for no != 0 {
		digit := no % 10
		sum += digit
		no /= 10
	}
	time.Sleep(2 * time.Second)
	return sum
}

func worker(wg *sync.WaitGroup) {
	for job := range jobs {
		output := Result{job, digits(job.Randomno)}
		results <- output
	}
	wg.Done()
}

func CreateWorkerPool(noOfWorkers int) {
	var wg sync.WaitGroup
	for i := 0; i < noOfWorkers; i++ {
		wg.Add(1)
		go worker(&wg)
	}
	wg.Wait()
	close(results)
}

func Allocate(noOfJobs int) {
	fmt.Println("job start allocate")
	for i := 0; i < noOfJobs; i++ {
		randNo := rand.Intn(999)
		job := Job{i, randNo}
		job.PrintSome("666")
		jobs <- job
	}
	close(jobs)
}

func GetResult(done chan int) {
	var1 := 0
	for result := range results {
		var1 += result.sumofdigits
		fmt.Printf("Job id %d, input random no %d , sum of digits %d\n", result.job.Id, result.job.Randomno, result.sumofdigits)
	}
	done <- var1
}

func Test() {
	fmt.Println("this is a test")
}

func server1(ch chan string) {
	time.Sleep(6 * time.Second)
	fmt.Print("server1 start work")
	ch <- "server1 start work"
}

func server2(ch chan string) {
	time.Sleep(3 * time.Second)
	fmt.Println("server2 start work")
	ch <- "server2 start work"
}

func SelectTest() {
	output1 := make(chan string)
	output2 := make(chan string)
	go server1(output1)
	go server2(output2)
	select {
	case s1 := <-output1:
		fmt.Println(s1)
	case s2 := <-output2:
		fmt.Println(s2)
	}
}

var x = 0
var flag1 = 0

func increament(wg *sync.WaitGroup) {
	x += 1
	wg.Done()
}
func NoMutex() {
	var wg sync.WaitGroup
	for i := 0; i < 100; i++ {
		wg.Add(1)
		go increament(&wg)
	}
	wg.Wait()
	fmt.Println("the result is", x)
}

func increamentWithMutex(wg *sync.WaitGroup, m *sync.Mutex) {
	m.Lock()
	x += 1
	m.Unlock()
	wg.Done()
}

func WithMutex() {
	if x != 0 && flag1 == 0 {
		x = 0
		flag1 = 1
	}
	var wg sync.WaitGroup
	var mutex sync.Mutex
	for i := 0; i < 100; i++ {
		wg.Add(1)
		go increamentWithMutex(&wg, &mutex)
	}
	wg.Wait()
	fmt.Println("the result is", x)
}

func increatWithChannel(wg *sync.WaitGroup, channel chan bool) {
	channel <- true
	x += 1
	<-channel
	wg.Done()
}

func ChannelParallel() {
	var wg sync.WaitGroup
	channel := make(chan bool, 1)
	for i := 0; i < 100; i++ {
		wg.Add(1)
		go increatWithChannel(&wg, channel)
	}
	wg.Wait()
	fmt.Println("the result is", x)
}
