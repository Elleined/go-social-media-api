package reaction

import "github.com/jmoiron/sqlx"

type Repository interface {
	save(reactorId, postId, emojiId int) (id int64, err error)

	findAll(postId int) ([]Reaction, error)
	findAllByEmoji(postId int, emojiId int) ([]Reaction, error)

	delete(reactorId, postId int) (affectedRows int64, err error)
}

type RepositoryImpl struct {
	db *sqlx.DB
}

func NewRepository(db *sqlx.DB) Repository {
	return &RepositoryImpl{
		db: db,
	}
}

func (r RepositoryImpl) save(reactorId, postId, emojiId int) (id int64, err error) {
	result, err := r.db.NamedExec("INSERT INTO post_reaction (reactor_id, post_id, emoji_id) VALUES (:reactor_id, :postId, :emojiId)", map[string]any{
		"reactor_id": reactorId,
		"post_id":    postId,
		"emoji_id":   emojiId,
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

func (r RepositoryImpl) findAll(postId int) ([]Reaction, error) {
	reactions := make([]Reaction, 0)

	err := r.db.Select(&reactions, "SELECT * FROM post_reaction WHERE post_id = ? ORDER BY created_at DESC", postId)
	if err != nil {
		return nil, err
	}

	return reactions, nil
}

func (r RepositoryImpl) findAllByEmoji(postId int, emojiId int) ([]Reaction, error) {
	reactions := make([]Reaction, 0)

	err := r.db.Select(&reactions, "SELECT * FROM post_reaction WHERE post_id = ? AND emoji_id = ? ORDER BY created_at DESC", postId, emojiId)
	if err != nil {
		return nil, err
	}
	
	return reactions, nil
}

func (r RepositoryImpl) delete(reactorId, postId int) (affectedRows int64, err error) {
	result, err := r.db.NamedExec("DELETE FROM post_reaction WHERE reactor_id = :reactor_id AND post_id = :post_id", map[string]any{
		"reactor_id": reactorId,
		"post_id":    postId,
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
