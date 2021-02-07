package db

import (
	"backend-grpc-challenge/server/db/entity"
	"sort"
	"sync"
	"time"
)

//NewInMemStorage simple db in memory supporting atomic operation (lock)
func NewInMemStorage() Storage {
	return &inMemStorage{
		users:    make(map[int64]entity.User),
		messages: make(map[int64]entity.Message),
	}
}

type inMemStorage struct {
	lock     sync.Mutex
	users    map[int64]entity.User
	messages map[int64]entity.Message
}

//Init insert preset data in database
func (m *inMemStorage) Init() (err error) {
	err = m.CreateUser(entity.User{
		ID:       "1",
		Name:     "admin",
		Password: "lgviX0dEsDNo5bi0aF93Dhey7LGPZCcHZwWODL8saqo=", //password
	})
	if err != nil {
		return
	}
	err = m.CreateUser(entity.User{
		ID:       "2",
		Name:     "test",
		Password: "p0LqcbYx6arPpbJpSDH5eiqwY6XVkZVUiJJB+HKMxEU=", //testpassword
	})
	return
}

func (m *inMemStorage) CreateUser(user entity.User) error {
	m.lock.Lock()
	defer m.lock.Unlock()
	id := len(m.users) + 1
	m.users[int64(id)] = user

	return nil
}

func (m *inMemStorage) FindAll() ([]entity.User, error) {
	m.lock.Lock()
	defer m.lock.Unlock()
	var users []entity.User
	for _, user := range m.users {
		users = append(users, user)
	}
	return users, nil

}

func (m *inMemStorage) FindUser(name string) (entity.User, error) {
	m.lock.Lock()
	defer m.lock.Unlock()
	for _, user := range m.users {
		if user.Name == name {
			return user, nil
		}
	}
	return entity.User{}, nil

}

func (m *inMemStorage) InsertMsg(msg entity.Message) error {
	m.lock.Lock()
	defer m.lock.Unlock()
	id := len(m.messages) + 1
	m.messages[int64(id)] = msg
	return nil
}

//RetrieveMsg filter msgs from "start" to "end"
func (m *inMemStorage) RetrieveMsgByTime(start, end time.Time) ([]entity.Message, error) {
	m.lock.Lock()
	defer m.lock.Unlock()
	var msgs []entity.Message
	for _, msg := range m.messages {
		tmp := *msg.Time
		msgTime := tmp.AsTime()

		if (msgTime.After(start) || msgTime.Equal(start)) && (msgTime.Before(end) || msgTime.Equal(end)) {
			msgs = append(msgs, msg)
		}
	}

	sort.Sort(entity.ByTime(msgs))

	return msgs, nil

}
