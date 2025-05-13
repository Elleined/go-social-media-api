package comment

import (
	"errors"
	"strings"
)

type Service interface {
	save(authorId, postId int, content string) (id int64, err error)

	getAll(currentUserId, postId, limit, offset int) ([]Comment, error)

	updateContent(currentUserId, postId, commentId int, newContent string) (affectedRows int64, err error)
	updateAttachment(currentUserId, postId, commentId int, newAttachment string) (affectedRows int64, err error)

	deleteById(currentUserId, postId, commentId int) (affectedRows int64, err error)
}

type ServiceImpl struct {
	repository Repository
}

func NewService(repository Repository) Service {
	return &ServiceImpl{
		repository: repository,
	}
}

func (s ServiceImpl) save(authorId, postId int, content string) (id int64, err error) {
	if authorId <= 0 {
		return 0, errors.New("author is required")
	}

	if postId <= 0 {
		return 0, errors.New("post is required")
	}

	if strings.TrimSpace(content) == "" {
		return 0, errors.New("content is required")
	}

	id, err = s.repository.save(authorId, postId, content)
	if err != nil {
		return 0, err
	}

	return id, nil
}

func (s ServiceImpl) getAll(currentUserId, postId, limit, offset int) ([]Comment, error) {
	if currentUserId <= 0 {
		return nil, errors.New("currentUserId is required")
	}

	if postId <= 0 {
		return nil, errors.New("postId is required")
	}

	if limit < 0 {
		return nil, errors.New("limit is required")
	}

	if offset < 0 {
		return nil, errors.New("offset is required")
	}

	comments, err := s.repository.findAll(currentUserId, postId, limit, offset)
	if err != nil {
		return nil, err
	}

	return comments, nil
}

func (s ServiceImpl) updateContent(currentUserId, postId, commentId int, newContent string) (affectedRows int64, err error) {
	if currentUserId <= 0 {
		return 0, errors.New("currentUserId is required")
	}

	if postId <= 0 {
		return 0, errors.New("postId is required")
	}

	if commentId <= 0 {
		return 0, errors.New("commentId is required")
	}

	if strings.TrimSpace(newContent) == "" {
		return 0, errors.New("newContent is required")
	}

	affectedRows, err = s.repository.updateContent(currentUserId, postId, commentId, newContent)
	if err != nil {
		return 0, err
	}

	return affectedRows, nil
}

func (s ServiceImpl) updateAttachment(currentUserId, postId, commentId int, newAttachment string) (affectedRows int64, err error) {
	if currentUserId <= 0 {
		return 0, errors.New("currentUserId is required")
	}

	if postId <= 0 {
		return 0, errors.New("postId is required")
	}

	if commentId <= 0 {
		return 0, errors.New("commentId is required")
	}

	if strings.TrimSpace(newAttachment) == "" {
		return 0, errors.New("newAttachment is required")
	}

	affectedRows, err = s.repository.updateAttachment(currentUserId, postId, commentId, newAttachment)
	if err != nil {
		return 0, err
	}

	return affectedRows, nil
}

func (s ServiceImpl) deleteById(currentUserId, postId, commentId int) (affectedRows int64, err error) {
	if currentUserId <= 0 {
		return 0, errors.New("currentUserId is required")
	}

	if postId <= 0 {
		return 0, errors.New("postId is required")
	}

	if commentId <= 0 {
		return 0, errors.New("commentId is required")
	}
	
	affectedRows, err = s.repository.deleteById(currentUserId, postId, commentId)
	if err != nil {
		return 0, err
	}

	return affectedRows, nil
}
