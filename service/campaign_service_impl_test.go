package service

import (
	"bwastartup-crowdfunding/repository"
	"context"
	"fmt"
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
