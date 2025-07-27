package db

import "database/sql"

type ReviewRepository interface {
}

type ReviewRepositoryImpl struct {
	db *sql.DB
}

func NewReviewRepositoryImpl(_db *sql.DB) *ReviewRepositoryImpl{
	return &ReviewRepositoryImpl{
		db : _db,
	}
}