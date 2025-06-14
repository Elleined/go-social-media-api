package post

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"log"
	"social-media-application/internal/paging"
	"social-media-application/utils"
)

type (
	Repository interface {
		save(authorId int, content, attachment string) (id int64, err error)

		findById(postId int) (Post, error)
		findAll(currentUserId int, isDeleted bool, request *paging.PageRequest) (*paging.Page[Post], error)

		findAllBy(currentUserId int, isDeleted bool, request *paging.PageRequest) (*paging.Page[Post], error)

		updateContent(currentUserId, postId int, newContent string) (affectedRows int64, err error)
		updateAttachment(currentUserId, postId int, newAttachment string) (affectedRows int64, err error)

		deleteById(currentUserId, postId int) (affectedRows int64, err error)

		hasPost(currentUserId, postId int) (exists bool, err error)
	}

	RepositoryImpl struct {
		*sqlx.DB
	}
)

func NewRepository(db *sqlx.DB) Repository {
	return &RepositoryImpl{
		DB: db,
	}
}

func (repository RepositoryImpl) save(authorId int, content, attachment string) (id int64, err error) {
	result, err := repository.NamedExec("INSERT INTO post (content, attachment, author_id) VALUES (:content, :attachment, :authorId)", map[string]any{
		"content":    content,
		"attachment": attachment,
		"authorId":   authorId,
	})

	if err != nil {
		return 0, err
	}

	id, err = result.LastInsertId()
	if err != nil {
		return 0, err
	}

	return id, nil
}

func (repository RepositoryImpl) findById(postId int) (Post, error) {
	var post Post
	err := repository.Get(&post, "SELECT * FROM post WHERE id = ?", postId)
	if err != nil {
		return Post{}, err
	}

	return post, nil
}

func (repository RepositoryImpl) findAll(currentUserId int, isDeleted bool, request *paging.PageRequest) (*paging.Page[Post], error) {
	if !utils.IsInDBTag(request.Field, Post{}) {
		request.Field = "created_at"
		log.Println("WARNING: field is not in database! defaulted to", request.Field)
	}

	if !utils.IsInSortingOrder(request.SortBy) {
		request.SortBy = "DESC"
		log.Println("WARNING: sortBy is not valid! defaulted to", request.SortBy)
	}

	var total int
	err := repository.Get(&total, "SELECT COUNT(*) FROM post WHERE author_id != ? AND is_deleted = ?", currentUserId, isDeleted)
	if err != nil {
		return nil, err
	}

	posts := make([]Post, request.PageSize)
	query := fmt.Sprintf("SELECT * FROM post WHERE author_id != ? AND is_deleted = ? ORDER BY %s %s LIMIT ? OFFSET ?", request.Field, request.SortBy)
	err = repository.Select(&posts, query, currentUserId, isDeleted, request.PageSize, request.Offset())
	if err != nil {
		return nil, err
	}

	return paging.NewPage(posts, request, total), nil
}

func (repository RepositoryImpl) findAllBy(currentUserId int, isDeleted bool, request *paging.PageRequest) (*paging.Page[Post], error) {
	if !utils.IsInDBTag(request.Field, Post{}) {
		request.Field = "created_at"
		log.Println("WARNING: field is not in database! defaulted to", request.Field)
	}

	if !utils.IsInSortingOrder(request.SortBy) {
		request.SortBy = "DESC"
		log.Println("WARNING: sortBy is not valid! defaulted to", request.SortBy)
	}

	var total int
	err := repository.Get(&total, "SELECT COUNT(*) FROM post WHERE author_id = ? AND is_deleted = ?", currentUserId, isDeleted)
	if err != nil {
		return nil, err
	}

	posts := make([]Post, request.PageSize)
	query := fmt.Sprintf("SELECT * FROM post WHERE author_id = ? AND is_deleted = ? ORDER BY %s %s LIMIT ? OFFSET ?", request.Field, request.SortBy)
	err = repository.Select(&posts, query, currentUserId, isDeleted, request.PageSize, request.Offset())
	if err != nil {
		return nil, err
	}

	return paging.NewPage(posts, request, total), nil
}

func (repository RepositoryImpl) updateContent(currentUserId, postId int, newContent string) (affectedRows int64, err error) {
	result, err := repository.NamedExec("UPDATE post SET content = :content WHERE id = :id AND author_id = :authorId", map[string]any{
		"content":  newContent,
		"id":       postId,
		"authorId": currentUserId,
	})

	if err != nil {
		return 0, err
	}

	affectedRows, err = result.RowsAffected()
	if err != nil {
		return 0, err
	}

	return affectedRows, nil
}

func (repository RepositoryImpl) updateAttachment(currentUserId, postId int, newAttachment string) (affectedRows int64, err error) {
	result, err := repository.NamedExec("UPDATE post SET attachment = :attachment WHERE id = :postId AND author_id = :authorId", map[string]any{
		"attachment": newAttachment,
		"postId":     postId,
		"authorId":   currentUserId,
	})
	if err != nil {
		return 0, err
	}

	affectedRows, err = result.RowsAffected()
	if err != nil {
		return 0, err
	}

	return affectedRows, nil
}

func (repository RepositoryImpl) deleteById(currentUserId, postId int) (affectedRows int64, err error) {
	result, err := repository.NamedExec("UPDATE post SET is_deleted = true WHERE id = :postId AND author_id = :currentUserId", map[string]any{
		"postId":        postId,
		"currentUserId": currentUserId,
	})

	if err != nil {
		return 0, err
	}

	affectedRows, err = result.RowsAffected()
	if err != nil {
		return 0, err
	}

	return affectedRows, nil
}

func (repository RepositoryImpl) hasPost(currentUserId, postId int) (exists bool, err error) {
	err = repository.Get(&exists, "SELECT EXISTS(SELECT 1 FROM post WHERE author_id = ? AND id = ?)", currentUserId, postId)
	if err != nil {
		return exists, err
	}

	return exists, nil
}
