package emoji

type (
	Service interface {
		save(name string) (id int64, err error)
		getAll() ([]Emoji, error)
		update(emojiId int, newName string) (affectedRows int64, err error)
		delete(emojiId int) (affectedRows int64, err error)
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

func (s ServiceImpl) save(name string) (id int64, err error) {
	//TODO implement me
	panic("implement me")
}

func (s ServiceImpl) getAll() ([]Emoji, error) {
	//TODO implement me
	panic("implement me")
}

func (s ServiceImpl) update(emojiId int, newName string) (affectedRows int64, err error) {
	//TODO implement me
	panic("implement me")
}

func (s ServiceImpl) delete(emojiId int) (affectedRows int64, err error) {
	//TODO implement me
	panic("implement me")
}
