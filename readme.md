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
Consumption | purchases-bkend-1, topic purchases, message: map[id:0 message:Hello World time:2025-09-29 22:18:10.20742 -0500 CDT m=+0.000574334] 
Consumption | purchases-bkend-1, topic purchases, message: map[id:1 message:Hello World time:2025-09-29 22:18:11.208776 -0500 CDT m=+1.001936334] 
Consumption | subscription-bkend-1, topic subscription, message: map[id:0 message:Hello World time:2025-09-29 22:18:10.20742 -0500 CDT m=+0.000574334] 
Consumption | subscription-bkend-1, topic subscription, message: map[id:1 message:Hello World time:2025-09-29 22:18:11.208776 -0500 CDT m=+1.001936334] 
Consumption | purchases-bkend-1, topic purchases, message: map[id:2 message:Hello World time:2025-09-29 22:18:12.210042 -0500 CDT m=+2.003208167] 
Consumption | purchases-bkend-1, topic purchases, message: map[id:3 message:Hello World time:2025-09-29 22:18:13.211259 -0500 CDT m=+3.004431001] 
Consumption | subscription-bkend-1, topic subscription, message: map[id:2 message:Hello World time:2025-09-29 22:18:12.210042 -0500 CDT m=+2.003208167] 
Consumption | purchases-bkend-1, topic purchases, message: map[id:4 message:Hello World time:2025-09-29 22:18:14.212481 -0500 CDT m=+4.005658084] 
Consumption | subscription-bkend-1, topic subscription, message: map[id:3 message:Hello World time:2025-09-29 22:18:13.211259 -0500 CDT m=+3.004431001] 
Consumption | subscription-bkend-1, topic subscription, message: map[id:4 message:Hello World time:2025-09-29 22:18:14.212481 -0500 CDT m=+4.005658084] 
Consumption | purchases-bkend-1, topic purchases, message: map[id:5 message:Hello World time:2025-09-29 22:18:15.212848 -0500 CDT m=+5.006031209] 
Consumption | purchases-bkend-1, topic purchases, message: map[id:6 message:Hello World time:2025-09-29 22:18:16.214048 -0500 CDT m=+6.007236001] 
Consumption | subscription-bkend-1, topic subscription, message: map[id:5 message:Hello World time:2025-09-29 22:18:15.212848 -0500 CDT m=+5.006031209] 
Consumption | purchases-bkend-1, topic purchases, message: map[id:7 message:Hello World time:2025-09-29 22:18:17.215252 -0500 CDT m=+7.008445751] 
Consumption | subscription-bkend-1, topic subscription, message: map[id:6 message:Hello World time:2025-09-29 22:18:16.214048 -0500 CDT m=+6.007236001] 
Consumption | subscription-bkend-1, topic subscription, message: map[id:7 message:Hello World time:2025-09-29 22:18:17.215252 -0500 CDT m=+7.008445751] 
Consumption | purchases-bkend-1, topic purchases, message: map[id:8 message:Hello World time:2025-09-29 22:18:18.216385 -0500 CDT m=+8.009584501] 
Consumption | subscription-bkend-1, topic subscription, message: map[id:8 message:Hello World time:2025-09-29 22:18:18.216385 -0500 CDT m=+8.009584501] 
Consumption | purchases-bkend-1, topic purchases, message: map[id:9 message:Hello World time:2025-09-29 22:18:19.217612 -0500 CDT m=+9.010817042] 
Consumption | subscription-bkend-1, topic subscription, message: map[id:9 message:Hello World time:2025-09-29 22:18:19.217612 -0500 CDT m=+9.010817042] 
Unsubscribe the topics
Shutting down the message queue...
```
