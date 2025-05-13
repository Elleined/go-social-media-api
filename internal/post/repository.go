package post

import (
	"github.com/jmoiron/sqlx"
)

type (
	Repository interface {
		save(authorId int, subject, content string) (id int64, err error)

		findAll(currentUserId int, isDeleted bool, limit, offset int) ([]Post, error)
		findAllBy(currentUserId int, isDeleted bool, limit, offset int) ([]Post, error)

		updateSubject(currentUserId, postId int, newSubject string) (affectedRows int64, err error)
		updateContent(currentUserId, postId int, newContent string) (affectedRows int64, err error)
		updateAttachment(currentUserId, postId int, newAttachment string) (affectedRows int64, err error)

		deleteById(currentUserId, postId int) (affectedRows int64, err error)

		hasPost(currentUserId, postId int) (exists bool, err error)
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

func (r RepositoryImpl) save(authorId int, subject, content string) (id int64, err error) {
	result, err := r.db.NamedExec("INSERT INTO post (subject, content, author_id) VALUES (:subject, :content, :authorId)", map[string]any{
		"subject":  subject,
		"content":  content,
		"authorId": authorId,
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

func (r RepositoryImpl) findAll(currentUserId int, isDeleted bool, limit, offset int) ([]Post, error) {
	posts := make([]Post, limit)

	err := r.db.Select(&posts, "SELECT * FROM post WHERE author_id != ? AND is_deleted = ? ORDER BY created_at DESC LIMIT ? OFFSET ?", currentUserId, isDeleted, limit, offset)
	if err != nil {
		return nil, err
	}

	return posts, nil
}

func (r RepositoryImpl) findAllBy(currentUserId int, isDeleted bool, limit, offset int) ([]Post, error) {
	posts := make([]Post, limit)

	err := r.db.Select(&posts, "SELECT * FROM post WHERE author_id = ? AND is_deleted = ? ORDER BY created_at DESC LIMIT ? OFFSET ?", currentUserId, isDeleted, limit, offset)
	if err != nil {
		return nil, err
	}

	return posts, nil
}

func (r RepositoryImpl) updateSubject(currentUserId int, postId int, newSubject string) (affectedRows int64, err error) {
	result, err := r.db.NamedExec("UPDATE post SET subject = :subject WHERE id = :id AND author_id = :authorId", map[string]any{
		"subject":  newSubject,
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

func (r RepositoryImpl) updateContent(currentUserId, postId int, newContent string) (affectedRows int64, err error) {
	result, err := r.db.NamedExec("UPDATE post SET content = :content WHERE id = :id AND author_id = :authorId", map[string]any{
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

func (r RepositoryImpl) updateAttachment(currentUserId, postId int, newAttachment string) (affectedRows int64, err error) {
	result, err := r.db.NamedExec("UPDATE post SET attachment = :attachment WHERE id = :postId AND author_id = :authorId", map[string]any{
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

func (r RepositoryImpl) deleteById(currentUserId, postId int) (affectedRows int64, err error) {
	result, err := r.db.NamedExec("UPDATE post SET is_deleted = true WHERE id = :postId AND author_id = :currentUserId", map[string]any{
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

func (r RepositoryImpl) hasPost(currentUserId, postId int) (exists bool, err error) {
	err = r.db.Get(&exists, "SELECT EXISTS(SELECT 1 FROM post WHERE author_id = ? AND post_id = ?)", currentUserId, postId)
	if err != nil {
		return exists, err
	}

	return exists, nil
}
