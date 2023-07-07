package service

import (
	"fmt"
	"fmt/api/model"
	"fmt/api/store"
)

// User table에 값 존재여부 확인
// error handling 추가
func FindUser(userID int64) (*model.User, error) {
	for _, user := range store.Users {
		if user.ID == userID {
			return &user, nil
		}
	}
	return nil, fmt.Errorf("User not found")
}
