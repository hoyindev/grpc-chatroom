# How to start & examples

* functions have been implemented:
1) Specify address of server being connected to.
2) Receive messages sent by other connected clients.
3) Send messages to the server.
4) Save chat messages into a database
5) Rejoining a chat will reload the last hour's worth of messages
6) Login by input user name and password (Authentication system)

* partially done
1) Write Unit Tests 

* have not been implmented yet
1) Write Integration Tests

## Prerequisite
build server and clint
```bash
make build
```

## Log info in server
print info log in server
```bash
export GRPC_GO_LOG_SEVERITY_LEVEL=info
```

## Client

Join a chatroom using the command (default address: 127.0.0.1:8080)

```bash
bin/client --address 127.0.0.1:8080 join 
```

DB now have users: 
1. name:"admin" pw:"password" 
2. name:"test" pw:"testpassword"

You can also specific the user (default: Anonymous)
```bash
bin/client --address 127.0.0.1:8080 join --name admin --password password
```


## Server

Start a chatroom server using the command

```bash
bin/server -port 8080
```

1) Specify port for accepting incoming messages.
2) Accept messages from clients and send them to other connected clients.





## Example

client1 : anonymous user
```
$ bin/client --address 127.0.0.1:8080 join
msg1
2021-02-07 03:33:37.8917699 +0000 UTC (Anonymous) : msg1
msg2
2021-02-07 03:33:40.9342148 +0000 UTC (Anonymous) : msg2
```

client2 : admin
```
$ ./client join --name admin --password password
2021-02-07 03:33:37.8917699 +0000 UTC (Anonymous) : msg1
2021-02-07 03:33:40.9342148 +0000 UTC (Anonymous) : msg2
admin msg1
2021-02-07 03:33:59.6100702 +0000 UTC (admin) : admin msg1
admin msg2
2021-02-07 03:34:02.5832109 +0000 UTC (admin) : admin msg2
```

clien3 : test
```
$ ./client join --name test --password testpassword
2021-02-07 03:33:37.8917699 +0000 UTC (Anonymous) : msg1
2021-02-07 03:33:40.9342148 +0000 UTC (Anonymous) : msg2
2021-02-07 03:33:59.6100702 +0000 UTC (admin) : admin msg1
2021-02-07 03:34:02.5832109 +0000 UTC (admin) : admin msg2
test user 1 msg1
2021-02-07 03:34:10.6042247 +0000 UTC (test) : test user 1 msg1
test user 1 msg2
2021-02-07 03:34:13.5351272 +0000 UTC (test) : test user 1 msg2
```

server: set GRPC_GO_LOG_SEVERITY_LEVEL=info
```
$ bin/server -port 8080
INFO: 2021/02/07 11:33:23 listening port: 8080
INFO: 2021/02/07 11:33:25 conn id: 54bc47a5b784af2dfa92f291b24ac48896c71cbf15e55984efd825ff2ebb39e1 connected
INFO: 2021/02/07 11:33:37 2021-02-07 03:33:37.8917699 +0000 UTC (Anonymous) : msg1
INFO: 2021/02/07 11:33:40 2021-02-07 03:33:40.9342148 +0000 UTC (Anonymous) : msg2
INFO: 2021/02/07 11:33:47 conn id: c5c5a327fe96f19a555ff8e9783846efc7eed719bf89d9d7c28b83b0050e0113 connected
INFO: 2021/02/07 11:33:59 2021-02-07 03:33:59.6100702 +0000 UTC (admin) : admin msg1
INFO: 2021/02/07 11:34:02 2021-02-07 03:34:02.5832109 +0000 UTC (admin) : admin msg2
INFO: 2021/02/07 11:34:05 conn id: a7eda006cade9a912e8af038d620bf91f467f233ff2e27a2930e2ccd89f36407 connected
INFO: 2021/02/07 11:34:10 2021-02-07 03:34:10.6042247 +0000 UTC (test) : test user 1 msg1
INFO: 2021/02/07 11:34:13 2021-02-07 03:34:13.5351272 +0000 UTC (test) : test user 1 msg2
```