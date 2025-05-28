package social

type (
	Service interface {
	}

	ServiceImpl struct {
		repository Repository
	}
)

func NewService(repository Repository) Service {
	return &ServiceImpl{
		repository: repository,
	}
}
