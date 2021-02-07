# Zenport Code Challenge for Backend Role

* Build a client server application for relaying messages in a chat room.

If you disagree without protocol messaging structure feel free to change it and explain why you feel your idea is more optimal.

## Client

Join a chatroom using the command

```bash
./client --address 127.0.0.1:8080 join 
```

You should now be able to:
1) Specify address of server being connected to.
2) Receive messages sent by other connected clients.
3) Send messages to the server.

## Server

Start a chatroom server using the command

```bash
./server -port 8080
```

1) Specify port for accepting incoming messages.
2) Accept messages from clients and send them to other connected clients.

## GRPC

* There should be a single rpc endpoint that is a bidirectional stream of Posts

Post messages should be formatted as such

```proto
message Post {
    string user_name = 1;
    string data = 2;
    google.protobuf.Timestamp post_time = 3;
}
```

## Bonus
* `+1` Write Unit Tests
* `+2` Write Integration Tests
* `+2` Save chat messages into a database
* `+2` Rejoining a chat will reload the last hour's worth of messages
* `+3` Authentication system

