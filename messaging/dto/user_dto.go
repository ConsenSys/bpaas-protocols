package dto

type UserDataTransferType struct {
	ID                 string
	Name               string
	Email              string
	PlatformRole       RoleDataTransferType
	CompanyRoles       []*CompanyRoleDataTransferType
	PhoneNumber        string
	CountryCode        string
	IsVerified         bool
	Status             string
	Gender             string
	VerificationToken_ string //Only to be returned when EXECUTION_MODE=test
}
