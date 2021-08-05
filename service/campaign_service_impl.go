package service

import (
	"bwastartup-crowdfunding/model"
	"bwastartup-crowdfunding/repository"
	"context"
)

type CampaignServiceImpl struct {
	repository repository.CampaignRepository
}

func NewCampaignServiceImpl(repository repository.CampaignRepository) *CampaignServiceImpl {
	return &CampaignServiceImpl{repository: repository}
}

func (c *CampaignServiceImpl) FindAll(ctx context.Context) ([]model.GetCampaignResponse, error) {
	campaigns, err := c.repository.FindAll(ctx)
	if err != nil {
		return nil, err
	}
	var campaignModels []model.GetCampaignResponse
	for _, campaign := range campaigns {
		campaignModels = append(campaignModels, model.GetCampaignResponse{
			Id:               campaign.Id,
			UserId:           campaign.UserId,
			Name:             campaign.Name,
			ShortDescription: campaign.ShortDescription,
			ImageUrl:         model.BASE_URL + "uploads/campaigns/" + campaign.CampaignImages[0].FileName,
			GoalAmount:       campaign.GoalAmount,
			CurrentAmount:    campaign.CurrentAmount,
		})
	}

	return campaignModels, nil
}

func (c *CampaignServiceImpl) FindByUserId(ctx context.Context, id uint32) ([]model.GetCampaignResponse, error) {
	campaigns, err := c.repository.FindByUserId(ctx, id)
	if err != nil {
		return nil, err
	}
	var campaignModels []model.GetCampaignResponse
	for _, campaign := range campaigns {
		campaignModels = append(campaignModels, model.GetCampaignResponse{
			Id:               campaign.Id,
			UserId:           campaign.UserId,
			Name:             campaign.Name,
			ShortDescription: campaign.ShortDescription,
			ImageUrl:         model.BASE_URL + "uploads/campaigns/" + campaign.CampaignImages[0].FileName,
			GoalAmount:       campaign.GoalAmount,
			CurrentAmount:    campaign.CurrentAmount,
		})
	}
	return campaignModels, nil
}
