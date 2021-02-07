package cmd

import (
	"context"
	"fmt"
	"log"
	"net"
	"strings"
	"testing"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/test/bufconn"

	proto "backend-grpc-challenge/chatroompb"
)

var mockDB map[string]proto.User

func init() {
	mockDB = make(map[string]proto.User)
	mockDB["admin"] = proto.User{}
}

func TestLogin(t *testing.T) {

	ctx := context.TODO()
	srv, listener := startGRPCServer()

	defer func() { time.Sleep(10 * time.Millisecond) }()
	defer srv.Stop()
	conn, err := grpc.DialContext(ctx, "", grpc.WithContextDialer(getBufDialer(listener)), grpc.WithInsecure())
	if err != nil {
		t.Fatalf("failed to dial: %v", err)
	}
	defer conn.Close()
	client := proto.NewChatroomClient(conn)

	tests := []struct {
		a    proto.User
		want string
	}{
		{
			*createProtoUser("admin", ""),
			"",
		},
		{
			*createProtoUser("noSuchUser", ""),
			"no such user",
		},
	}

	for _, tt := range tests {
		testname := fmt.Sprintf("%v", &tt.a)
		t.Run(testname, func(t *testing.T) {

			_, rt := client.Login(context.Background(), &tt.a)
			equal := ErrorContains(rt, tt.want)
			if equal == false {
				t.Errorf("Error actual = %v, and Expected = %v.", &tt.a, tt.want)
			}

		})
	}
}

func TestBroadcastMessage(t *testing.T) {

	ctx := context.TODO()
	srv, listener := startGRPCServer()

	defer func() { time.Sleep(10 * time.Millisecond) }()
	defer srv.Stop()
	conn, err := grpc.DialContext(ctx, "", grpc.WithContextDialer(getBufDialer(listener)), grpc.WithInsecure())
	if err != nil {
		t.Fatalf("failed to dial: %v", err)
	}
	defer conn.Close()
	client := proto.NewChatroomClient(conn)

	msg := proto.Post{}
	_, err = client.BroadcastMessage(ctx, &msg)
	if err != nil {
		t.Errorf("Error BoradcastMessage func %v", err)
	}
}

func TestJoin(t *testing.T) {

	ctx := context.TODO()
	srv, listener := startGRPCServer()

	defer func() { time.Sleep(10 * time.Millisecond) }()
	defer srv.Stop()
	conn, err := grpc.DialContext(ctx, "", grpc.WithContextDialer(getBufDialer(listener)), grpc.WithInsecure())
	if err != nil {
		t.Fatalf("failed to dial: %v", err)
	}
	defer conn.Close()
	client := proto.NewChatroomClient(conn)

	in := proto.Connect{}
	_, err = client.Join(ctx, &in)
	if err != nil {
		t.Errorf("Error Join func %v", err)
	}
}

// func startGRPCServer(impl *vacancies.Implementation) (*grpc.Server, *bufconn.Listener) {
func startGRPCServer() (*grpc.Server, *bufconn.Listener) {
	bufferSize := 1024 * 1024
	listener := bufconn.Listen(bufferSize)
	server := grpc.NewServer()

	proto.RegisterChatroomServer(server, &mockServer{})
	go func() {
		if err := server.Serve(listener); err != nil {
			log.Fatalf("failed to start grpc server: %v", err)
		}
	}()
	return server, listener
}

func getBufDialer(listener *bufconn.Listener) func(context.Context, string) (net.Conn, error) {
	return func(ctx context.Context, url string) (net.Conn, error) {
		return listener.Dial()
	}
}

//mockServer mock chatroom server
type mockServer struct {
	// Connection []*Connection
}

//Join join the chatroom
func (s *mockServer) Join(pconn *proto.Connect, stream proto.Chatroom_JoinServer) error {
	return nil
}

//BroadcastMessage broadcast
func (s *mockServer) BroadcastMessage(ctx context.Context, msg *proto.Post) (*proto.Close, error) {
	return &proto.Close{}, nil
}

//Login validate user info to login
func (s *mockServer) Login(ctx context.Context, user *proto.User) (close *proto.Close, err error) {
	if _, ok := mockDB[user.Name]; !ok {
		err = fmt.Errorf("no such user")
	}
	return &proto.Close{}, err
}

// ErrorContains checks if the error message in out contains the text in
// want.
//
// This is safe when out is nil. Use an empty string for want if you want to
// test that err is nil.
func ErrorContains(out error, want string) bool {
	if out == nil {
		return want == ""
	}
	if want == "" {
		return false
	}
	return strings.Contains(out.Error(), want)
}
