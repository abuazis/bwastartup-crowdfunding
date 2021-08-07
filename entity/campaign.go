package entity

import (
	"time"
)

type Campaign struct {
	Id, UserId, BackerCount                          uint32
	Name, ShortDescription, Description, Perks, Slug string
	GoalAmount, CurrentAmount                        uint64
	CreatedAt                                        time.Time
	UpdatedAt                                        *time.Time
	CampaignImages                                   []CampaignImage
	User                                             User
}

type CampaignImage struct {
	Id, CampaignId uint32
	FileName       string
	IsPrimary      bool
	CreatedAt      time.Time
	UpdatedAt      *time.Time
}
