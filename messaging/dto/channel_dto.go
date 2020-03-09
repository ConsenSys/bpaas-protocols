package dto

//import model "github.com/consensys/bpaas-network/persistence/model"

type ChannelDTO struct {
	ID                 string
	Name               string
	ConsortiumID       string
	OriginalName       string
	OrderersCount      uint
	PeersCount         uint
	OrganizationsCount uint
	ChaincodesCount    uint
	Status             string
	Public             bool
	ConsortiumName     string
}
