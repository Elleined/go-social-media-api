package post

import (
	"github.com/jmoiron/sqlx"
)

type Repository interface {
	save(authorId int, subject, content string) (id int64, err error)

	getAll(currentUserId, limit, offset int) ([]Post, error)
	getAllBy(currentUserId, limit, offset int) ([]Post, error)

	updateSubject(currentUserId, postId int, newSubject string) (affectedRows int64, err error)
	updateContent(currentUserId, postId int, newContent string) (affectedRows int64, err error)
	updateAttachment(currentUserId, postId int, newAttachment string) (affectedRows int64, err error)

	deleteById(currentUserId, postId int) (affectedRows int64, err error)
}

type RepositoryImpl struct {
	db *sqlx.DB
}

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

func (r RepositoryImpl) getAll(currentUserId, limit, offset int) ([]Post, error) {
	posts := make([]Post, offset)

	err := r.db.Select(&posts, "SELECT * FROM post WHERE author_id != ? ORDER BY created_at DESC LIMIT ? OFFSET ?", currentUserId, limit, offset)
	if err != nil {
		return nil, err
	}

	return posts, nil
}

func (r RepositoryImpl) getAllBy(currentUserId, limit, offset int) ([]Post, error) {
	posts := make([]Post, offset)

	err := r.db.Select(&posts, "SELECT * FROM post WHERE author_id = ? ORDER BY created_at DESC LIMIT ? OFFSET ?", currentUserId, limit, offset)
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
	result, err := r.db.NamedExec("UPDATE post SET attachment = :attachment WHERE id = :id AND author_id = :authorId", map[string]any{
		"attachment": newAttachment,
		"id":         postId,
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
	result, err := r.db.NamedExec("DELETE FROM post WHERE id = :postId AND author_id = :authorId", map[string]any{
		"postId":   postId,
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
