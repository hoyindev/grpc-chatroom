package convertor

import (
	proto "backend-grpc-challenge/chatroompb"
	"backend-grpc-challenge/server/db/entity"
)

//ConvertMsgToPost map Msg to Post
func ConvertMsgToPost(msg entity.Message) proto.Post {

	return proto.Post{
		PostTime: msg.Time,
		UserName: msg.UserName,
		Id:       msg.ID,
		Data:     msg.Message,
	}
}
