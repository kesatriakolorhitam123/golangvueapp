package transaction

import "golangvueapp/user"

type GetCampaignTransactionInput struct {
	ID   int `uri:"id" binding:"required"`
	User user.User
}
