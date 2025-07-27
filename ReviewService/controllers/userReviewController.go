package controller

import (
	service "GoReview/services"
	"fmt"
	"net/http"
)

type UserReviewController struct {
	userReviewService service.UserReviewService
}

func NewUserReviewController(_userReviewService service.UserReviewService) *UserReviewController {
	return &UserReviewController{
		userReviewService: _userReviewService,
	}
}

func (urc *UserReviewController) CreateReviewController(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Controller Layer")
	w.Write([]byte("create review endpoints"))
	urc.userReviewService.CreateUserReview()
	
}