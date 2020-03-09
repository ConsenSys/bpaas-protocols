package dto

type CompanyDataTransferType struct {
	ID                string
	Name              string
	Admin             *UserDataTransferType
	OrganizationCount uint
	UserCount         uint
}
