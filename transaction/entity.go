package transaction

import (
	"golangvueapp/user"
	"time"
)

type Transaction struct {
	ID         int
	CampaignId int
	UserID     int
	Amount     int
	Status     string
	Code       string
	User       user.User
	CreatedAt  time.Time
	UpdatedAt  time.Time
}
