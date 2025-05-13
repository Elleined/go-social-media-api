package post

import (
	"errors"
	"strings"
)

type Service interface {
	save(authorId int, subject, content string) (id int64, err error)

	findAll(currentUserId, limit, offset int) ([]Post, error)
	findAllBy(currentUserId, limit, offset int) ([]Post, error)

	updateSubject(currentUserId int, postId int, newSubject string) (affectedRows int64, err error)
	updateContent(currentUserId, postId int, newContent string) (affectedRows int64, err error)
	updateAttachment(currentUserId, postId int, newAttachment string) (affectedRows int64, err error)

	deleteById(currentUserId, postId int) (affectedRows int64, err error)
}

type ServiceImpl struct {
	repository Repository
}

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

func (s ServiceImpl) findAll(currentUserId, limit, offset int) ([]Post, error) {
	if limit < 0 {
		return nil, errors.New("limit is required")
	}

	if offset < 0 {
		return nil, errors.New("offset is required")
	}

	posts, err := s.repository.findAll(currentUserId, limit, offset)
	if err != nil {
		return nil, err
	}

	return posts, nil
}

func (s ServiceImpl) findAllBy(currentUserId, limit, offset int) ([]Post, error) {
	if currentUserId <= 0 {
		return nil, errors.New("author id is required")
	}

	if limit < 0 {
		return nil, errors.New("limit is required")
	}

	if offset < 0 {
		return nil, errors.New("offset is required")
	}

	posts, err := s.repository.findAllBy(currentUserId, limit, offset)
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

	return affectedRows, nil
}

func NewService(repository Repository) Service {
	return &ServiceImpl{
		repository: repository,
	}
}
