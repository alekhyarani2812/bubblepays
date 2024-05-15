package learning

import (
	"time"

	"github.com/bubblepays/log"
)

func ProducerConsumerChGoRoutines() {
	ch := make(chan int)
	go ProducerCode(ch)
	go ConsumerCode(ch)
	time.Sleep(time.Second * 10)
}
func ProducerCode(ch chan int) {
	for i := 0; i < 10; i++ {
		time.Sleep(time.Second)
		log.Infof("producer : sendimg %v ", i)
		ch <- i
	}
	close(ch)
}
func ConsumerCode(ch <-chan int) {
	for i := range ch {
		time.Sleep(time.Second)
		log.Infof("consumer : receiving %v ", i)
	}
}
