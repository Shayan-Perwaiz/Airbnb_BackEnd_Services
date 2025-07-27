package service

import (
	db "GoReview/db/repository"
	"fmt"
	"time"
)

type UserReviewService interface {
	CreateUserReview() error
}

type UserReviewServiceImpl struct {
	userReviewRepository db.ReviewRepository
}

func NewUserReviewServiceImpl(_userReviewRepository db.ReviewRepository) *UserReviewServiceImpl {
	return &UserReviewServiceImpl{
		userReviewRepository: _userReviewRepository,
	}
}

func(urs *UserReviewServiceImpl) CreateUserReview() error{
	fmt.Println("This is Service Layer")
	err := urs.userReviewRepository.CreateUserReview(101, 202, 303, "Great stay with excellent service!", 4.5, time.Now(), time.Now())
	if err != nil{
		fmt.Println("Review Service -> Error inserting the value", err)
		return err
	}
	return nil
}
