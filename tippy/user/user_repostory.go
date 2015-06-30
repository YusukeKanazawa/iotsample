package user

import (
	"errors"
	"log"
)

// Repository
type UserRepository struct {
	store map[UserKey]User
}

func (this *UserRepository) Store(user *User) error {
	if user == nil {
		return errors.New("Object is nil.")
	}
	if user.UserName == "" {
		return errors.New("Blank is not allowed for UserName.")
	}
	this.store[user.UserName] = *user
	return nil
}
func (this *UserRepository) Find(userName UserKey) *User {
	user := this.store[userName]
	return &user
}
func (this *UserRepository) GetAll() *map[UserKey]User{
	return &(this.store)
}

// Singleton
var instance *UserRepository = &UserRepository{
	store: make(map[UserKey]User, 0),
}
func GetRepository() *UserRepository {
	log.Println("get repo.")
	return instance
}
