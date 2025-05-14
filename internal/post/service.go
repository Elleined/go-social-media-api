package post

import (
	"errors"
	"social-media-application/internal/paging"
	"strings"
)

type (
	Service interface {
		save(authorId int, subject, content string) (id int64, err error)

		getAll(currentUserId int, isDeleted bool, pageRequest *paging.PageRequest) (*paging.Page[Post], error)
		getAllBy(currentUserId int, isDeleted bool, pageRequest *paging.PageRequest) (*paging.Page[Post], error)

		updateSubject(currentUserId int, postId int, newSubject string) (affectedRows int64, err error)
		updateContent(currentUserId, postId int, newContent string) (affectedRows int64, err error)
		updateAttachment(currentUserId, postId int, newAttachment string) (affectedRows int64, err error)

		deleteById(currentUserId, postId int) (affectedRows int64, err error)
	}

	ServiceImpl struct {
		repository Repository
	}
)

func (s ServiceImpl) save(authorId int, subject, content string) (id int64, err error) {
	if authorId <= 0 {
		return 0, errors.New("author id is required")
	}

	if strings.TrimSpace(subject) == "" {
		return 0, errors.New("subject is required")
	}

	if strings.TrimSpace(content) == "" {
		return 0, errors.New("content is required")
	}

	id, err = s.repository.save(authorId, subject, content)
	if err != nil {
		return 0, err
	}

	return id, nil
}

func (s ServiceImpl) getAll(currentUserId int, isDeleted bool, pageRequest *paging.PageRequest) (*paging.Page[Post], error) {
	if currentUserId <= 0 {
		return nil, errors.New("author id is required")
	}

	posts, err := s.repository.findAll(currentUserId, isDeleted, pageRequest)
	if err != nil {
		return nil, err
	}

	return posts, nil
}

func (s ServiceImpl) getAllBy(currentUserId int, isDeleted bool, pageRequest *paging.PageRequest) (*paging.Page[Post], error) {
	if currentUserId <= 0 {
		return nil, errors.New("author id is required")
	}

	posts, err := s.repository.findAllBy(currentUserId, isDeleted, pageRequest)
	if err != nil {
		return nil, err
	}

	return posts, nil
}

func (s ServiceImpl) updateSubject(currentUserId int, postId int, newSubject string) (affectedRows int64, err error) {
	if currentUserId <= 0 {
		return 0, errors.New("author id is required")
	}

	if postId <= 0 {
		return 0, errors.New("post id is required")
	}

	if strings.TrimSpace(newSubject) == "" {
		return 0, errors.New("new subject is required")
	}

	affectedRows, err = s.repository.updateSubject(currentUserId, postId, newSubject)
	if err != nil {
		return 0, err
	}

	if affectedRows <= 0 {
		return 0, errors.New("current user is not the author of post")
	}

	return affectedRows, nil
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

func NewService(repository Repository) Service {
	return &ServiceImpl{
		repository: repository,
	}
}
