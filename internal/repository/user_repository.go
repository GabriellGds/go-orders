package repository

import (
	"github.com/GabriellGds/go-orders/internal/models"
)

func (r repository) CreateUserRepository(user models.User) (int, error) {
	stmt, err := r.db.Preparex(`INSERT INTO users (name, email, password) VALUES($1, $2, $3) RETURNING id`)
	if err != nil {
		return 0, err
	}
	defer stmt.Close()

	var id int
	err = stmt.QueryRow(user.Name, user.Email, user.Password).Scan(&id)
	if err != nil {
		return 0, err
	}

	return id, nil
}

func (r repository) UpdateUserRepository(id int, user *models.User) error {
	stmt, err := r.db.Preparex(`UPDATE users SET name = $1 where id = $2 AND deleted_at IS NULL`)
	if err != nil {
		return err
	}
	defer stmt.Close()

	if _, err := stmt.Exec(user.Name, id); err != nil {
		return err
	}

	return nil
}

func (r repository) DeleteUserRepository(id int) error {
	stmt, err := r.db.Preparex(`UPDATE users SET deleted_at = NOW() WHERE id = $1 AND deleted_at IS NULL`)
	if err != nil {
		return err
	}
	defer stmt.Close()

	if _, err := stmt.Exec(id); err != nil {
		return err
	}

	return nil
}

func (r repository) FindUserRepository(id int) (models.User, error) {
	stmt, err := r.db.Preparex(`SELECT * FROM users WHERE id = $1 AND deleted_at IS NULL`)
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

func (r repository) FindUserByEmailRepository(email string) (*models.User, error) {
	stmt, err := r.db.Preparex("SELECT * FROM users WHERE email = $1 AND deleted_at IS NULL")
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

