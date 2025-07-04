package post

import (
	"errors"
	"social-media-application/internal/paging"
	"strings"
)

type (
	Service interface {
		save(authorId int, content, attachment string) (id int64, err error)

		getById(postId int) (Post, error)
		getAll(currentUserId int, isDeleted bool, request *paging.PageRequest) (*paging.Page[Post], error)
		getAllBy(currentUserId int, isDeleted bool, request *paging.PageRequest) (*paging.Page[Post], error)

		updateContent(currentUserId, postId int, newContent string) (affectedRows int64, err error)
		updateAttachment(currentUserId, postId int, newAttachment string) (affectedRows int64, err error)

		deleteById(currentUserId, postId int) (affectedRows int64, err error)
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

func (s ServiceImpl) save(authorId int, content, attachment string) (id int64, err error) {
	if authorId <= 0 {
		return 0, errors.New("author id is required")
	}

	if strings.TrimSpace(content) == "" {
		return 0, errors.New("content is required")
	}

	id, err = s.repository.save(authorId, content, attachment)
	if err != nil {
		return 0, err
	}

	return id, nil
}

func (s ServiceImpl) getById(postId int) (Post, error) {
	if postId <= 0 {
		return Post{}, errors.New("post id is required")
	}

	post, err := s.repository.findById(postId)
	if err != nil {
		return Post{}, err
	}

	return post, nil
}

func (s ServiceImpl) getAll(currentUserId int, isDeleted bool, request *paging.PageRequest) (*paging.Page[Post], error) {
	if currentUserId <= 0 {
		return nil, errors.New("author id is required")
	}

	posts, err := s.repository.findAll(currentUserId, isDeleted, request)
	if err != nil {
		return nil, err
	}

	return posts, nil
}

func (s ServiceImpl) getAllBy(currentUserId int, isDeleted bool, request *paging.PageRequest) (*paging.Page[Post], error) {
	if currentUserId <= 0 {
		return nil, errors.New("author id is required")
	}

	posts, err := s.repository.findAllBy(currentUserId, isDeleted, request)
	if err != nil {
		return nil, err
	}

	return posts, nil
}

func (s ServiceImpl) updateContent(currentUserId, postId int, newContent string) (affectedRows int64, err error) {
	if currentUserId <= 0 {
		return 0, errors.New("author id is required")
	}

	if postId <= 0 {
		return 0, errors.New("post id is required")
	}

	if strings.TrimSpace(newContent) == "" {
		return 0, errors.New("new content is required")
	}

	affectedRows, err = s.repository.updateContent(currentUserId, postId, newContent)
	if err != nil {
		return 0, err
	}

	if affectedRows <= 0 {
		return 0, errors.New("current user is not the author of post")
	}

	return affectedRows, nil
}

func (s ServiceImpl) updateAttachment(currentUserId, postId int, newAttachment string) (affectedRows int64, err error) {
	if currentUserId <= 0 {
		return 0, errors.New("author id is required")
	}

	if postId <= 0 {
		return 0, errors.New("post id is required")
	}

	if strings.TrimSpace(newAttachment) == "" {
		return 0, errors.New("new attachment is required")
	}

	affectedRows, err = s.repository.updateAttachment(currentUserId, postId, newAttachment)
	if err != nil {
		return 0, err
	}

	if affectedRows <= 0 {
		return 0, errors.New("current user is not the author of post")
	}

	return affectedRows, nil
}

func (s ServiceImpl) deleteById(currentUserId, postId int) (affectedRows int64, err error) {
	if currentUserId <= 0 {
		return 0, errors.New("author id is required")
	}

	if postId <= 0 {
		return 0, errors.New("post id is required")
	}

	affectedRows, err = s.repository.deleteById(currentUserId, postId)
	if err != nil {
		return 0, err
	}

	if affectedRows <= 0 {
		return 0, errors.New("current user is not the author of post")
	}

	return affectedRows, nil
}
