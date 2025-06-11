package commentreaction

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"social-media-application/internal/paging"
)

type (
	Repository interface {
		save(reactorId, postId, commentId, emojiId int) (id int64, err error)

		findById(postId, commentId, reactionId int) (Reaction, error)
		findAll(postId, commentId int, request *paging.PageRequest) (*paging.Page[Reaction], error)
		findAllByEmoji(postId, commentId, emojiId int, request *paging.PageRequest) (*paging.Page[Reaction], error)

		update(reactorId, postId, commentId, newEmojiId int) (affectedRows int64, err error)

		delete(reactorId, postId, commentId int) (affectedRows int64, err error)

		isAlreadyReacted(reactorId, postId, commentId int) (bool, error)
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

func (repository RepositoryImpl) save(reactorId, postId, commentId, emojiId int) (id int64, err error) {
	result, err := repository.NamedExec("INSERT INTO comment_reaction(reactor_id, comment_id, emoji_id) VALUES (:reactorId, :commentId, :emojiId)", map[string]any{
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

func (repository RepositoryImpl) findById(postId, commentId, reactionId int) (Reaction, error) {
	var reaction Reaction
	query := `
		SELECT cr.*
		FROM comment_reaction cr
		JOIN comment c ON c.id = cr.comment_id
		JOIN post p ON p.id = c.post_id
		WHERE p.id = ?
		AND cr.comment_id = ?
		AND cr.id = ?
	`
	err := repository.Get(&reaction, query, postId, commentId, reactionId)
	if err != nil {
		return Reaction{}, err
	}

	return reaction, nil
}

func (repository RepositoryImpl) findAll(postId, commentId int, request *paging.PageRequest) (*paging.Page[Reaction], error) {
	var total int
	query := `
		SELECT COUNT(*) 
		FROM comment_reaction cr
		JOIN comment c ON c.id = cr.comment_id
		JOIN post p ON p.id = c.post_id
		WHERE p.id = ?
		AND cr.comment_id = ?
	`
	err := repository.Get(&total, query, postId, commentId)
	if err != nil {
		return nil, err
	}

	reactions := make([]Reaction, request.PageSize)
	query = fmt.Sprintf(`
		SELECT cr.* 
		FROM comment_reaction cr
		JOIN comment c ON c.id = cr.comment_id
		JOIN post p ON p.id = c.post_id
		WHERE p.id = ?
		AND cr.comment_id = ?
		ORDER BY cr.%s %s
		LIMIT ?
		OFFSET ?
	`, request.Field, request.SortBy)
	err = repository.Select(&reactions, query, postId, commentId, request.PageSize, request.Offset())
	if err != nil {
		return nil, err
	}

	return paging.NewPage(reactions, request, total), nil
}

func (repository RepositoryImpl) findAllByEmoji(postId, commentId, emojiId int, request *paging.PageRequest) (*paging.Page[Reaction], error) {
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
	err := repository.Get(&total, query, postId, commentId, emojiId)
	if err != nil {
		return nil, err
	}

	reactions := make([]Reaction, request.PageSize)
	query = fmt.Sprintf(`
		SELECT cr.* 
		FROM comment_reaction cr
		JOIN comment c ON c.id = cr.comment_id
		JOIN post p ON p.id = c.post_id
		WHERE p.id = ?
		AND cr.comment_id = ?
		AND cr.emoji_id = ?
		ORDER BY cr.%s %s
		LIMIT ?
		OFFSET ?
	`, request.Field, request.SortBy)
	err = repository.Select(&reactions, query, postId, commentId, emojiId, request.PageSize, request.Offset())
	if err != nil {
		return nil, err
	}

	return paging.NewPage(reactions, request, total), nil
}

func (repository RepositoryImpl) update(reactorId, postId, commentId, newEmojiId int) (affectedRows int64, err error) {
	query := `
		UPDATE comment_reaction cr
		JOIN comment c ON c.id = cr.comment_id
		JOIN post p ON p.id = c.post_id
		SET cr.emoji_id = :newEmojiId
		WHERE p.id = :postId
		AND cr.comment_id = :commentId
		AND cr.reactor_id = :reactorId
	`
	result, err := repository.NamedExec(query, map[string]any{
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

func (repository RepositoryImpl) delete(reactorId, postId, commentId int) (affectedRows int64, err error) {
	query := `
		DELETE cr
		FROM comment_reaction cr
		JOIN comment c ON c.id = cr.comment_id
		JOIN post p ON p.id = c.post_id
		WHERE p.id = :postId
		AND cr.comment_id = :commentId
		AND cr.reactor_id = :reactorId
	`
	result, err := repository.NamedExec(query, map[string]any{
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

func (repository RepositoryImpl) isAlreadyReacted(reactorId, postId, commentId int) (bool, error) {
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
	err := repository.Get(&exists, query, postId, commentId, reactorId)
	if err != nil {
		return false, err
	}

	return exists, nil
}
