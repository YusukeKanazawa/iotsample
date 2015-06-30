package user

import (
	"log"
	"testing"
)

func TestUser(t *testing.T) {
	repository := NewUserRepository()
	user1 := NewUser("User1", "password")
	repository.Store(user1)
	log.Println(user1)
	client1 := NewClient("client#1", "client")
	user1.AddClient(client1)
	repository.Store(user1)
	log.Println(repository.Find("User1"))
}
