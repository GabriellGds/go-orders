package repository

import (
	"github.com/GabriellGds/go-orders/internal/models"
	"github.com/jmoiron/sqlx"
)

type userRepository struct {
	db *sqlx.DB
}

func NewUserRepository(db *sqlx.DB) UserRepositoryInterface {
	return &userRepository{db: db}
}

func (ur userRepository) CreateUser(user models.User) (models.User, error) {
	stmt, err := ur.db.Preparex(`INSERT INTO users (name, email, password) VALUES($1, $2, $3) RETURNING id`)
	if err != nil {
		return models.User{}, err
	}
	defer stmt.Close()

	var lastInsertID int
	err = stmt.QueryRow(user.Name, user.Email, user.Password).Scan(&lastInsertID)
	if err != nil {
		return models.User{}, err
	}

	user.ID = lastInsertID

	return user, nil
}

func (ur userRepository) UpdateUser(id int, user models.UserUpdateRequest) error {
	stmt, err := ur.db.Preparex(`UPDATE users SET name = $1 where id = $2 AND deleted_at IS NULL`)
	if err != nil {
		return err
	}
	defer stmt.Close()

	if _, err := stmt.Exec(user.Name, id); err != nil {
		return err
	}

	return nil
}

func (ur userRepository) DeleteUser(id int) error {
	stmt, err := ur.db.Preparex(`UPDATE users SET deleted_at = NOW() WHERE id = $1 AND deleted_at IS NULL`)
	if err != nil {
		return err
	}
	defer stmt.Close()

	if _, err := stmt.Exec(id); err != nil {
		return err
	}

	return nil
}

func (ur userRepository) User(id int) (models.User, error) {
	stmt, err := ur.db.Preparex(`SELECT * FROM users WHERE id = $1 AND deleted_at IS NULL`)
	if err != nil {
		return models.User{}, err
	}
	defer stmt.Close()

	var user models.User
	if err := stmt.Get(&user, id); err != nil {
		return models.User{}, err
	}

	return user, nil
}

func (ur userRepository) FindEmail(email string) (*models.User, error) {
	stmt, err := ur.db.Preparex("SELECT * FROM users WHERE email = $1 AND deleted_at IS NULL")
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	var user models.User
	if err := stmt.Get(&user, email); err != nil {
		return nil, err
	}

	return &user, nil
}

