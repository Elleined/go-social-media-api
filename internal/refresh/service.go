package refresh

type (
	Service interface {
	}

	ServiceImpl struct {
		Repository
	}
)

func NewService(repository Repository) Service {
	return &ServiceImpl{
		Repository: repository,
	}
}
