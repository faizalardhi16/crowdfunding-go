package campaign

import (
	"time"
)

type Campaign struct {
	ID               int       `json:"id"`
	UserID           int       `json:"user_id"`
	Name             string    `json:"name"`
	ShortDescription string    `json:"shortDescription"`
	Description      string    `json:"description"`
	Pearks           string    `json:"pearks"`
	BackerCount      int       `json:"backerCount"`
	GoalAmount       int       `json:"goalAmount"`
	CurrentAmount    int       `json:"currentAmount"`
	Slug             string    `json:"slug"`
	CreatedAt        time.Time `json:"createdAt"`
	UpdatedAt        time.Time `json:"updatedAt"`
	CampaignImages   []CampaignImage
}

type CampaignImage struct {
	ID         int       `json:"id"`
	CampaignID int       `json:"campaign_id"`
	FileName   string    `json:"fileName"`
	IsPrimary  int       `json:"isPrimary"`
	CreatedAt  time.Time `json:"createdAt"`
	UpdatedAt  time.Time `json:"updatedAt"`
}
