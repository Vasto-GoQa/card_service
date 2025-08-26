package repository

import (
	"card_service/internal/database"
	"card_service/internal/models"
	"database/sql"
	"fmt"
)

type CardRepository struct {
	db *database.DB
}

func NewCardRepository(db *database.DB) *CardRepository {
	return &CardRepository{db: db}
}

func (r *CardRepository) GetByID(id int) (*models.Card, error) {
	query := `SELECT 
			bc.id, bc.user_id, bc.card_number, bc.operator_id, 
			bc.issue_date, bc.expiry_date, bc.is_active, bc.balance, bc.created_at,
			co.name as operator_name, co.code as operator_code, co.id as operator_id
		FROM bank_cards bc
		JOIN card_operators co ON bc.operator_id = co.id
		WHERE bc.id = $1`

	card := models.NewCard()
	err := r.db.QueryRow(query, id).Scan(
		&card.ID, &card.UserID, &card.CardNumber, &card.OperatorID,
		&card.IssueDate, &card.ExpiryDate, &card.IsActive, &card.Balance, &card.CreatedAt,
		&card.Operator.Name, &card.Operator.Code, &card.Operator.ID)
	if err == sql.ErrNoRows {
		return nil, models.ErrNotFound
	}
	if err != nil {
		return nil, fmt.Errorf("failed to get card: %w", err)
	}

	return card, nil
}

func (r *CardRepository) GetOperatorByID(id int) (*models.CardOperator, error) {
	query := `SELECT co.id, co.name, co.code
		FROM card_operators co
		WHERE co.id = $1`

	operator := &models.CardOperator{}
	err := r.db.QueryRow(query, id).Scan(&operator.ID, &operator.Name, &operator.Code)
	if err == sql.ErrNoRows {
		return nil, models.ErrNotFound
	}
	if err != nil {
		return nil, fmt.Errorf("failed to get operator: %w", err)
	}

	return operator, nil
}

func (r *CardRepository) Create(card *models.Card) (*models.Card, error) {
	query := `INSERT INTO bank_cards (user_id, card_number, operator_id, issue_date, expiry_date, balance)
		VALUES ($1, $2, $3, $4, $5, $6)
		RETURNING id, created_at`

	err := r.db.QueryRow(query, card.UserID, card.CardNumber, card.OperatorID,
		card.IssueDate, card.ExpiryDate, card.Balance).
		Scan(&card.ID, &card.CreatedAt)
	if err != nil {
		return nil, fmt.Errorf("failed to create card: %w", err)
	}

	// Return the created card with all fields populated
	return r.GetByID(card.ID)
}

func (r *CardRepository) Delete(id int) error {
	query := `DELETE FROM bank_cards WHERE id = $1`

	result, err := r.db.Exec(query, id)
	if err != nil {
		return fmt.Errorf("failed to delete card: %w", err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("failed to check delete result: %w", err)
	}

	if rowsAffected == 0 {
		return models.ErrNotFound
	}

	return nil
}
