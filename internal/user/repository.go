package user

import (
	"github.com/jmoiron/sqlx"
)

type (
	Repository interface {
		save(firstName, lastName, email, password string) (id int64, err error)

		findById(id int) (User, error)
		findByEmail(email string) (User, error)

		findAll(isActive bool, limit, offset int) ([]User, error)

		deleteById(id int) (affectedRows int64, err error)

		changeStatus(userId int, isActive bool) (affectedRows int64, err error)
		changePassword(userId int, newPassword string) (affectedRows int64, err error)

		isEmailExists(email string) (bool, error)
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

func (r *RepositoryImpl) save(firstName, lastName, email, password string) (id int64, err error) {
	result, err := r.db.NamedExec(`INSERT INTO user (first_name, last_name, email, password) VALUES (:firstName, :lastName, :email, :password)`, map[string]any{
		"firstName": firstName,
		"lastName":  lastName,
		"email":     email,
		"password":  password,
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

func (r *RepositoryImpl) findById(id int) (User, error) {
	var user User

	err := r.db.Get(&user, "SELECT * FROM user WHERE id = ?", id)
	if err != nil {
		return User{}, err
	}

	return user, nil
}

func (r *RepositoryImpl) findByEmail(email string) (User, error) {
	var user User

	err := r.db.Get(&user, "SELECT * FROM user WHERE email = ?", email)
	if err != nil {
		return User{}, err
	}

	return user, nil
}

func (r *RepositoryImpl) findAll(isActive bool, limit, offset int) ([]User, error) {
	users := make([]User, limit)

	err := r.db.Select(&users, "SELECT * FROM user WHERE is_active = ? LIMIT ? OFFSET ?", isActive, limit, offset)
	if err != nil {
		return nil, err
	}

	return users, nil
}

func (r *RepositoryImpl) deleteById(id int) (affectedRows int64, err error) {
	result, err := r.db.NamedExec("DELETE FROM user WHERE id = :id", map[string]any{
		"id": id,
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

func (r *RepositoryImpl) changeStatus(userId int, isActive bool) (affectedRows int64, err error) {
	result, err := r.db.NamedExec("UPDATE user SET is_active = :isActive WHERE id = :id", map[string]any{
		"isActive": isActive,
		"id":       userId,
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

func (r *RepositoryImpl) changePassword(userId int, newPassword string) (affectedRows int64, err error) {
	result, err := r.db.NamedExec("UPDATE user SET password = :password WHERE id = :id", map[string]any{
		"password": newPassword,
		"id":       userId,
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

func (r *RepositoryImpl) isEmailExists(email string) (bool, error) {
	var exists bool
	err := r.db.Get(&exists, "SELECT EXISTS(SELECT 1 FROM user WHERE email = ?", email)
	if err != nil {
		return exists, err
	}

	return exists, nil
}
