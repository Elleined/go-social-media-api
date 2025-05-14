package comment

import (
	"github.com/jmoiron/sqlx"
	"social-media-application/internal/paging"
)

type (
	Repository interface {
		save(authorId, postId int, content string) (id int64, err error)

		findAll(postId int, isDeleted bool, request *paging.PageRequest) (*paging.Page[Comment], error)

		updateContent(currentUserId, postId, commentId int, newContent string) (affectedRows int64, err error)
		updateAttachment(currentUserId, postId, commentId int, newAttachment string) (affectedRows int64, err error)

		deleteById(currentUserId, postId, commentId int) (affectedRows int64, err error)
	}

	RepositoryImpl struct {
		db *sqlx.DB
	}
)

func NewRepository(db *sqlx.DB) Repository {
	return &RepositoryImpl{
		db: db,
	}
}

func (r RepositoryImpl) save(authorId, postId int, content string) (id int64, err error) {
	result, err := r.db.NamedExec("INSERT INTO comment (author_id, post_id, content) VALUE (:authorId, :postId, :content)", map[string]any{
		"authorId": authorId,
		"postId":   postId,
		"content":  content,
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

func (r RepositoryImpl) findAll(postId int, isDeleted bool, request *paging.PageRequest) (*paging.Page[Comment], error) {
	var total int
	err := r.db.Get(&total, "SELECT COUNT(*) FROM comment WHERE post_id = ? AND is_deleted = ? ORDER BY created_at DESC", postId, isDeleted)

	comments := make([]Comment, request.PageSize)
	err = r.db.Select(&comments, "SELECT * FROM comment WHERE post_id = ? AND is_deleted = ? ORDER BY created_at DESC LIMIT ? OFFSET ?", postId, isDeleted, request.PageSize, request.Offset())
	if err != nil {
		return nil, err
	}

	return paging.NewPage(comments, request, total), nil
}

func (r RepositoryImpl) updateContent(currentUserId, postId, commentId int, newContent string) (affectedRows int64, err error) {
	result, err := r.db.NamedExec("UPDATE comment SET content = :content WHERE id = :commentId AND author_id = :authorId AND post_id = :postId", map[string]any{
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

func (r RepositoryImpl) updateAttachment(currentUserId, postId, commentId int, newAttachment string) (affectedRows int64, err error) {
	result, err := r.db.NamedExec("UPDATE comment SET attachment = :attachment WHERE id = :commentId AND author_id = :authorId AND post_id = :postId", map[string]any{
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

func (r RepositoryImpl) deleteById(currentUserId, postId, commentId int) (affectedRows int64, err error) {
	result, err := r.db.NamedExec("UPDATE comment SET is_deleted = true WHERE id = :commentId AND author_id = :currentUserId AND post_id = :postId", map[string]any{
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
