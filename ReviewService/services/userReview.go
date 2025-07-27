package service

import (
	db "GoReview/db/repository"
	"GoReview/dto"
	"fmt"
)

type UserReviewService interface {
	CreateUserReview(payload *dto.UserReviewDto) error
}

type UserReviewServiceImpl struct {
	userReviewRepository db.ReviewRepository
}

func NewUserReviewServiceImpl(_userReviewRepository db.ReviewRepository) *UserReviewServiceImpl {
	return &UserReviewServiceImpl{
		userReviewRepository: _userReviewRepository,
	}
}

func(urs *UserReviewServiceImpl) CreateUserReview(payload *dto.UserReviewDto) error{
	fmt.Println("This is Service Layer")
	// err := urs.userReviewRepository.CreateUserReview(101, 202, 303, "Great stay with excellent service!", 4.5, time.Now(), time.Now())
	err := urs.userReviewRepository.CreateUserReview(payload)

	if err != nil{
		fmt.Println("Review Service -> Error inserting the value", err)
		return err
	}
	return nil
}
