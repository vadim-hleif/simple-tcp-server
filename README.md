## simple-tcp-server

![Test](https://github.com/vadim-hleif/simple-tcp-server/workflows/Test/badge.svg)

### What is it
TCP service in Go. 

Allow connecting by TCP. It notifies user's friends about his status changing (online or not).

How to connect by TCP:
```shell
telnet localhost 8080
```
Send a json payload:
```json
{"user_id": 1, "friends": [2, 3, 4]}
```
It will be stored in-memory.

When another connection is established with the `user_id` from the list of any other user's friends section, they will be notified about it with message: 
```json
{"friend_id": 1, "online": true}
```
When the user goes offline, their friends (if they have some and any of them are online) will receive a message:
```json
{"friend_id": 1, "online": false}
```

### Build:

```shell
docker build -t simple-tcp-server . 
```

### Run:

```shell
 docker run -it -p=8080:8080 simple-tcp-server
```