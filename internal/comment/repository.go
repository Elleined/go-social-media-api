package comment

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"log"
	"social-media-application/internal/paging"
	"social-media-application/utils"
)

type (
	Repository interface {
		save(authorId, postId int, content, attachment string) (id int64, err error)

		findById(postId, commentId int) (Comment, error)
		findAll(postId int, isDeleted bool, request *paging.PageRequest) (*paging.Page[Comment], error)

		updateContent(currentUserId, postId, commentId int, newContent string) (affectedRows int64, err error)
		updateAttachment(currentUserId, postId, commentId int, newAttachment string) (affectedRows int64, err error)

		deleteById(currentUserId, postId, commentId int) (affectedRows int64, err error)
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

func (repository RepositoryImpl) save(authorId, postId int, content, attachment string) (id int64, err error) {
	result, err := repository.NamedExec("INSERT INTO comment (author_id, post_id, content, attachment) VALUE (:authorId, :postId, :content, :attachment)", map[string]any{
		"authorId":   authorId,
		"postId":     postId,
		"content":    content,
		"attachment": attachment,
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

func (repository RepositoryImpl) findById(postId, commentId int) (Comment, error) {
	var comment Comment
	err := repository.Get(&comment, "SELECT * FROM comment WHERE post_id = ? AND id = ?", postId, commentId)
	if err != nil {
		return Comment{}, err
	}

	return comment, nil
}

func (repository RepositoryImpl) findAll(postId int, isDeleted bool, request *paging.PageRequest) (*paging.Page[Comment], error) {
	if !utils.IsInDBTag(request.Field, Comment{}) {
		request.Field = "created_at"
		log.Println("WARNING: field is not in database! defaulted to", request.Field)
	}

	if !utils.IsInSortingOrder(request.SortBy) {
		request.SortBy = "DESC"
		log.Println("WARNING: sortBy is not valid! defaulted to", request.SortBy)
	}

	var total int
	err := repository.Get(&total, "SELECT COUNT(*) FROM comment WHERE post_id = ? AND is_deleted = ?", postId, isDeleted)

	comments := make([]Comment, request.PageSize)
	query := fmt.Sprintf("SELECT * FROM comment WHERE post_id = ? AND is_deleted = ? ORDER BY %s %s LIMIT ? OFFSET ?", request.Field, request.SortBy)
	err = repository.Select(&comments, query, postId, isDeleted, request.PageSize, request.Offset())
	if err != nil {
		return nil, err
	}

	return paging.NewPage(comments, request, total), nil
}

func (repository RepositoryImpl) updateContent(currentUserId, postId, commentId int, newContent string) (affectedRows int64, err error) {
	result, err := repository.NamedExec("UPDATE comment SET content = :content WHERE id = :commentId AND author_id = :authorId AND post_id = :postId", map[string]any{
		"authorId":  currentUserId,
		"postId":    postId,
		"commentId": commentId,
		"content":   newContent,
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

func (repository RepositoryImpl) updateAttachment(currentUserId, postId, commentId int, newAttachment string) (affectedRows int64, err error) {
	result, err := repository.NamedExec("UPDATE comment SET attachment = :attachment WHERE id = :commentId AND author_id = :authorId AND post_id = :postId", map[string]any{
		"authorId":   currentUserId,
		"postId":     postId,
		"commentId":  commentId,
		"attachment": newAttachment,
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

func (repository RepositoryImpl) deleteById(currentUserId, postId, commentId int) (affectedRows int64, err error) {
	result, err := repository.NamedExec("UPDATE comment SET is_deleted = true WHERE id = :commentId AND author_id = :currentUserId AND post_id = :postId", map[string]any{
		"currentUserId": currentUserId,
		"postId":        postId,
		"commentId":     commentId,
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
