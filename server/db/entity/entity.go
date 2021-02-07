package entity

import "google.golang.org/protobuf/types/known/timestamppb"

type User struct {
	ID       string
	Name     string
	Password string
}

//Message client messages
type Message struct {
	Time     *timestamppb.Timestamp
	UserName string
	Message  string
	ID       string
}

// ByTime implements sort.Interface based on the Time field.
type ByTime []Message

func (bt ByTime) Len() int { return len(bt) }
func (bt ByTime) Less(i, j int) bool {
	return bt[i].Time.AsTime().UnixNano() < bt[j].Time.AsTime().UnixNano()
}
func (bt ByTime) Swap(i, j int) { bt[i], bt[j] = bt[j], bt[i] }
