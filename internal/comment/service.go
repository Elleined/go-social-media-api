package comment

import (
	"errors"
	"social-media-application/internal/paging"
	"strings"
)

type (
	Service interface {
		save(authorId, postId int, content, attachment string) (id int64, err error)

		getById(postId, commentId int) (Comment, error)
		getAll(postId int, isDeleted bool, request *paging.PageRequest) (*paging.Page[Comment], error)

		updateContent(currentUserId, postId, commentId int, newContent string) (affectedRows int64, err error)
		updateAttachment(currentUserId, postId, commentId int, newAttachment string) (affectedRows int64, err error)

		deleteById(currentUserId, postId, commentId int) (affectedRows int64, err error)
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

func (s ServiceImpl) save(authorId, postId int, content, attachment string) (id int64, err error) {
	if authorId <= 0 {
		return 0, errors.New("author is required")
	}

	if postId <= 0 {
		return 0, errors.New("post is required")
	}

	if strings.TrimSpace(content) == "" {
		return 0, errors.New("content is required")
	}

	id, err = s.repository.save(authorId, postId, content, attachment)
	if err != nil {
		return 0, err
	}

	return id, nil
}

func (s ServiceImpl) getById(postId, commentId int) (Comment, error) {
	comment, err := s.repository.findById(postId, commentId)
	if err != nil {
		return Comment{}, err
	}

	return comment, nil
}

func (s ServiceImpl) getAll(postId int, isDeleted bool, request *paging.PageRequest) (*paging.Page[Comment], error) {
	if postId <= 0 {
		return nil, errors.New("postId is required")
	}

	comments, err := s.repository.findAll(postId, isDeleted, request)
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

	if affectedRows <= 0 {
		return 0, errors.New("current user doesn't have this comment")
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

	if affectedRows <= 0 {
		return 0, errors.New("current user doesn't have this comment")
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

	if affectedRows <= 0 {
		return 0, errors.New("current user doesn't have this comment")
	}

	return affectedRows, nil
}
