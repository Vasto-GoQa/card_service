package repository

import (
	"card_service/internal/database"
	"card_service/internal/models"
	"database/sql"
	"errors"
	"fmt"
)

type UserRepository struct {
	db *database.DB
}

func NewUserRepository(db *database.DB) *UserRepository {
	return &UserRepository{db: db}
}

func (r *UserRepository) Create(user *models.User) (*models.User, error) {
	query := `INSERT INTO users (first_name, last_name, email, phone, birth_date)
				VALUES ($1, $2, $3, $4, $5)
				RETURNING id, created_at`

	err := r.db.QueryRow(query, user.FirstName, user.LastName, user.Email, user.Phone, user.BirthDate).
		Scan(&user.ID, &user.CreatedAt)
	if err != nil {
		return nil, fmt.Errorf("failed to create user: %w", err)
	}

	return user, nil
}

func (r *UserRepository) GetByID(id int) (*models.User, error) {
	query := `SELECT id, first_name, last_name, email, phone, birth_date, created_at
		FROM users WHERE id = $1`

	user := &models.User{}
	err := r.db.QueryRow(query, id).Scan(
		&user.ID, &user.FirstName, &user.LastName, &user.Email,
		&user.Phone, &user.BirthDate, &user.CreatedAt)
	if errors.Is(err, sql.ErrNoRows) {
		return nil, fmt.Errorf("user with id %d not found", id)
	}
	if err != nil {
		return nil, fmt.Errorf("failed to get user: %w", err)
	}

	return user, nil
}

func (r *UserRepository) GetAll() ([]*models.User, error) {
	query := `SELECT id, first_name, last_name, email, phone, birth_date, created_at
		FROM users ORDER BY id`

	rows, err := r.db.Query(query)
	if err != nil {
		return nil, fmt.Errorf("failed to get users: %w", err)
	}
	defer rows.Close()

	var users []*models.User
	for rows.Next() {
		user := &models.User{}
		err := rows.Scan(&user.ID, &user.FirstName, &user.LastName, &user.Email,
			&user.Phone, &user.BirthDate, &user.CreatedAt)
		if err != nil {
			return nil, fmt.Errorf("failed to scan user: %w", err)
		}
		users = append(users, user)
	}

	return users, nil
}

func (r *UserRepository) Update(user *models.User) (*models.User, error) {
	query := `UPDATE users 
		SET first_name = $2, last_name = $3, email = $4, phone = $5, birth_date = $6
		WHERE id = $1
		RETURNING id, first_name, last_name, email, phone, birth_date, created_at`

	err := r.db.QueryRow(query, user.ID, user.FirstName, user.LastName,
		user.Email, user.Phone, user.BirthDate).Scan(
		&user.ID, &user.FirstName, &user.LastName, &user.Email,
		&user.Phone, &user.BirthDate, &user.CreatedAt)
	if errors.Is(err, sql.ErrNoRows) {
		return nil, fmt.Errorf("user with id %d not found", user.ID)
	}
	if err != nil {
		return nil, fmt.Errorf("failed to update user: %w", err)
	}

	return user, nil
}

func (r *UserRepository) Delete(id int) error {
	query := `DELETE FROM users WHERE id = $1`

	result, err := r.db.Exec(query, id)
	if err != nil {
		return fmt.Errorf("failed to delete user: %w", err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("failed to check delete result: %w", err)
	}

	if rowsAffected == 0 {
		return fmt.Errorf("user with id %d not found", id)
	}

	return nil
}
