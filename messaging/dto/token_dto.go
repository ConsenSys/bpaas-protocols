package dto

//userSvc "github.com/consensys/bpaas-user/service/user"

//TokenDataTransferType - Representation of a token, suitable for transmission as JSON
type TokenDataTransferType struct {
	ID       string
	Code     string
	User     *UserDataTransferType
	Type     string
	Used     bool
	DateUsed string
}
