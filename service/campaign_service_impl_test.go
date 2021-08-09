package service

import (
	"bwastartup-crowdfunding/model"
	"bwastartup-crowdfunding/repository"
	"context"
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

var campaignRepository repository.CampaignRepository
var campaignService CampaignService

func TestCampaignServiceImpl_FindAll(t *testing.T) {
	ctx := context.Background()
	campaigns, err := campaignService.FindAll(ctx)
	if err != nil {
		t.Error(err.Error())
	}
	for _, campaign := range campaigns {
		fmt.Println(campaign)
	}
}
func TestCampaignServiceImpl_FindByUserId(t *testing.T) {
	ctx := context.Background()
	campaigns, err := campaignService.FindByUserId(ctx, 1)
	if err != nil {
		t.Error(err.Error())
	}
	for _, campaign := range campaigns {
		fmt.Println(campaign)
	}
}

func TestCampaignServiceImpl_FindById(t *testing.T) {
	ctx := context.Background()
	response, err := campaignService.FindById(ctx, 1)
	if err != nil {
		t.Error(err.Error())
	}
	fmt.Println(response)
}

func TestCampaignServiceImpl_GenerateSlug(t *testing.T) {
	slug := campaignService.GenerateSlug("Test Slug", 10)
	assert.Equal(t, "test-slug-10", slug)
}

func TestCampaignServiceImpl_Create(t *testing.T) {
	request := model.CreateCampaignRequest{
		Name:             "Test Create Campaign",
		ShortDescription: "shot",
		Description:      "longgggg",
		GoalAmount:       500_000,
		Perks:            "Ilmu yang bermanfaat,pahala",
		UserId:           2,
	}

	ctx := context.Background()
	response, err := campaignService.Create(ctx, request)
	if err != nil {
		t.Error(err.Error())
	}
	fmt.Println(response)
}

func TestCampaignServiceImpl_Update(t *testing.T) {
	ctx := context.Background()
	request := model.CreateCampaignRequest{
		Name:             "Test Update From Service",
		ShortDescription: "ngetes",
		Description:      "-",
		GoalAmount:       1_000_000_000,
		Perks:            "pahalaaa,pahalaaa",
		UserId:           1,
	}

	update, err := campaignService.Update(ctx, request, 1)
	if err != nil {
		t.Error(err.Error())
	}
	fmt.Println(update)
}
