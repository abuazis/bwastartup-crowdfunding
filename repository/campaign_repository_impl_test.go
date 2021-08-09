package repository

import (
	"bwastartup-crowdfunding/entity"
	"context"
	"fmt"
	"testing"
)

var campaignRepository CampaignRepository

func TestCampaignRepositoryImpl_FindByUserId(t *testing.T) {
	ctx := context.Background()
	campaigns, err := campaignRepository.FindByUserId(ctx, 1)
	if err != nil {
		t.Error(err.Error())
	}
	for _, campaign := range campaigns {
		//fmt.Println(campaign)
		fmt.Println(campaign)
		fmt.Println()
	}
}

func TestCampaignRepositoryImpl_FindAll(t *testing.T) {
	ctx := context.Background()
	campaigns, err := campaignRepository.FindAll(ctx)
	if err != nil {
		t.Error(err.Error())
	}
	for _, campaign := range campaigns {
		//fmt.Println(campaign)
		fmt.Println(campaign)
		fmt.Println()
	}
}

func TestCampaignRepositoryImpl_FindById(t *testing.T) {
	ctx := context.Background()
	campaign, err := campaignRepository.FindById(ctx, 1)
	if err != nil {
		t.Error(err.Error())
	}
	fmt.Println(campaign)
}

func TestCampaignRepositoryImpl_Update(t *testing.T) {
	ctx := context.Background()
	campaign := entity.Campaign{
		Id:               5,
		UserId:           1,
		Name:             "Test Update Repository",
		ShortDescription: "shorttttt",
		Description:      "long.",
		Perks:            "pahala,pahala",
		Slug:             "test-update-repository",
		GoalAmount:       1_000_000,
	}

	update, err := campaignRepository.Update(ctx, campaign)
	if err != nil {
		t.Error(err.Error())
	}
	fmt.Println(update)
}
