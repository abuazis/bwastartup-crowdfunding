package service

import (
	"bwastartup-crowdfunding/model"
	"bwastartup-crowdfunding/repository"
	"context"
	"strings"
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
			Slug:             campaign.Slug,
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
			Slug:             campaign.Slug,
		})
	}
	return campaignModels, nil
}

func (c *CampaignServiceImpl) FindById(ctx context.Context, id uint32) (model.GetCampaignDetailResponse, error) {
	campaign, err := c.repository.FindById(ctx, id)
	if err != nil {
		return model.GetCampaignDetailResponse{}, err
	}

	campaignResponse := model.GetCampaignDetailResponse{
		Campaign: model.CampaignDetailResponse{
			Id:               campaign.Id,
			Title:            campaign.Name,
			ShortDescription: campaign.ShortDescription,
			Description:      campaign.Description,
			GoalAmount:       campaign.GoalAmount,
			CurrentAmount:    campaign.CurrentAmount,
			Perks:            strings.Split(campaign.Perks, ","),
		},
		User: model.UserDetailResponse{
			Id:        campaign.User.Id,
			Name:      campaign.User.Name,
			AvatarUrl: model.BASE_URL + "uploads/users/" + campaign.User.AvatarFileName,
		},
	}

	for _, image := range campaign.CampaignImages {
		campaignResponse.Campaign.Images = append(campaignResponse.Campaign.Images, model.ImageDetailResponse{
			ImageUrl:  model.BASE_URL + "uploads/campaigns/" + image.FileName,
			IsPrimary: image.IsPrimary,
		})
	}

	return campaignResponse, nil
}
