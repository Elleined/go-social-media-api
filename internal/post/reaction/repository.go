package reaction

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"log"
	"social-media-application/internal/paging"
	"social-media-application/utils"
)

type (
	Repository interface {
		save(reactorId, postId, emojiId int) (id int64, err error)

		findById(postId, reactionId int) (Reaction, error)
		findAll(postId int, request *paging.PageRequest) (*paging.Page[Reaction], error)
		findAllByEmoji(postId int, emojiId int, request *paging.PageRequest) (*paging.Page[Reaction], error)

		update(reactorId, postId, newEmojiId int) (affectedRows int64, err error)

		delete(reactorId, postId int) (affectedRows int64, err error)

		isAlreadyReacted(reactorId, postId int) (bool, error)
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

func (repository RepositoryImpl) save(reactorId, postId, emojiId int) (id int64, err error) {
	result, err := repository.NamedExec("INSERT INTO post_reaction (reactor_id, post_id, emoji_id) VALUES (:reactorId, :postId, :emojiId)", map[string]any{
		"reactorId": reactorId,
		"postId":    postId,
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

func (repository RepositoryImpl) findById(postId, reactionId int) (Reaction, error) {
	var reaction Reaction
	err := repository.Get(&reaction, "SELECT * FROM post_reaction WHERE post_id = ? AND id = ?", postId, reactionId)
	if err != nil {
		return Reaction{}, err
	}

	return reaction, nil
}

func (repository RepositoryImpl) findAll(postId int, request *paging.PageRequest) (*paging.Page[Reaction], error) {
	if !utils.IsInDBTag(request.Field, Reaction{}) {
		request.Field = "created_at"
		log.Println("WARNING: field is not in database! defaulted to", request.Field)
	}

	if !utils.IsInSortingOrder(request.SortBy) {
		request.SortBy = "DESC"
		log.Println("WARNING: sortBy is not valid! defaulted to", request.SortBy)
	}

	var total int
	err := repository.Get(&total, "SELECT COUNT(*) FROM post_reaction WHERE post_id = ?", postId)
	if err != nil {
		return nil, err
	}

	reactions := make([]Reaction, 10)
	query := fmt.Sprintf("SELECT * FROM post_reaction WHERE post_id = ? ORDER BY %s %s LIMIT ? OFFSET ?", request.Field, request.SortBy)
	err = repository.Select(&reactions, query, postId, request.PageSize, request.Offset())
	if err != nil {
		return nil, err
	}

	return paging.NewPage(reactions, request, total), nil
}

func (repository RepositoryImpl) findAllByEmoji(postId int, emojiId int, request *paging.PageRequest) (*paging.Page[Reaction], error) {
	if !utils.IsInDBTag(request.Field, Reaction{}) {
		request.Field = "created_at"
		log.Println("WARNING: field is not in database! defaulted to", request.Field)
	}

	if !utils.IsInSortingOrder(request.SortBy) {
		request.SortBy = "DESC"
		log.Println("WARNING: sortBy is not valid! defaulted to", request.SortBy)
	}

	var total int
	err := repository.Get(&total, "SELECT COUNT(*) FROM post_reaction WHERE post_id = ? AND emoji_id = ?", postId, emojiId)

	reactions := make([]Reaction, 10)
	query := fmt.Sprintf("SELECT * FROM post_reaction WHERE post_id = ? AND emoji_id = ? ORDER BY %s %s LIMIT ? OFFSET ?", request.Field, request.SortBy)
	err = repository.Select(&reactions, query, postId, emojiId, request.PageSize, request.Offset())
	if err != nil {
		return nil, err
	}

	return paging.NewPage(reactions, request, total), nil
}

func (repository RepositoryImpl) update(reactorId, postId, newEmojiId int) (affectedRows int64, err error) {
	result, err := repository.NamedExec("UPDATE post_reaction SET emoji_id = :newEmojiId WHERE reactor_id = :reactorId AND post_id = :postId", map[string]any{
		"reactorId":  reactorId,
		"postId":     postId,
		"newEmojiId": newEmojiId,
	})

	if err != nil {
		return 0, err
	}

	affectedRow, err := result.RowsAffected()
	if err != nil {
		return 0, err
	}

	return affectedRow, nil
}

func (repository RepositoryImpl) delete(reactorId, postId int) (affectedRows int64, err error) {
	result, err := repository.NamedExec("DELETE FROM post_reaction WHERE reactor_id = :reactorId AND post_id = :postId", map[string]any{
		"reactorId": reactorId,
		"postId":    postId,
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

func (repository RepositoryImpl) isAlreadyReacted(reactorId, postId int) (bool, error) {
	var exists bool
	err := repository.Get(&exists, "SELECT EXISTS (SELECT 1 FROM post_reaction WHERE reactor_id = ? AND post_id = ?)", reactorId, postId)
	if err != nil {
		return exists, err
	}

	return exists, nil
}
