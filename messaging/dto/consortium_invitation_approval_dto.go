package dto

type ConsortiumInvitationApprovalDTO struct {
	ConsortiumInvitationID string
	OrganizationID         string
	Status                 string
	Invitation             InvitationDTO
}
