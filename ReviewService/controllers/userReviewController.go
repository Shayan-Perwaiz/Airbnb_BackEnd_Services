package controller

import (
	"GoReview/dto"
	service "GoReview/services"
	"GoReview/utils"
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
	var payload dto.UserReviewDto
	fmt.Println("Controller Layer")
	if RequestBodyError := utils.ReadJsonBody(r, &payload); RequestBodyError != nil{
		utils.WriteJsonErrorResponse(w, "Invalid Input data", "false")
		return
	}
	if err := urc.userReviewService.CreateUserReview(&payload); err != nil{
		utils.WriteJsonErrorResponse(w, "Error creating review", "false")
		return
	}
	utils.WriteJsonSuccessResponse(w, "user review created success", "true")
	
}