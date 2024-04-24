package learning

import (
	"fmt"
	"time"
)
func ProducerConsumerChGoRoutines(){
	ch := make(chan int)
	go ProducerCode(ch)
	go ConsumerCode(ch)
	time.Sleep(time.Second*10)
}
func ProducerCode(ch chan int){
	for i:=0; i<10; i++ {
		time.Sleep(time.Second)
		fmt.Println("producer : sendimg ",i)
		ch <- i
	}
	close(ch)	
}
func ConsumerCode(ch <- chan int){
	for i := range ch {
		time.Sleep(time.Second)
		fmt.Println("consumer : receiving ",i)
	}
}