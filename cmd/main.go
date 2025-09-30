package main

import (
	"context"
	"fmt"
	"time"

	"github.com/ckalagara/pub-a-msgq.git/mq"
)

const (
	TopicSub     = "subscription"
	TopicPur     = "purchases"
	CidSubBkend1 = "subscription-bkend-1"
	CidPurBkend1 = "purchases-bkend-1"
)

func main() {
	ctxB := context.Background()
	q := mq.NewQueueWithChannelImpl([]string{TopicSub, TopicPur})

	stream, err := q.Subscribe(CidSubBkend1, TopicSub)
	if err != nil {
		fmt.Printf("Failed to sunbcribe with %v \n", err)
		return
	}

	stream2, err := q.Subscribe(CidPurBkend1, TopicPur)
	if err != nil {
		fmt.Printf("Failed to subscribe with %v \n", err)
		return
	}

	ctx, canFun := context.WithDeadline(ctxB, time.Now().Add(100*time.Second))
	defer canFun()
	go Consumer(ctx, CidSubBkend1, TopicSub, stream)
	go Consumer(ctx, CidPurBkend1, TopicPur, stream2)

	for i := 0; i < 10; i++ {
		m := make(map[string]interface{})
		m["id"] = fmt.Sprintf("%d", i)
		m["message"] = "Hello World"
		m["time"] = time.Now().String()

		if err = q.Publish(TopicPur, m); err != nil {
			fmt.Printf("Failed to publish %v \n", err)
		}
		err := q.Publish(TopicSub, m)
		if err != nil {
			fmt.Printf("Failed to publish %v \n", err)
		}
	}

	time.Sleep(1 * time.Minute)
	err = q.Unsubscribe(CidSubBkend1, TopicSub)
	if err != nil {
		fmt.Printf("Failed to unsubscribe %v \n", err)
	}
	err = q.Unsubscribe(CidPurBkend1, TopicPur)
	if err != nil {
		fmt.Printf("Failed to unsubscribe %v \n", err)
	}

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
