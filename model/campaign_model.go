package model

type GetCampaignResponse struct {
	Id               uint32 `json:"id"`
	UserId           uint32 `json:"user_id"`
	Name             string `json:"name"`
	ShortDescription string `json:"short_description"`
	ImageUrl         string `json:"image_url"`
	GoalAmount       uint64 `json:"goal_amount"`
	CurrentAmount    uint64 `json:"current_amount"`
	Slug             string `json:"slug"`
}

type GetCampaignDetailRequest struct {
	Id uint32 `uri:"id" binding:"required"`
}

type GetCampaignDetailResponse struct {
	Campaign CampaignDetailResponse `json:"campaign"`
	User     UserDetailResponse     `json:"user"`
}

type CampaignDetailResponse struct {
	Id               uint32                `json:"id"`
	Title            string                `json:"title"`
	ShortDescription string                `json:"short_description"`
	Description      string                `json:"description"`
	GoalAmount       uint64                `json:"goal_amount"`
	CurrentAmount    uint64                `json:"current_amount"`
	Perks            []string              `json:"perks"`
	Images           []ImageDetailResponse `json:"images"`
}

type ImageDetailResponse struct {
	ImageUrl  string `json:"image_url"`
	IsPrimary bool   `json:"is_primary"`
}

type UserDetailResponse struct {
	Id        uint32 `json:"id"`
	Name      string `json:"name"`
	AvatarUrl string `json:"avatar_url"`
}

type CreateCampaignRequest struct {
	Name             string `json:"name" binding:"required"`
	ShortDescription string `json:"short_description" binding:"required"`
	Description      string `json:"description" binding:"required"`
	GoalAmount       uint64 `json:"goal_amount" binding:"required"`
	Perks            string `json:"perks" binding:"required"`
	UserId           uint32
}
type CreateCampaignResponse struct {
	Name             string `json:"name"`
	ShortDescription string `json:"short_description"`
	Description      string `json:"description"`
	GoalAmount       uint64 `json:"goal_amount"`
	Perks            string `json:"perks"`
}
