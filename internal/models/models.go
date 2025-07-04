package models

import (
	"database/sql"
	"time"
)

type User struct {
	ID        int            `db:"id"`
	FirstName string         `db:"first_name"`
	LastName  string         `db:"last_name"`
	Email     string         `db:"email"`
	Phone     sql.NullString `db:"phone"`
	BirthDate sql.NullTime   `db:"birth_date"`
	CreatedAt time.Time      `db:"created_at"`
}

type Card struct {
	ID                 int       `db:"id"`
	UserID             int       `db:"user_id"`
	CardNumber         string    `db:"card_number"`
	OperatorID         int       `db:"operator_id"`
	IssueDate          time.Time `db:"issue_date"`
	ExpiryDate         time.Time `db:"expiry_date"`
	IsActive           bool      `db:"is_active"`
	Balance            float64   `db:"balance"`
	CreatedAt          time.Time `db:"created_at"`
	OperatorName       string    `db:"operator_name"`
	OperatorCode       string    `db:"operator_code"`
	CardholderFullName string    `db:"cardholder_full_name"`
}

type CardOperator struct {
	ID        int       `db:"id"`
	Name      string    `db:"name"`
	Code      string    `db:"code"`
	CreatedAt time.Time `db:"created_at"`
}
