package model

type GetCampaignResponse struct {
	Id               uint32 `json:"id"`
	UserId           uint32 `json:"user_id"`
	Name             string `json:"name"`
	ShortDescription string `json:"short_description"`
	ImageUrl         string `json:"image_url"`
	GoalAmount       uint64 `json:"goal_amount"`
	CurrentAmount    uint64 `json:"current_amount"`
}
