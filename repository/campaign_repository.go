package repository

import (
	"bwastartup-crowdfunding/entity"
	"context"
)

type CampaignRepository interface {
	FindAll(ctx context.Context) ([]entity.Campaign, error)
	FindByUserId(ctx context.Context, userId uint32) ([]entity.Campaign, error)
}
