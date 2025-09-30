# pub-a-msgQ
A simple message queue, implemented using go channels.

## run
```
cd cmd
go run main.go
```
```
Creating new queue with channel impl with topics [subscription purchases] 
Creating consumers
Publishing messages to topics
Subscribing subscription-bkend-1 client to subscription topic 
Subscribing purchases-bkend-1 client to purchases topic 
Publishing | cart-bkend-1, topic purchases, message: map[id:1 message:Hello World time:2025-09-30 00:01:05.721966 -0500 CDT m=+0.000432209] 
Consuming  | purchases-bkend-1, topic purchases, message: map[id:1 message:Hello World time:2025-09-30 00:01:05.721966 -0500 CDT m=+0.000432209] 
Publishing | cart-bkend-1, topic subscription, message: map[id:1 message:Hello World time:2025-09-30 00:01:05.721966 -0500 CDT m=+0.000432209] 
Publishing | cart-bkend-1, topic purchases, message: map[id:2 message:Hello World time:2025-09-30 00:01:06.223833 -0500 CDT m=+0.502297334] 
Consuming  | purchases-bkend-1, topic purchases, message: map[id:2 message:Hello World time:2025-09-30 00:01:06.223833 -0500 CDT m=+0.502297334] 
Consuming  | subscription-bkend-1, topic subscription, message: map[id:1 message:Hello World time:2025-09-30 00:01:05.721966 -0500 CDT m=+0.000432209] 
Publishing | cart-bkend-1, topic subscription, message: map[id:2 message:Hello World time:2025-09-30 00:01:06.223833 -0500 CDT m=+0.502297334] 
Publishing | cart-bkend-1, topic purchases, message: map[id:3 message:Hello World time:2025-09-30 00:01:06.725176 -0500 CDT m=+1.003639834] 
Consuming  | purchases-bkend-1, topic purchases, message: map[id:3 message:Hello World time:2025-09-30 00:01:06.725176 -0500 CDT m=+1.003639834] 
Consuming  | subscription-bkend-1, topic subscription, message: map[id:2 message:Hello World time:2025-09-30 00:01:06.223833 -0500 CDT m=+0.502297334] 
Publishing | cart-bkend-1, topic subscription, message: map[id:3 message:Hello World time:2025-09-30 00:01:06.725176 -0500 CDT m=+1.003639834] 
Publishing | cart-bkend-1, topic purchases, message: map[id:4 message:Hello World time:2025-09-30 00:01:07.226565 -0500 CDT m=+1.505027667] 
Consuming  | purchases-bkend-1, topic purchases, message: map[id:4 message:Hello World time:2025-09-30 00:01:07.226565 -0500 CDT m=+1.505027667] 
Consuming  | subscription-bkend-1, topic subscription, message: map[id:3 message:Hello World time:2025-09-30 00:01:06.725176 -0500 CDT m=+1.003639834] 
Publishing | cart-bkend-1, topic subscription, message: map[id:4 message:Hello World time:2025-09-30 00:01:07.226565 -0500 CDT m=+1.505027667] 
Publishing | cart-bkend-1, topic purchases, message: map[id:5 message:Hello World time:2025-09-30 00:01:07.728047 -0500 CDT m=+2.006509167] 
Consuming  | purchases-bkend-1, topic purchases, message: map[id:5 message:Hello World time:2025-09-30 00:01:07.728047 -0500 CDT m=+2.006509167] 
Consuming  | subscription-bkend-1, topic subscription, message: map[id:4 message:Hello World time:2025-09-30 00:01:07.226565 -0500 CDT m=+1.505027667] 
Publishing | cart-bkend-1, topic subscription, message: map[id:5 message:Hello World time:2025-09-30 00:01:07.728047 -0500 CDT m=+2.006509167] 
Publishing | cart-bkend-1, topic purchases, message: map[id:6 message:Hello World time:2025-09-30 00:01:08.22937 -0500 CDT m=+2.507830709] 
Consuming  | purchases-bkend-1, topic purchases, message: map[id:6 message:Hello World time:2025-09-30 00:01:08.22937 -0500 CDT m=+2.507830709] 
Consuming  | subscription-bkend-1, topic subscription, message: map[id:5 message:Hello World time:2025-09-30 00:01:07.728047 -0500 CDT m=+2.006509167] 
Publishing | cart-bkend-1, topic subscription, message: map[id:6 message:Hello World time:2025-09-30 00:01:08.22937 -0500 CDT m=+2.507830709] 
Publishing | cart-bkend-1, topic purchases, message: map[id:7 message:Hello World time:2025-09-30 00:01:08.730195 -0500 CDT m=+3.008655126] 
Consuming  | purchases-bkend-1, topic purchases, message: map[id:7 message:Hello World time:2025-09-30 00:01:08.730195 -0500 CDT m=+3.008655126] 
Consuming  | subscription-bkend-1, topic subscription, message: map[id:6 message:Hello World time:2025-09-30 00:01:08.22937 -0500 CDT m=+2.507830709] 
Publishing | cart-bkend-1, topic subscription, message: map[id:7 message:Hello World time:2025-09-30 00:01:08.730195 -0500 CDT m=+3.008655126] 
Publishing | cart-bkend-1, topic purchases, message: map[id:8 message:Hello World time:2025-09-30 00:01:09.231329 -0500 CDT m=+3.509788084] 
Consuming  | purchases-bkend-1, topic purchases, message: map[id:8 message:Hello World time:2025-09-30 00:01:09.231329 -0500 CDT m=+3.509788084] 
Consuming  | subscription-bkend-1, topic subscription, message: map[id:7 message:Hello World time:2025-09-30 00:01:08.730195 -0500 CDT m=+3.008655126] 
Publishing | cart-bkend-1, topic subscription, message: map[id:8 message:Hello World time:2025-09-30 00:01:09.231329 -0500 CDT m=+3.509788084] 
Publishing | cart-bkend-1, topic purchases, message: map[id:9 message:Hello World time:2025-09-30 00:01:09.732388 -0500 CDT m=+4.010845376] 
Consuming  | purchases-bkend-1, topic purchases, message: map[id:9 message:Hello World time:2025-09-30 00:01:09.732388 -0500 CDT m=+4.010845376] 
Consuming  | subscription-bkend-1, topic subscription, message: map[id:8 message:Hello World time:2025-09-30 00:01:09.231329 -0500 CDT m=+3.509788084] 
Publishing | cart-bkend-1, topic subscription, message: map[id:9 message:Hello World time:2025-09-30 00:01:09.732388 -0500 CDT m=+4.010845376] 
Publishing | cart-bkend-1, topic purchases, message: map[id:10 message:Hello World time:2025-09-30 00:01:10.233511 -0500 CDT m=+4.511967626] 
Consuming  | purchases-bkend-1, topic purchases, message: map[id:10 message:Hello World time:2025-09-30 00:01:10.233511 -0500 CDT m=+4.511967626] 
Consuming  | subscription-bkend-1, topic subscription, message: map[id:9 message:Hello World time:2025-09-30 00:01:09.732388 -0500 CDT m=+4.010845376] 
Publishing | cart-bkend-1, topic subscription, message: map[id:10 message:Hello World time:2025-09-30 00:01:10.233511 -0500 CDT m=+4.511967626] 
Consuming  | subscription-bkend-1, topic subscription, message: map[id:10 message:Hello World time:2025-09-30 00:01:10.233511 -0500 CDT m=+4.511967626] 
Unsubscribing purchases-bkend-1 client to purchases topic 
Unsubscribing subscription-bkend-1 client to subscription topic 
Shutting down the message queue..
```
