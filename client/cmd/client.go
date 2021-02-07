/*
Copyright © 2021 NAME HERE <EMAIL ADDRESS>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package cmd

import (
	"bufio"
	"context"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"log"
	"os"
	"sync"
	"time"

	"github.com/golang/protobuf/ptypes"
	"github.com/spf13/cobra"
	"google.golang.org/grpc"

	proto "backend-grpc-challenge/chatroompb"
	"backend-grpc-challenge/internal/auth"
)

var (
	clientCmdIPAddress, clientCmdName, clientCmdPW string
)

// clientCmd represents the client command
var clientCmd = &cobra.Command{
	Use:   "join",
	Short: "join the chatroom",
	Long:  `It is a client side who want to join the chatroom server`,
	Run: func(cmd *cobra.Command, args []string) {
		user := createProtoUser(clientCmdName, clientCmdPW)

		conn, err := grpc.Dial(clientCmdIPAddress, grpc.WithInsecure())
		if err != nil {
			log.Fatalf("Could not dial to server %v", err)
		}

		client := proto.NewChatroomClient(conn)

		//login
		err = login(user, client)
		if err != nil {
			log.Fatalf("Could not login to server %v", err)
		}

		//connect to chatroom
		err = connect(user, client)
		if err != nil {
			log.Fatalf("Could not connect to server %v", err)
		}

	},
}

func init() {
	rootCmd.AddCommand(clientCmd)
	clientCmd.PersistentFlags().StringVarP(&clientCmdIPAddress, "address", "a", "127.0.0.1:8080", "Server host address")
	clientCmd.PersistentFlags().StringVarP(&clientCmdName, "name", "n", "Anonymous", "user name")
	clientCmd.PersistentFlags().StringVarP(&clientCmdPW, "password", "p", "", "user password")

}

func login(user *proto.User, client proto.ChatroomClient) error {
	_, err := client.Login(context.Background(), user)
	return err
}

func connect(user *proto.User, client proto.ChatroomClient) error {
	wait := sync.WaitGroup{}
	done := make(chan int)

	var streamError error

	stream, err := client.Join(context.Background(), &proto.Connect{
		User:   user,
		Active: true,
	})

	if err != nil {
		return fmt.Errorf("Connect failed: %v", err)
	}

	wait.Add(1)

	go func(str proto.Chatroom_JoinClient) {
		defer wait.Done()
		for {
			msg, err := str.Recv()

			if err != nil {
				streamError = fmt.Errorf("Error reading message: %v", err)
				break
			}

			fmt.Printf("%v (%s) : %s\n", msg.PostTime.AsTime(), msg.UserName, msg.Data)
		}
	}(stream)

	wait.Add(1)
	go func() {
		defer wait.Done()
		scanner := bufio.NewScanner(os.Stdin)
		ts := time.Now()
		msgID := sha256.Sum256([]byte(ts.String() + user.Name))
		for scanner.Scan() {
			msg := &proto.Post{
				Id:       hex.EncodeToString(msgID[:]),
				UserName: user.GetName(),
				Data:     scanner.Text(),
				PostTime: ptypes.TimestampNow(),
			}

			_, err := client.BroadcastMessage(context.Background(), msg)
			if err != nil {
				fmt.Printf("Error sending message: %v", err)
				break
			}
		}
	}()

	go func() {
		wait.Wait()
		close(done)
	}()

	<-done
	return streamError

}

func createProtoUser(name, password string) *proto.User {
	ts := time.Now()
	id := sha256.Sum256([]byte(ts.String() + name))
	hashedPW := auth.MakeClientHashedPW(password)

	return &proto.User{
		Id:       hex.EncodeToString(id[:]),
		Name:     name,
		Password: hashedPW,
	}
}
