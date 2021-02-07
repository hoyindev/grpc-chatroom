package main

import (
	"context"
	"flag"
	"fmt"
	"net"
	"sync"
	"time"

	proto "backend-grpc-challenge/chatroompb"
	"backend-grpc-challenge/internal/auth"
	"backend-grpc-challenge/server/convertor"
	"backend-grpc-challenge/server/db"
	"backend-grpc-challenge/server/db/entity"

	"google.golang.org/grpc"
	grpcLog "google.golang.org/grpc/grpclog"
)

//Connection represent each connection with client
type Connection struct {
	stream proto.Chatroom_JoinServer
	id     string
	active bool
	err    chan error
}

//Server chatroom server
type Server struct {
	Connection []*Connection
}

var database db.Storage

func main() {
	port := flag.String("port", "8080", "port number")
	flag.Parse()

	//init db
	database = db.NewInMemStorage()
	err := database.Init()
	if err != nil {
		grpcLog.Fatalf("db err: %v", err)
	}

	//start server
	server := grpc.NewServer()
	proto.RegisterChatroomServer(server, &Server{})

	lis, err := net.Listen("tcp", ":"+*port)
	if err != nil {
		grpcLog.Fatalf("net.Listen err: %v", err)
	}
	grpcLog.Infof("listening port: %s\n", *port)

	server.Serve(lis)
}

//Login validate user info to login
func (s *Server) Login(ctx context.Context, user *proto.User) (close *proto.Close, err error) {

	if user.Name != "Anonymous" {
		//validate user
		dbRt, dbErr := database.FindUser(user.Name)
		if dbErr != nil {
			err = fmt.Errorf("db err : %s", user.Name)
			return
		}
		if dbRt.Name == "" {
			err = fmt.Errorf("no such user : %s", user.Name)
			return
		}

		if dbRt.Password != auth.MakeServerHashedPW(user.Password) {
			err = fmt.Errorf("incorrect password for: %s", user.Name)
			return
		}
	}
	return &proto.Close{}, err
}

//Join join the chatroom
func (s *Server) Join(pconn *proto.Connect, stream proto.Chatroom_JoinServer) error {

	conn := &Connection{
		stream: stream,
		id:     pconn.User.Id,
		active: true,
		err:    make(chan error),
	}
	s.Connection = append(s.Connection, conn)

	//get old msgs 60mins before (in order by time)
	end := time.Now()
	count := 60
	start := end.Add(time.Duration(-count) * time.Minute)
	records, err := database.RetrieveMsgByTime(start, end)
	if err != nil {
		grpcLog.Errorf("DB Error %v", conn.stream, err)
		conn.active = false
		conn.err <- err
	}

	//send each record
	for _, msgRecord := range records {
		post := convertor.ConvertMsgToPost(msgRecord)
		err := conn.stream.Send(&post)
		if err != nil {
			grpcLog.Errorf("Error with stream %v. Error: %v", conn.stream, err)
			conn.active = false
			conn.err <- err
		}
	}

	grpcLog.Info("conn id: ", conn.id, " connected")
	return <-conn.err
}

//BroadcastMessage broadcast
func (s *Server) BroadcastMessage(ctx context.Context, msg *proto.Post) (*proto.Close, error) {
	wait := sync.WaitGroup{}
	done := make(chan int)

	//map post to message and save to db
	message := entity.Message{
		ID:       msg.GetId(),
		UserName: msg.GetUserName(),
		Message:  msg.GetData(),
		Time:     msg.GetPostTime(),
	}
	database.InsertMsg(message)

	//log the msg shown in client
	grpcLog.Infof("%v (%s) : %s\n", msg.PostTime.AsTime(), msg.UserName, msg.Data)

	//send msg to each cliet
	for _, conn := range s.Connection {
		wait.Add(1)

		go func(msg *proto.Post, conn *Connection) {

			defer wait.Done()

			if conn.active {
				err := conn.stream.Send(msg)
				if err != nil {
					grpcLog.Errorf("Error with stream %v. Error: %v", conn.stream, err)
					conn.active = false
					conn.err <- err
				}
				// grpcLog.Infof("Sending message %v at %v", msg.Id, conn.id)

			}
		}(msg, conn)
	}

	go func() {
		wait.Wait()
		close(done)
	}()

	<-done
	return &proto.Close{}, nil
}
