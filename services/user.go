package services

import "test-go-microservice/domain"

func GetUser(userId int64) (*domain.User, error) {
	return domain.GetUser(userId)
}
