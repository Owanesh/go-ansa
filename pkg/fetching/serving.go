package fetching

// Repository is the interface that defines the storage functions
// used by Fetching service
type Repository interface {
	GetChannelByTopic(string) (*RSS, error)
}

// Service defines the methods to interact with this Fetching service
type Service interface {
	GetChannelByTopic(string) (*RSS, error)
}

// NewService creates a Fetching service
func NewService(r Repository) Service {
	return &service{r}
}

type service struct {
	r Repository
}

func (s *service) GetChannelByTopic(topic string) (*RSS, error) {
	return s.r.GetChannelByTopic(topic)
}
