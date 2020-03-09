package dto

type ConsortiumDTO struct {
	ID           string
	Name         string
	OriginalName string
	Description  string
}

type ChannelDTO struct {
	ConsortiumID string
	Name         string
	OriginalName string
	Status       string
	Public       bool
}

type InvitationDTO struct {
	ConsortiumID      string
	InviterOrgID      string
	InvitedEmail      string
	InvitedID         string
	InitialOrgsCount  uint
	SelectedPeers     string
	InvitedOrgID      string
	ChannelID         string
	ApprovedOrgsCount uint
	Status            string
	Consortium        ConsortiumDTO
	Channel           ChannelDTO
	//InvitationApprovals []*ConsortiumInvitationApproval //HasMany Relationship
}

// Invitation             ConsortiumInvitation `gorm:"foreignkey:ConsortiumInvitationID"`

type ConsortiumInvitationApprovalDTO struct {
	ConsortiumInvitationID string
	OrganizationID         string
	Status                 string
	Invitation             InvitationDTO
}
