package router

import (
	controller "GoReview/controllers"

	"github.com/go-chi/chi/v5"
)

type ReviewRouter struct {
	userReviewController *controller.UserReviewController
}

func NewReviewRouter(_userReviewController *controller.UserReviewController) Router{
	return &ReviewRouter{
		userReviewController : _userReviewController,
	}

}

func (rr *ReviewRouter) RegisterRouter(r chi.Router) {
	r.Post("/create", rr.userReviewController.CreateReviewController)
}
