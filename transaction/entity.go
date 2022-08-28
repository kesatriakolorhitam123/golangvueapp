package transaction

import (
	"golangvueapp/campaign"
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
	Campaign   campaign.Campaign
	CreatedAt  time.Time
	UpdatedAt  time.Time
}
