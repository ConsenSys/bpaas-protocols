package dto

type InviteDataTransferType struct {
	ID         string
	Sender     UserDataTransferType
	Receiver   UserDataTransferType
	Company    CompanyDataTransferType
	Status     string
	IsOutgoing bool
	RoleID     string
}
