package vk2go

type Publication struct {
	Id   int	`json:"id"`
	Text string	`json:"text"`
}

type PublicationsStorage interface {
	IsPublished(pubId int, domain string) (bool, error)
	MarkAsPublished(pubId int, domain string) error
}
