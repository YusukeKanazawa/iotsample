package app

import (
	user "broker/tippy/user"
)
var repo = user.GetRepository()
func NewUser(userName string, password string){
	usr := user.NewUser(user.UserKey(userName), password)
	repo.Store(usr)
}
func GetUserList() *map[user.UserKey]user.User{
	return repo.GetAll()
}
