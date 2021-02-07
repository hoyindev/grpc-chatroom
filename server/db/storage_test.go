package db

import (
	"backend-grpc-challenge/server/db/entity"
	"fmt"
	"testing"
	"time"

	"google.golang.org/protobuf/types/known/timestamppb"
)

func TestCreateUser(t *testing.T) {
	tests := []struct {
		a    entity.User
		b    entity.User
		want int
	}{
		{
			entity.User{
				ID:       "1",
				Name:     "admin",
				Password: "password",
			},
			entity.User{
				ID:       "2",
				Name:     "testadmin",
				Password: "testpassword",
			},
			2,
		},
	}

	for _, tt := range tests {
		testname := fmt.Sprintf("%v,%v", tt.a, tt.b)
		t.Run(testname, func(t *testing.T) {
			mockDB := NewMockDB()
			mockDB.CreateUser(tt.a)
			mockDB.CreateUser(tt.b)
			ans, err := mockDB.FindAll()
			if err != nil {
				t.Errorf("FindUser err %d", err)
			}

			if len(ans) != tt.want {
				t.Errorf("got %v, want %v", len(ans), tt.want)
			}
		})
	}
}

func TestFindUser(t *testing.T) {
	tests := []struct {
		a    entity.User
		b    entity.User
		want entity.User
	}{
		{
			entity.User{
				ID:       "1",
				Name:     "admin",
				Password: "password",
			},
			entity.User{
				ID:       "2",
				Name:     "testadmin",
				Password: "testpassword",
			},
			entity.User{
				ID:       "1",
				Name:     "admin",
				Password: "password",
			},
		},
	}

	for _, tt := range tests {
		testname := fmt.Sprintf("%v,%v", tt.a, tt.b)
		t.Run(testname, func(t *testing.T) {
			mockDB := NewMockDB()
			mockDB.CreateUser(tt.a)
			mockDB.CreateUser(tt.b)
			ans, err := mockDB.FindUser("admin")
			if err != nil {
				t.Errorf("FindUser err %d", err)
			}

			if ans != tt.want {
				t.Errorf("got %v, want %v", ans, tt.want)
			}
		})
	}

}

func TestInsertMsg(t *testing.T) {
	now := timestamppb.Now()
	msg := entity.Message{
		Time:     now,
		UserName: "a",
		Message:  "hello",
	}
	mockDB := NewMockDB()
	mockDB.InsertMsg(msg)
	mockDB.InsertMsg(msg)

	end := time.Now()
	count := 60
	start := end.Add(time.Duration(-count) * time.Minute)

	rt, _ := mockDB.RetrieveMsgByTime(start, end)
	t.Error(rt)

}
