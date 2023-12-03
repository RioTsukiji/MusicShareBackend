package domain

import "time"

type Music struct {
	ID       int64     `json:"id"`
	Title    string    `json:"title"`
	Artist   string    `json:"artist"`
	Link     string    `json:"link"`
	SharedAt time.Time `json:"shared_at"`
	UserID   int64     `json:"user_name"`
}
