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

func (c *CampaignRepositoryImpl) FindById(ctx context.Context, id uint32) (entity.Campaign, error) {
	var campaign entity.Campaign
	err := c.Db.WithContext(ctx).Preload("User").Preload("CampaignImages", func(db *gorm.DB) *gorm.DB {
		return db.Order("campaign_images.is_primary DESC")
	}).Where("id=?", id).First(&campaign).Error
	if err != nil {
		return entity.Campaign{}, err
	}
	return campaign, nil
}

func (c *CampaignRepositoryImpl) Save(ctx context.Context, campaign entity.Campaign) (entity.Campaign, error) {
	err := c.Db.WithContext(ctx).Create(&campaign).Error
	if err != nil {
		return entity.Campaign{}, err
	}
	return campaign, nil
}

func (c *CampaignRepositoryImpl) Update(ctx context.Context, campaign entity.Campaign) (entity.Campaign, error) {
	err := c.Db.WithContext(ctx).Model(&entity.Campaign{}).Where("id=?", campaign.Id).Updates(campaign).Error
	if err != nil {
		return entity.Campaign{}, err
	}
	return campaign, nil
}

func (c *CampaignRepositoryImpl) CreateImage(ctx context.Context, image entity.CampaignImage) (entity.CampaignImage, error) {
	err := c.Db.WithContext(ctx).Create(&image).Error
	if err != nil {
		return entity.CampaignImage{}, err
	}
	return image, nil
}

func (c *CampaignRepositoryImpl) MarkAllImagesAsNonPrimary(campaignId uint32) error {
	err := c.Db.Model(&entity.CampaignImage{}).Where("campaign_id=?", campaignId).Update("is_primary", false).Error
	if err != nil {
		return err
	}
	return nil
}
