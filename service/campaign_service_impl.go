package service

import (
	"bwastartup-crowdfunding/entity"
	"bwastartup-crowdfunding/model"
	"bwastartup-crowdfunding/repository"
	"context"
	"errors"
	"fmt"
	"path/filepath"
	"strconv"
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

func (c *CampaignServiceImpl) Create(ctx context.Context, request model.CreateCampaignRequest) (model.GetCampaignResponse, error) {
	campaign := entity.Campaign{
		UserId:           request.UserId,
		Name:             request.Name,
		ShortDescription: request.ShortDescription,
		Description:      request.Description,
		Perks:            request.Perks,
		GoalAmount:       request.GoalAmount,
		Slug:             c.GenerateSlug(request.Name, request.UserId),
	}

	save, err := c.repository.Save(ctx, campaign)
	if err != nil {
		return model.GetCampaignResponse{}, err
	}

	return model.GetCampaignResponse{
		Id:               save.Id,
		Name:             save.Name,
		UserId:           save.UserId,
		GoalAmount:       save.GoalAmount,
		ShortDescription: save.ShortDescription,
		Slug:             save.Slug,
	}, nil
}

// GenerateSlug Generate slug : name-id
func (c *CampaignServiceImpl) GenerateSlug(name string, userId uint32) string {
	return strings.Join(strings.Split(strings.ToLower(strings.Trim(name, " ")), " "), "-") + "-" + strconv.Itoa(int(userId))
}

func (c *CampaignServiceImpl) Update(ctx context.Context, request model.CreateCampaignRequest, id uint32) (model.GetCampaignResponse, error) {
	campaign, err := c.repository.FindById(ctx, id)
	if err != nil {
		return model.GetCampaignResponse{}, err
	}
	if campaign.UserId != request.UserId {
		return model.GetCampaignResponse{}, errors.New("campaign: not an owner of the campaign")
	}

	campaign.Name = request.Name
	campaign.ShortDescription = request.ShortDescription
	campaign.Description = request.Description
	campaign.Perks = request.Perks
	campaign.GoalAmount = request.GoalAmount

	update, err := c.repository.Update(ctx, campaign)
	if err != nil {
		return model.GetCampaignResponse{}, err
	}

	return model.GetCampaignResponse{
		Id:               update.Id,
		UserId:           update.UserId,
		Name:             update.Name,
		ShortDescription: update.ShortDescription,
		ImageUrl:         update.CampaignImages[0].FileName,
		GoalAmount:       update.GoalAmount,
		CurrentAmount:    update.CurrentAmount,
		Slug:             update.Slug,
	}, nil
}

func (c *CampaignServiceImpl) SaveCampaignImage(ctx context.Context, request model.CreateCampaignImageRequest, fileName string) (string, error) {
	campaign, err := c.repository.FindById(ctx, request.CampaignId)
	if err != nil {
		return "", err
	}

	if campaign.UserId != request.UserId {
		return "", errors.New("campaign: not an owner of the campaign")
	}

	if request.IsPrimary {
		err := c.repository.MarkAllImagesAsNonPrimary(request.UserId)
		if err != nil {
			return "", err
		}
	}

	ext := filepath.Ext(fileName)
	// Validate image extension
	if ext != ".jpg" && ext != ".jpeg" && ext != ".png" {
		return "", errors.New("upload: invalid file extension")
	}
	generateName := c.GenerateCampaignImageName(request.CampaignId, request.UserId, fileName, ext)

	campaignImage := entity.CampaignImage{
		CampaignId: request.CampaignId,
		FileName:   generateName,
		IsPrimary:  request.IsPrimary,
	}

	_, err = c.repository.CreateImage(ctx, campaignImage)
	if err != nil {
		return "", err
	}
	return generateName, nil
}

func (c *CampaignServiceImpl) GenerateCampaignImageName(campaignId uint32, userId uint32, fileName string, extension string) string {
	name := strings.Join(strings.Split(fileName, " "), "-")
	return fmt.Sprintf("%d-user%d-%s-campaign%s", campaignId, userId, name, extension)
}
