package models

import (
	"database/sql"
	"errors"
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

type CardOperator struct {
	ID   int    `db:"operator_id"`
	Name string `db:"operator_name"`
	Code string `db:"operator_code"`
}

type Card struct {
	ID         int       `db:"id"`
	UserID     int       `db:"user_id"`
	CardNumber string    `db:"card_number"`
	OperatorID int       `db:"operator_id"`
	IssueDate  time.Time `db:"issue_date"`
	ExpiryDate time.Time `db:"expiry_date"`
	IsActive   bool      `db:"is_active"`
	Balance    float64   `db:"balance"`
	CreatedAt  time.Time `db:"created_at"`

	// Operator for convenience mapping to proto Card
	Operator *CardOperator `db:"-"`
}

type Transaction struct {
	ID         int       `db:"id"`
	FromCardID int       `db:"from_card_id"`
	ToCardID   int       `db:"to_card_id"`
	Amount     float64   `db:"amount"`
	CreatedAt  time.Time `db:"created_at"`
}

func NewCard() *Card {
	return &Card{
		Operator: &CardOperator{},
	}
}

var ErrNotFound = errors.New("record not found")
