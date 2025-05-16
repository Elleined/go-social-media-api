package commentreaction

import (
	"github.com/jmoiron/sqlx"
	"social-media-application/internal/paging"
)

type (
	Repository interface {
		save(reactorId, postId, commentId, emojiId int) (id int64, err error)

		findAll(postId, commentId int, request *paging.PageRequest) (*paging.Page[Reaction], error)
		findAllByEmoji(postId, commentId, emojiId int, request *paging.PageRequest) (*paging.Page[Reaction], error)

		update(reactorId, postId, commentId, newEmojiId int) (affectedRows int64, err error)

		delete(reactorId, postId, commentId int) (affectedRows int64, err error)

		isAlreadyReacted(reactorId, postId, commentId int) (bool, error)
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

func (r RepositoryImpl) save(reactorId, postId, commentId, emojiId int) (id int64, err error) {
	result, err := r.db.NamedExec("INSERT INTO comment_reaction(reactor_id, comment_id, emoji_id) VALUES (:reactorId, :commentId, :emojiId)", map[string]any{
		"reactorId": reactorId,
		"commentId": commentId,
		"emojiId":   emojiId,
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

func (r RepositoryImpl) findAll(postId, commentId int, pageRequest *paging.PageRequest) (*paging.Page[Reaction], error) {
	var total int
	query := `
		SELECT COUNT(*) 
		FROM comment_reaction cr
		JOIN comment c ON c.id = cr.comment_id
		JOIN post p ON p.id = c.post_id
		WHERE p.id = ?
		AND cr.comment_id = ?
	`
	err := r.db.Get(&total, query, postId, commentId)
	if err != nil {
		return nil, err
	}

	reactions := make([]Reaction, pageRequest.PageSize)
	query = `
		SELECT cr.* 
		FROM comment_reaction cr
		JOIN comment c ON c.id = cr.comment_id
		JOIN post p ON p.id = c.post_id
		WHERE p.id = ?
		AND cr.comment_id = ?
		ORDER BY cr.created_at DESC
		LIMIT ?
		OFFSET ?
	`
	err = r.db.Select(&reactions, query, postId, commentId, pageRequest.PageSize, pageRequest.Offset())
	if err != nil {
		return nil, err
	}

	return paging.NewPage(reactions, pageRequest, total), nil
}

func (r RepositoryImpl) findAllByEmoji(postId, commentId, emojiId int, pageRequest *paging.PageRequest) (*paging.Page[Reaction], error) {
	var total int
	query := `
		SELECT COUNT(*) 
		FROM comment_reaction cr
		JOIN comment c ON c.id = cr.comment_id
		JOIN post p ON p.id = c.post_id
		WHERE p.id = ?
		AND cr.comment_id = ?
		AND cr.emoji_id = ?
	`
	err := r.db.Get(&total, query, postId, commentId, emojiId)
	if err != nil {
		return nil, err
	}

	reactions := make([]Reaction, pageRequest.PageSize)
	query = `
		SELECT cr.* 
		FROM comment_reaction cr
		JOIN comment c ON c.id = cr.comment_id
		JOIN post p ON p.id = c.post_id
		WHERE p.id = ?
		AND cr.comment_id = ?
		AND cr.emoji_id = ?
		ORDER BY cr.created_at DESC
		LIMIT ?
		OFFSET ?
	`
	err = r.db.Select(&reactions, query, postId, commentId, emojiId, pageRequest.PageSize, pageRequest.Offset())
	if err != nil {
		return nil, err
	}

	return paging.NewPage(reactions, pageRequest, total), nil
}

func (r RepositoryImpl) update(reactorId, postId, commentId, newEmojiId int) (affectedRows int64, err error) {
	query := `
		UPDATE comment_reaction cr
		JOIN comment c ON c.id = cr.comment_id
		JOIN post p ON p.id = c.post_id
		SET cr.emoji_id = :newEmojiId
		WHERE p.id = :postId
		AND cr.comment_id = :commentId
		AND cr.reactor_id = :reactorId
	`
	result, err := r.db.NamedExec(query, map[string]any{
		"reactorId":  reactorId,
		"postId":     postId,
		"commentId":  commentId,
		"newEmojiId": newEmojiId,
	})
	if err != nil {
		return 0, err
	}

	affectedRows, err = result.RowsAffected()
	if err != nil {
		return 0, err
	}

	return affectedRows, err
}

func (r RepositoryImpl) delete(reactorId, postId, commentId int) (affectedRows int64, err error) {
	query := `
		DELETE cr
		FROM comment_reaction cr
		JOIN comment c ON c.id = cr.comment_id
		JOIN post p ON p.id = c.post_id
		WHERE p.id = :postId
		AND cr.comment_id = :commentId
		AND cr.reactor_id = :reactorId
	`
	result, err := r.db.NamedExec(query, map[string]any{
		"reactorId": reactorId,
		"postId":    postId,
		"commentId": commentId,
	})
	if err != nil {
		return 0, err
	}

	affectedRows, err = result.RowsAffected()
	if err != nil {
		return 0, err
	}

	return affectedRows, err
}

func (r RepositoryImpl) isAlreadyReacted(reactorId, postId, commentId int) (bool, error) {
	var exists bool
	query := `
		SELECT EXISTS (
			SELECT 1 FROM comment_reaction cr
			JOIN comment c ON c.id = cr.comment_id
			JOIN post p ON p.id = c.post_id
			WHERE p.id = ?
			AND cr.comment_id = ?
			AND cr.reactor_id = ?
		)
	`
	err := r.db.Get(&exists, query, postId, commentId, reactorId)
	if err != nil {
		return false, err
	}

	return exists, nil
}
