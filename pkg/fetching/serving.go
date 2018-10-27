package fetching

type Repository interface {
	GetChannelByTopic(string) (*RSS, error)
}

type Service interface {
	GetChannelByTopic(string) (*RSS, error)
}

func NewService(r Repository) Service {
	return &service{r}
}

type service struct {
	r Repository
}

func (s *service) GetChannelByTopic(topic string) (*RSS, error) {
	return s.r.GetChannelByTopic(topic)
}
