package repository

import (
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
		fmt.Println(campaign)
	}
}
