package repository

import (
	"bwastartup-crowdfunding/entity"
	"context"
	"gorm.io/gorm"
)

type CampaignRepositoryImpl struct {
	Db *gorm.DB
}

func NewCampaignRepositoryImpl(db *gorm.DB) *CampaignRepositoryImpl {
	return &CampaignRepositoryImpl{Db: db}
}

func (c *CampaignRepositoryImpl) FindAll(ctx context.Context) ([]entity.Campaign, error) {
	var campaigns []entity.Campaign
	err := c.Db.WithContext(ctx).Preload("CampaignImages", "campaign_images.is_primary=1").Find(&campaigns).Error
	if err != nil {
		return nil, err
	}
	return campaigns, nil
}

func (c *CampaignRepositoryImpl) FindByUserId(ctx context.Context, userId uint32) ([]entity.Campaign, error) {
	var campaigns []entity.Campaign
	err := c.Db.WithContext(ctx).Where("user_id=?", userId).Preload("CampaignImages", "campaign_images.is_primary=1").Find(&campaigns).Error
	if err != nil {
		return nil, err
	}
	return campaigns, nil
}
