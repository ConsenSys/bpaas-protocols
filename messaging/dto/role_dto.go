package dto

type RoleDataTransferType struct {
	ID    string
	Name  string
	Title string
}

type CompanyRoleDataTransferType struct {
	Company CompanyDataTransferType
	Role    RoleDataTransferType
}
