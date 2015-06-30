package user

import (
	//	"log"
	//	"errors"
	"crypto/sha512"
)
//key type
type UserKey string

// Entity: Aggregate
type User struct {
	UserName UserKey
	Password [sha512.Size]byte
	clients  clientList
}

func (this *User) AddClient(client *Client) {
	this.clients.add(client)
}

// FactoryMethod
func NewUser(userName UserKey, password string) *User {
	return &User{
		UserName: userName,
		Password: sha512.Sum512([]byte(password)),
		clients:  clientList{},
	}
}
