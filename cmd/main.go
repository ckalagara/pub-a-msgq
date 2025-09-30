package main

import (
	"context"
	"fmt"
	"sync"
	"time"

	"github.com/ckalagara/pub-a-msgq/mq"
)

const (
	TopicSub     = "subscription"
	TopicPur     = "purchases"
	CidSubBkend1 = "subscription-bkend-1"
	CidPurBkend1 = "purchases-bkend-1"

	demoMsgCount = 10

	messagePublishDelay    = 500 * time.Millisecond
	consumerSessionTimeout = 15 * time.Second
)

func main() {
	ctxB := context.Background()
	// new msgQ
	q := mq.NewQueueWithChannelImpl([]string{TopicSub, TopicPur})

	wg := sync.WaitGroup{}
	// creating consumer as routines
	fmt.Println("Creating consumers")
	ctx, canFun := context.WithDeadline(ctxB, time.Now().Add(consumerSessionTimeout))
	defer canFun()

	wg.Go(func() {
		Consumer(ctx, CidSubBkend1, TopicSub, q)
	})

	wg.Go(func() {
		Consumer(ctx, CidPurBkend1, TopicPur, q)
	})

	// Publishing few messages
	PublishMessages(q, demoMsgCount)

	wg.Wait()

	// shutting down
	q.Shutdown()

}

func PublishMessages(q mq.Queue, count int) {
	fmt.Println("Publishing messages to topics")
	for i := 1; i <= count; i++ {
		m := make(map[string]interface{})
		m["id"] = fmt.Sprintf("%d", i)
		m["message"] = "Hello World"
		m["time"] = time.Now().String()

		if err := q.Publish(TopicPur, m); err != nil {
			fmt.Printf("Failed to publish %v \n", err)
		}
		// artificial dealy
		time.Sleep(messagePublishDelay)

		if err := q.Publish(TopicSub, m); err != nil {
			fmt.Printf("Failed to publish %v \n", err)
		}
	}
}

func Consumer(ctx context.Context, clientID, topic string, q mq.Queue) {
	// subscribing
	fmt.Printf("Subscribing %s from %s \n", clientID, topic)
	ch, err := q.Subscribe(clientID, topic)
	if err != nil {
		fmt.Printf("Failed to subscribe with %v \n", err)
		return
	}

	// consuming
readerLoop:
	for {
		select {
		case <-ctx.Done():
			break readerLoop
		case msg := <-ch:
			fmt.Printf("Consumption | %s, topic %s, message: %v \n", clientID, topic, msg)
		}
	}

	// unsubscribing
	fmt.Printf("Unsubscribing %s from %s \n", clientID, topic)
	err = q.Unsubscribe(clientID, topic)
	if err != nil {
		fmt.Printf("Failed to unsubscribe %v \n", err)
	}
}
