//////////////////////////////////////////////////////////////////////
//
// Given is a producer-consumer szenario, where a producer reads in
// tweets from a mockstream and a consumer is processing the
// data. Your task is to change the code so that the producer as well
// as the consumer can run concurrently
//

package main

import (
	"fmt"
	"sync"
	"time"
)

func producer(stream Stream) (tweets []*Tweet) {
	defer wait.Done()
	for {
		tweet, err := stream.Next()
		if err == ErrEOF {
			close(cha)
			return tweets
		}
		cha <- tweet
		tweets = append(tweets, tweet)
	}
}

func consumer() {
	defer wait.Done()
	for {
		t, ok := <-cha
		if !ok {
			return
		}
		if !t.IsTalkingAboutGo() {
			fmt.Println(t.Username, "\tdoes not tweet about golang")
		} else {
			fmt.Println(t.Username, "\ttweets about golang")
		}
	}
}

var wait sync.WaitGroup
var cha chan *Tweet

func main() {

	wait.Add(2)
	cha = make(chan *Tweet)
	start := time.Now()
	stream := GetMockStream()

	// Producer
	go producer(stream)

	// Consumer
	go consumer()
	wait.Wait()

	fmt.Printf("Process took %s\n", time.Since(start))
}
