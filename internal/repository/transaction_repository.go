package repository

import (
	"card_service/internal/database"
	"card_service/internal/models"
	"fmt"
	"time"
)

type TransactionRepository struct {
	db *database.DB
}

func NewTransactionRepository(db *database.DB) *TransactionRepository {
	return &TransactionRepository{db: db}
}

func (r *TransactionRepository) GetById(fromCardID, toCardID int) ([]*models.Transaction, error) {
	query := `
        SELECT id, from_card_id, to_card_id, amount, created_at
        FROM transactions
        WHERE 
            ($1 = 0 OR from_card_id = $1) AND
            ($2 = 0 OR to_card_id = $2)
        ORDER BY created_at DESC
    `
	rows, err := r.db.Query(query, fromCardID, toCardID)
	if err != nil {
		return nil, fmt.Errorf("failed to get transactions: %w", err)
	}
	defer rows.Close()

	var transactions []*models.Transaction
	for rows.Next() {
		tx := &models.Transaction{}
		err := rows.Scan(&tx.ID, &tx.FromCardID, &tx.ToCardID, &tx.Amount, &tx.CreatedAt)
		if err != nil {
			return nil, fmt.Errorf("failed to scan transaction: %w", err)
		}
		transactions = append(transactions, tx)
	}
	return transactions, nil
}

func (r *TransactionRepository) Create(tx *models.Transaction) (*models.Transaction, error) {
	query := `
        INSERT INTO transactions (from_card_id, to_card_id, amount, created_at)
        VALUES ($1, $2, $3, $4)
        RETURNING id, created_at
    `
	now := time.Now()
	err := r.db.QueryRow(query, tx.FromCardID, tx.ToCardID, tx.Amount, now).
		Scan(&tx.ID, &tx.CreatedAt)
	if err != nil {
		return nil, fmt.Errorf("failed to create transaction: %w", err)
	}
	return tx, nil
}
