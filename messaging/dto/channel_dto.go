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

func ToDTOList(channels []*model.Channel) []*ChannelDTO {

	channelsDTO := []*ChannelDTO{}

	for _, channel := range channels {

		channelsDTO = append(channelsDTO, ToDTO(channel))

	}

	return channelsDTO
}

func ToDTO(channel *model.Channel) *ChannelDTO {

	return &ChannelDTO{
		ID:                 channel.ID,
		Name:               channel.Name,
		ConsortiumID:       channel.ConsortiumID,
		OriginalName:       channel.OriginalName,
		OrderersCount:      uint(len(channel.Orderers)),
		PeersCount:         uint(len(channel.Peers)),
		OrganizationsCount: uint(len(channel.Organizations)),
		ChaincodesCount:    uint(len(channel.Chaincodes)),
		Status:             channel.Status,
		Public:             channel.Public,
		ConsortiumName:     channel.Consortium.Name,
	}

}
