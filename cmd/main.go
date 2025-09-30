package main

import (
	"context"
	"fmt"
	"time"

	"github.com/ckalagara/pub-a-msgq/mq"
)

const (
	TopicSub     = "subscription"
	TopicPur     = "purchases"
	CidSubBkend1 = "subscription-bkend-1"
	CidPurBkend1 = "purchases-bkend-1"
)

func main() {
	ctxB := context.Background()
	// new msgQ
	q := mq.NewQueueWithChannelImpl([]string{TopicSub, TopicPur})

	// sub1
	stream, err := q.Subscribe(CidSubBkend1, TopicSub)
	if err != nil {
		fmt.Printf("Failed to sunbcribe with %v \n", err)
		return
	}

	// sub2
	stream2, err := q.Subscribe(CidPurBkend1, TopicPur)
	if err != nil {
		fmt.Printf("Failed to subscribe with %v \n", err)
		return
	}

	// creating consumer as routines
	fmt.Println("Creating consumers")
	ctx, canFun := context.WithDeadline(ctxB, time.Now().Add(20*time.Second))
	defer canFun()
	go Consumer(ctx, CidSubBkend1, TopicSub, stream)
	go Consumer(ctx, CidPurBkend1, TopicPur, stream2)

	// Publishing few messages
	fmt.Println("Publishing messages to topics")
	for i := 0; i < 10; i++ {
		m := make(map[string]interface{})
		m["id"] = fmt.Sprintf("%d", i)
		m["message"] = "Hello World"
		m["time"] = time.Now().String()

		if err = q.Publish(TopicPur, m); err != nil {
			fmt.Printf("Failed to publish %v \n", err)
		}
		// artifical dealy
		time.Sleep(1 * time.Second)

		err = q.Publish(TopicSub, m)
		if err != nil {
			fmt.Printf("Failed to publish %v \n", err)
		}
	}

	// waiting for the consumers tor read
	time.Sleep(10 * time.Second)

	// Unsubscribing
	fmt.Println("Unsubscribe the topics")
	err = q.Unsubscribe(CidSubBkend1, TopicSub)
	if err != nil {
		fmt.Printf("Failed to unsubscribe %v \n", err)
	}
	err = q.Unsubscribe(CidPurBkend1, TopicPur)
	if err != nil {
		fmt.Printf("Failed to unsubscribe %v \n", err)
	}

	// shutting down
	q.Shutdown()

}

func Consumer(ctx context.Context, clientID, topic string, ch chan mq.Message) {
	for {
		select {
		case <-ctx.Done():
			return
		case msg := <-ch:
			fmt.Printf("Consumption | %s, topic %s, message: %v \n", clientID, topic, msg)
		}
	}
}
