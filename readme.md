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
Subscribing subscription-bkend-1 from subscription 
Subscribing purchases-bkend-1 from purchases 
Consumption | purchases-bkend-1, topic purchases, message: map[id:1 message:Hello World time:2025-09-29 23:48:26.915037 -0500 CDT m=+0.000171001] 
Consumption | purchases-bkend-1, topic purchases, message: map[id:2 message:Hello World time:2025-09-29 23:48:27.416305 -0500 CDT m=+0.501434917] 
Consumption | subscription-bkend-1, topic subscription, message: map[id:1 message:Hello World time:2025-09-29 23:48:26.915037 -0500 CDT m=+0.000171001] 
Consumption | subscription-bkend-1, topic subscription, message: map[id:2 message:Hello World time:2025-09-29 23:48:27.416305 -0500 CDT m=+0.501434917] 
Consumption | purchases-bkend-1, topic purchases, message: map[id:3 message:Hello World time:2025-09-29 23:48:27.917887 -0500 CDT m=+1.003013709] 
Consumption | purchases-bkend-1, topic purchases, message: map[id:4 message:Hello World time:2025-09-29 23:48:28.419073 -0500 CDT m=+1.504195459] 
Consumption | subscription-bkend-1, topic subscription, message: map[id:3 message:Hello World time:2025-09-29 23:48:27.917887 -0500 CDT m=+1.003013709] 
Consumption | subscription-bkend-1, topic subscription, message: map[id:4 message:Hello World time:2025-09-29 23:48:28.419073 -0500 CDT m=+1.504195459] 
Consumption | purchases-bkend-1, topic purchases, message: map[id:5 message:Hello World time:2025-09-29 23:48:28.920419 -0500 CDT m=+2.005538251] 
Consumption | subscription-bkend-1, topic subscription, message: map[id:5 message:Hello World time:2025-09-29 23:48:28.920419 -0500 CDT m=+2.005538251] 
Consumption | purchases-bkend-1, topic purchases, message: map[id:6 message:Hello World time:2025-09-29 23:48:29.421623 -0500 CDT m=+2.506738626] 
Consumption | purchases-bkend-1, topic purchases, message: map[id:7 message:Hello World time:2025-09-29 23:48:29.92282 -0500 CDT m=+3.007932042] 
Consumption | subscription-bkend-1, topic subscription, message: map[id:6 message:Hello World time:2025-09-29 23:48:29.421623 -0500 CDT m=+2.506738626] 
Consumption | subscription-bkend-1, topic subscription, message: map[id:7 message:Hello World time:2025-09-29 23:48:29.92282 -0500 CDT m=+3.007932042] 
Consumption | purchases-bkend-1, topic purchases, message: map[id:8 message:Hello World time:2025-09-29 23:48:30.42401 -0500 CDT m=+3.509117417] 
Consumption | subscription-bkend-1, topic subscription, message: map[id:8 message:Hello World time:2025-09-29 23:48:30.42401 -0500 CDT m=+3.509117417] 
Consumption | purchases-bkend-1, topic purchases, message: map[id:9 message:Hello World time:2025-09-29 23:48:30.925198 -0500 CDT m=+4.010302417] 
Consumption | purchases-bkend-1, topic purchases, message: map[id:10 message:Hello World time:2025-09-29 23:48:31.426416 -0500 CDT m=+4.511516584] 
Consumption | subscription-bkend-1, topic subscription, message: map[id:9 message:Hello World time:2025-09-29 23:48:30.925198 -0500 CDT m=+4.010302417] 
Consumption | subscription-bkend-1, topic subscription, message: map[id:10 message:Hello World time:2025-09-29 23:48:31.426416 -0500 CDT m=+4.511516584] 
Unsubscribing subscription-bkend-1 from subscription 
Unsubscribing purchases-bkend-1 from purchases 
Shutting down the message queue...
```
