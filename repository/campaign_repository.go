package repository

import (
	"bwastartup-crowdfunding/entity"
	"context"
)

type CampaignRepository interface {
	FindAll(ctx context.Context) ([]entity.Campaign, error)
	FindByUserId(ctx context.Context, userId uint32) ([]entity.Campaign, error)
	FindById(ctx context.Context, id uint32) (entity.Campaign, error)
	Save(ctx context.Context, campaign entity.Campaign) (entity.Campaign, error)
	Update(ctx context.Context, campaign entity.Campaign) (entity.Campaign, error)
}
