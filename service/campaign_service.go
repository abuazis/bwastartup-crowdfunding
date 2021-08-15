package service

import (
	"bwastartup-crowdfunding/model"
	"context"
)

type CampaignService interface {
	FindAll(ctx context.Context) ([]model.GetCampaignResponse, error)
	FindByUserId(ctx context.Context, id uint32) ([]model.GetCampaignResponse, error)
	FindById(ctx context.Context, id uint32) (model.GetCampaignDetailResponse, error)
	Create(ctx context.Context, request model.CreateCampaignRequest) (model.GetCampaignResponse, error)
	GenerateSlug(name string, id uint32) string
	Update(ctx context.Context, request model.CreateCampaignRequest, id uint32) (model.GetCampaignResponse, error)
	SaveCampaignImage(ctx context.Context, request model.CreateCampaignImageRequest, fileName string) (string, error)
	GenerateCampaignImageName(campaignId uint32, userId uint32, fileName string, extension string) string
}
