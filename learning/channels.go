package learning

import (
	"fmt"
	"net/http"
	"time"
)

func ChanneleLinks() {
	links := []string{
		"http://google.com",
		"http://facebook.com",
		"http://golang.org",
		"http://stackoverflow.com",
		"http://amazon.com",
	}
	for _, link := range links {
		CheckLink(link)
	}
}
func CheckLink(link string) {
	_, err := http.Get(link)
	if err != nil {
		fmt.Println(link, " might be down")
	}
	fmt.Println(link, " is up")
}

//ChanneleLinksGo with go routines
func ChanneleLinksGo() {
	links := []string{
		"http://google.com",
		"http://facebook.com",
		"http://golang.org",
		"http://stackoverflow.com",
		"http://amazon.com",
	}
	c := make(chan string)  //buffered channel
	for _, link := range links {
		go CheckLinkGo(link, c)
		//fmt.Println(<-c, i)
	}
	for l := range c {
		go CheckLinkGo(l, c)
	}
}
func CheckLinkGo(link string, c chan string) {
	time.Sleep(5 * time.Second)
	_, err := http.Get(link)
	if err != nil {
		fmt.Println(link, " might be down")
		c <- "mighe be down i think"
		return
	}
	fmt.Println(link, " is up")
	c <- "yeh its up"
}

func ChanneleLinksGor() {
	links := []string{
		"http://google.com",
		"http://facebook.com",
		"http://golang.org",
		"http://stackoverflow.com",
		"http://amazon.com",
	}
	c := make(chan string)
	for _, link := range links {
		go CheckLinkGor(link, c)
		//fmt.Println(<-c, i)
	}
	for i := 0; i < len(links); i++ {
		fmt.Println(<-c)
	}
}
func CheckLinkGor(link string, c chan string) {
	time.Sleep(5 * time.Second)
	_, err := http.Get(link)
	if err != nil {
		fmt.Println(link, " might be down")
		c <- "mighe be down i think"
		return
	}
	fmt.Println(link, " is up")
	c <- "yeh its up"
}

//ChanneleLinksGoRepeat repeating go routines
func ChanneleLinksGoRepeat() {
	links := []string{
		"http://google.com",
		"http://facebook.com",
		"http://golang.org",
		"http://stackoverflow.com",
		"http://amazon.com",
	}
	c := make(chan string)
	for _, link := range links {
		go CheckLinkGoRepeat(link, c)
		//fmt.Println(<-c, i)
	}
	//go repeat
	// for {
	// 	go CheckLinkGoRepeat(<-c , c)
	// }

	//alternate for loop of go repeat
	// for l := range c {
	// 	go CheckLinkGoRepeat(l , c)
	// }
	//function literal
	for l := range c {
		go func(link string) {
			time.Sleep(5 * time.Second)
			CheckLinkGoRepeat(link, c)
		}(l)
	}
}
func CheckLinkGoRepeat(link string, c chan string) {
	time.Sleep(5 * time.Second)
	_, err := http.Get(link)
	if err != nil {
		fmt.Println(link, " might be down")
		c <- link
		return
	}
	fmt.Println(link, " is up")
	c <- link
}

func StrLteral() {
	//greetings := "Hi there"
	// (func (){
	// 	log.Info( "greetings ",greetings)
	// })()

	// c := make(chan string)
	// c <- []byte("Hi there!")

	c := make(chan string)
	c <- "Hi there!"
}
