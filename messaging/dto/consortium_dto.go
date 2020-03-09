package dto

type ConsortiumDTO struct {
	ID                         string
	Name                       string
	OriginalName               string
	Description                string
	OrganizationsCount         uint
	ChannelsCount              uint
	ConsortiumInvitationsCount uint
	Organizations              []*OrganizationDTO
}

type OrganizationDTO struct {
	ID           string
	Name         string
	OriginalName string
}
