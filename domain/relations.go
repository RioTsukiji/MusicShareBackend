package domain

type Relations struct {
	Follow   int64 `json:"follow"`
	Follower int64 `json:"follower"`
}
