package repository

import (
	"context"

	"github.com/GabriellGds/go-orders/internal/models"
)

func (r repository) CreateUserRepository(ctx context.Context, user models.User) (int, error) {
	stmt, err := r.db.PreparexContext(ctx, `INSERT INTO users (name, email, password) VALUES($1, $2, $3) RETURNING id`)
	if err != nil {
		return 0, err
	}
	defer stmt.Close()

	var id int
	err = stmt.QueryRowContext(ctx, user.Name, user.Email, user.Password).Scan(&id)
	if err != nil {
		return 0, err
	}

	return id, nil
}

func (r repository) UpdateUserRepository(ctx context.Context, id int, user *models.User) error {
	stmt, err := r.db.PreparexContext(ctx, `UPDATE users SET name = $1 where id = $2 AND deleted_at IS NULL`)
	if err != nil {
		return err
	}
	defer stmt.Close()

	if _, err := stmt.ExecContext(ctx, user.Name, id); err != nil {
		return err
	}

	return nil
}

func (r repository) DeleteUserRepository(ctx context.Context, id int) error {
	stmt, err := r.db.PreparexContext(ctx, `UPDATE users SET deleted_at = NOW() WHERE id = $1 AND deleted_at IS NULL`)
	if err != nil {
		return err
	}
	defer stmt.Close()

	if _, err := stmt.ExecContext(ctx, id); err != nil {
		return err
	}

	return nil
}

func (r repository) FindUserRepository(ctx context.Context, id int) (models.User, error) {
	stmt, err := r.db.PreparexContext(ctx, `SELECT * FROM users WHERE id = $1 AND deleted_at IS NULL`)
	if err != nil {
		return models.User{}, err
	}
	defer stmt.Close()

	var user models.User
	if err := stmt.GetContext(ctx, &user, id); err != nil {
		return models.User{}, err
	}

	return user, nil
}

func (r repository) FindUserByEmailRepository(ctx context.Context, email string) (*models.User, error) {
	stmt, err := r.db.PreparexContext(ctx, "SELECT * FROM users WHERE email = $1 AND deleted_at IS NULL")
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	var user models.User
	if err := stmt.GetContext(ctx, &user, email); err != nil {
		return nil, err
	}

	return &user, nil
}

func (r repository) ListUserRepository(ctx context.Context) ([]models.User, error) {
	stmt, err := r.db.PreparexContext(ctx, `SELECT * FROM users WHERE deleted_at IS NULL`)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	var users []models.User
	if err := stmt.SelectContext(ctx, &users); err != nil {
		return nil, err
	}

	return users, nil
}
