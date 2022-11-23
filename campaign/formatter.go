package campaign

import (
	"strings"
)

type CampaignFormatter struct {
	ID               int    `json:"id"`
	UserID           int    `json:"user_id"`
	Name             string `json:"name"`
	ShortDescription string `json:"short_description"`
	ImageURL         string `json:"image_url"`
	GoalAmount       int    `json:"goal_amount"`
	CurrentAmount    int    `json:"current_amount"`
	Slug             string `json:"slug"`
}

type CampaignDetailFormatter struct {
	ID               int                    `json:"id"`
	UserID           int                    `json:"user_id"`
	Name             string                 `json:"name"`
	ShortDescription string                 `json:"short_description"`
	ImageURL         string                 `json:"image_url"`
	GoalAmount       int                    `json:"goal_amount"`
	CurrentAmount    int                    `json:"current_amount"`
	Perks            []string               `json:"perks"`
	Slug             string                 `json:"slug"`
	Description      string                 `json:"description"`
	User             UserCampaignDetail     `json:"user"`
	Images           []CampaignImagesDetail `json:"images"`
}

type UserCampaignDetail struct {
	Name           string `json:"name"`
	AvatarFileName string `json:"avatar_url"`
}

type CampaignImagesDetail struct {
	ImageURL  string `json:"image_url"`
	IsPrimary int    `json:"is_primary"`
}

func FormatCampaign(campaign Campaign) CampaignFormatter {
	campaignFormatter := CampaignFormatter{}
	campaignFormatter.ID = campaign.ID
	campaignFormatter.UserID = campaign.UserID
	campaignFormatter.Name = campaign.Name
	campaignFormatter.ShortDescription = campaign.ShortDescription
	campaignFormatter.GoalAmount = campaign.GoalAmount
	campaignFormatter.CurrentAmount = campaign.CurrentAmount
	campaignFormatter.ImageURL = ""
	campaignFormatter.Slug = campaign.Slug

	if len(campaign.CampaignImages) > 0 {
		campaignFormatter.ImageURL = campaign.CampaignImages[0].FileName
	}

	return campaignFormatter
}

func FormatCampaigns(campaigns []Campaign) []CampaignFormatter {
	campaignsFormatter := []CampaignFormatter{}

	for _, r := range campaigns {
		campaignFormatter := FormatCampaign(r)
		campaignsFormatter = append(campaignsFormatter, campaignFormatter)
	}

	return campaignsFormatter
}

func FormatDetailCampaign(campaign Campaign) CampaignDetailFormatter {
	f := CampaignDetailFormatter{}
	fileThumb := ""

	//binding user
	user := UserCampaignDetail{}
	user.Name = campaign.User.Name
	user.AvatarFileName = campaign.User.AvatarFileName

	//binding Images and binding as primary file
	image := CampaignImagesDetail{}
	var images []CampaignImagesDetail

	for _, d := range campaign.CampaignImages {
		if d.IsPrimary == 1 {
			fileThumb = d.FileName
		}
		image.ImageURL = d.FileName
		image.IsPrimary = d.IsPrimary
		images = append(images, image)
	}

	f.ID = campaign.ID
	f.Name = campaign.Name
	f.ShortDescription = campaign.ShortDescription
	f.ImageURL = fileThumb
	f.GoalAmount = campaign.GoalAmount
	f.CurrentAmount = campaign.CurrentAmount
	f.UserID = campaign.UserID
	f.Description = campaign.Description
	f.User = user
	f.Perks = strings.Split(campaign.Perks, ", ")
	f.Slug = campaign.Slug
	f.Images = images

	return f
}
