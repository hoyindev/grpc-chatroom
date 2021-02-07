package auth

import (
	"backend-grpc-challenge/internal/misc"
	"crypto/sha256"
	"encoding/base64"
)

//MakeHashedPW base64(hash pw + client salt)
func MakeClientHashedPW(input string) string {
	str := input + getClientSalt()
	h := sha256.New()
	h.Write([]byte(str))
	return base64.StdEncoding.EncodeToString(h.Sum(nil))
}

//MakeHashedPW base64(hash pw + server salt)
func MakeServerHashedPW(input string) string {
	str := input + getServerSalt()
	h := sha256.New()
	h.Write([]byte(str))
	return base64.StdEncoding.EncodeToString(h.Sum(nil))
}

func getClientSalt() string {
	return misc.GetEnv("chatroomSalt", "salt")
}

func getServerSalt() string {
	return misc.GetEnv("chatroomServerSalt", "itisthesaltforserver")
}
