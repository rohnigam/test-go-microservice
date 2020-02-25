package domain

import "errors"

var users = map[int64]*User{
	123: {123, "John", "Doe", "johndoe@gmail.com"},
}

func GetUser(userId int64) (*User, error) {
	user := users[userId]
	if user == nil {
		return nil, errors.New("No such user present!")
	}
	return user, nil
}
