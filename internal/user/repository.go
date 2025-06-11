package user

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"log"
	"social-media-application/internal/paging"
	"social-media-application/utils"
)

type (
	Repository interface {
		saveLocal(firstName, lastName, email, password, attachment string) (id int64, err error)
		saveSocial(firstName, lastName, email string) (id int64, err error)

		findById(id int) (User, error)
		findByEmail(email string) (User, error)

		findAll(isActive bool, request *paging.PageRequest) (*paging.Page[User], error)

		deleteById(id int) (affectedRows int64, err error)

		changeAttachment(userId int, attachment string) (affectedRows int64, err error)
		changeStatus(userId int, isActive bool) (affectedRows int64, err error)
		changePassword(userId int, newPassword string) (affectedRows int64, err error)

		isEmailExists(email string) (bool, error)
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

func (repository *RepositoryImpl) saveLocal(firstName, lastName, email, password, attachment string) (id int64, err error) {
	result, err := repository.NamedExec(`INSERT INTO user (first_name, last_name, email, password, attachment) VALUES (:firstName, :lastName, :email, :password, :attachment)`, map[string]any{
		"firstName":  firstName,
		"lastName":   lastName,
		"email":      email,
		"password":   password,
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

func (repository *RepositoryImpl) saveSocial(firstName, lastName, email string) (id int64, err error) {
	result, err := repository.NamedExec(`INSERT INTO user (first_name, last_name, email) VALUES (:firstName, :lastName, :email)`, map[string]any{
		"firstName": firstName,
		"lastName":  lastName,
		"email":     email,
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

func (repository *RepositoryImpl) findById(id int) (User, error) {
	var user User

	err := repository.Get(&user, "SELECT * FROM user WHERE id = ?", id)
	if err != nil {
		return user, err
	}

	return user, nil
}

func (repository *RepositoryImpl) findByEmail(email string) (User, error) {
	var user User

	err := repository.Get(&user, "SELECT * FROM user WHERE email = ?", email)
	if err != nil {
		return User{}, err
	}

	return user, nil
}

func (repository *RepositoryImpl) findAll(isActive bool, request *paging.PageRequest) (*paging.Page[User], error) {
	if !utils.IsInDBTag(request.Field, User{}) {
		request.Field = "created_at"
		log.Println("WARNING: field is not in database! defaulted to", request.Field)
	}

	if !utils.IsInSortingOrder(request.SortBy) {
		request.SortBy = "DESC"
		log.Println("WARNING: sortBy is not valid! defaulted to", request.SortBy)
	}

	var total int
	err := repository.Get(&total, "SELECT COUNT(*) FROM user WHERE is_active = ?", isActive)
	if err != nil {
		return nil, err
	}

	users := make([]User, request.PageSize)
	query := fmt.Sprintf("SELECT * FROM user WHERE is_active = ? ORDER BY %s %s LIMIT ? OFFSET ?", request.Field, request.SortBy)
	err = repository.Select(&users, query, isActive, request.PageSize, request.Offset())
	if err != nil {
		return nil, err
	}

	return paging.NewPage(users, request, total), nil
}

func (repository *RepositoryImpl) deleteById(id int) (affectedRows int64, err error) {
	result, err := repository.NamedExec("DELETE FROM user WHERE id = :id", map[string]any{
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

func (repository *RepositoryImpl) changeAttachment(userId int, attachment string) (affectedRows int64, err error) {
	result, err := repository.NamedExec("UPDATE user SET attachment = :attachment WHERE id = :userId", map[string]any{
		"userId":     userId,
		"attachment": attachment,
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

func (repository *RepositoryImpl) changeStatus(userId int, isActive bool) (affectedRows int64, err error) {
	result, err := repository.NamedExec("UPDATE user SET is_active = :isActive WHERE id = :userId", map[string]any{
		"isActive": isActive,
		"userId":   userId,
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

func (repository *RepositoryImpl) changePassword(userId int, newPassword string) (affectedRows int64, err error) {
	result, err := repository.NamedExec("UPDATE user SET password = :password WHERE id = :userId", map[string]any{
		"password": newPassword,
		"userId":   userId,
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

func (repository *RepositoryImpl) isEmailExists(email string) (bool, error) {
	var exists bool
	err := repository.Get(&exists, "SELECT EXISTS(SELECT 1 FROM user WHERE email = ?)", email)
	if err != nil {
		return exists, err
	}

	return exists, nil
}
