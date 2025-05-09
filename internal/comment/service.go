package comment

type Service interface {
}

type ServiceImpl struct {
	repository Repository
}

func NewService(repository Repository) Service {
	return &ServiceImpl{
		repository: repository,
	}
}
