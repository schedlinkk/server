package users

import "time"

type UID string

type User struct {
	UID           UID       `json:"uid"`
	Username      string    `json:"username"`
	Friends       []string  `json:"friends"`
	CreatedAt     time.Time `json:"createdAt"`
	DayPlansCount int       `json:"dayPlansCount"`
}
