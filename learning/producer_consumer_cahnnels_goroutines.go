package learning

import (
	"strconv"
	"sync"
	"time"

	"github.com/bubblepays/log"
)

func ProducerConsumerChGoRoutines() {
	ch := make(chan string)
	defer close(ch)
	var wg sync.WaitGroup
	//ch := make(chan int)
	ProducerCode(ch, &wg)
	ConsumerCode(ch, &wg)
	wg.Wait()
	//defer close(ch)
	//time.Sleep(time.Second * 60)
}
func ProducerCode(ch chan string, wg *sync.WaitGroup) {

	log.Info("producer function is calling")
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(ch chan string, i int) {
			time.Sleep(time.Second)
			chString := "\n " + strconv.Itoa(i) + " Alekhya \n"
			//log.Infof("producer************ : sendimg %v ",chString)
			ch <- chString

		}(ch, i)

	}
}
func ConsumerCode(ch chan string, wg *sync.WaitGroup) {
	log.Info("consumer function is calling")
	chLen := len(ch)
	log.Infof("#############length of channel %v", chLen)
	index := make(map[int]int)
	index[0] = 0
	for i := range ch {
		go func(i string, index *map[int]int, ch chan string, wg *sync.WaitGroup) {
			defer wg.Done()
			time.Sleep(time.Second)
			(*index)[0] = (*index)[0] + 1
			log.Infof("consumer @@@@@@@@@@@@@@ : receiving %v , index: %d ", i, (*index))

			if (*index)[0] == 10 {
				log.Infof("consumer channel closed ############ : receiving %v , index: %d ", i, (*index))
				//close(ch)
			}
		}(i, &index, ch, wg)

	}
	//defer close(ch)
	//defer wg.Done()
}

func GoChannels() {
	//var ch chan int
	log.Info("this is channels example program...")
	ch := make(chan int)
	ch <- 10
	v := <-ch
	log.Infof("channel value %v", v)
}
