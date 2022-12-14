package campaign

import (
	"errors"
	"fmt"

	"github.com/gosimple/slug"
)

type Service interface {
	GetCampaigns(userID int) ([]Campaign, error)
	GetCampaign(input GetCampaignDetailInput) (Campaign, error)
	CreateCampaign(input CreateCampaignInput) (Campaign, error)
	UpdateCampaign(campaignID GetCampaignDetailInput, input CreateCampaignInput) (Campaign, error)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) *service {
	return &service{repository}
}

func (s *service) GetCampaigns(userID int) ([]Campaign, error) {
	if userID != 0 {
		campaigns, err := s.repository.FindByUserID(userID)

		if err != nil {
			return campaigns, err
		}

		return campaigns, nil
	}

	campaigns, err := s.repository.FindAll()

	if err != nil {
		return campaigns, err
	}

	return campaigns, nil
}

func (s *service) GetCampaign(input GetCampaignDetailInput) (Campaign, error) {
	campaign, err := s.repository.FindByID(input.ID)

	if err != nil {
		return campaign, err
	}

	return campaign, nil
}

func (s *service) CreateCampaign(input CreateCampaignInput) (Campaign, error) {
	campaign := Campaign{}
	campaign.Name = input.Name
	campaign.ShortDescription = input.ShortDescription
	campaign.Description = input.Description
	campaign.GoalAmount = input.GoalAmount
	campaign.Perks = input.Perks
	campaign.UserID = input.User.ID

	slugCandidate := fmt.Sprintf("%s %d", input.Name, input.User.ID)

	campaign.Slug = slug.Make(slugCandidate)

	newCampaign, err := s.repository.Save(campaign)

	if err != nil {
		return newCampaign, err
	}

	return newCampaign, nil
}

func (s *service) UpdateCampaign(campaignID GetCampaignDetailInput, campaign CreateCampaignInput) (Campaign, error) {
	campaigns, err := s.repository.FindByID(campaignID.ID)

	if err != nil {
		return campaigns, err
	}

	if campaigns.UserID != campaign.User.ID {
		return campaigns, errors.New("Not an owner of this campaign")
	}

	campaigns.Name = campaign.Name
	campaigns.ShortDescription = campaign.ShortDescription
	campaigns.Description = campaign.Description
	campaigns.GoalAmount = campaign.GoalAmount
	campaigns.Perks = campaign.Perks

	slugCandidate := fmt.Sprintf("%s %d", campaign.Name, campaign.User.ID)

	campaigns.Slug = slug.Make(slugCandidate)

	updatedCampaign, err := s.repository.Update(campaigns)

	if err != nil {
		return updatedCampaign, err
	}

	return updatedCampaign, nil

}
