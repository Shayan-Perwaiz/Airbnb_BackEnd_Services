package model

import "time"

type Hotel struct {
	Id        int64
	Name      string
	Address   string
	City      string
	CreatedAt time.Time
	UpdatedAt time.Time
}