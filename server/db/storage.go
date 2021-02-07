package db

import (
	"backend-grpc-challenge/server/db/entity"
	"time"
)

//NewMockDB return db for testing, in this case we use inMem storage
func NewMockDB() Storage {
	return NewInMemStorage()
}

//Storage represents a db
//*all dbs must implement this interface
type Storage interface {
	Init() error
	CreateUser(user entity.User) error
	FindUser(name string) (entity.User, error)
	FindAll() ([]entity.User, error)

	InsertMsg(msg entity.Message) error
	RetrieveMsgByTime(start, end time.Time) ([]entity.Message, error)
}
