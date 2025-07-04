package repository

import (
	"card_service/internal/database"
	"card_service/internal/models"
	"database/sql"
	"fmt"
	"math/rand"
	"time"
)

type CardRepository struct {
	db *database.DB
}

func NewCardRepository(db *database.DB) *CardRepository {
	return &CardRepository{db: db}
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

	// Получаем полную информацию о карте
	return r.GetByID(card.ID)
}

func (r *CardRepository) GetByID(id int) (*models.Card, error) {
	query := `SELECT 
			bc.id, bc.user_id, bc.card_number, bc.operator_id, 
			bc.issue_date, bc.expiry_date, bc.is_active, bc.balance, bc.created_at,
			co.name as operator_name, co.code as operator_code,
			u.first_name || ' ' || u.last_name as cardholder_full_name
		FROM bank_cards bc
		JOIN card_operators co ON bc.operator_id = co.id
		JOIN users u ON bc.user_id = u.id
		WHERE bc.id = $1`

	card := &models.Card{}
	err := r.db.QueryRow(query, id).Scan(
		&card.ID, &card.UserID, &card.CardNumber, &card.OperatorID,
		&card.IssueDate, &card.ExpiryDate, &card.IsActive, &card.Balance, &card.CreatedAt,
		&card.OperatorName, &card.OperatorCode, &card.CardholderFullName)
	if err == sql.ErrNoRows {
		return nil, fmt.Errorf("card with id %d not found", id)
	}
	if err != nil {
		return nil, fmt.Errorf("failed to get card: %w", err)
	}

	return card, nil
}

func (r *CardRepository) GetAll() ([]*models.Card, error) {
	query := `SELECT 
			bc.id, bc.user_id, bc.card_number, bc.operator_id, 
			bc.issue_date, bc.expiry_date, bc.is_active, bc.balance, bc.created_at,
			co.name as operator_name, co.code as operator_code,
			u.first_name || ' ' || u.last_name as cardholder_full_name
		FROM bank_cards bc
		JOIN card_operators co ON bc.operator_id = co.id
		JOIN users u ON bc.user_id = u.id
		ORDER BY bc.id`

	rows, err := r.db.Query(query)
	if err != nil {
		return nil, fmt.Errorf("failed to get cards: %w", err)
	}
	defer rows.Close()

	var cards []*models.Card
	for rows.Next() {
		card := &models.Card{}
		err := rows.Scan(
			&card.ID, &card.UserID, &card.CardNumber, &card.OperatorID,
			&card.IssueDate, &card.ExpiryDate, &card.IsActive, &card.Balance, &card.CreatedAt,
			&card.OperatorName, &card.OperatorCode, &card.CardholderFullName)
		if err != nil {
			return nil, fmt.Errorf("failed to scan card: %w", err)
		}
		cards = append(cards, card)
	}

	return cards, nil
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
		return fmt.Errorf("card with id %d not found", id)
	}

	return nil
}

func (r *CardRepository) GenerateCard(userID, operatorID int, balance float64) (*models.Card, error) {
	// Генерируем номер карты
	cardNumber := r.generateCardNumber(operatorID)

	// Проверяем уникальность номера
	for r.cardNumberExists(cardNumber) {
		cardNumber = r.generateCardNumber(operatorID)
	}

	// Устанавливаем даты
	issueDate := time.Now()
	expiryDate := issueDate.AddDate(5, 0, 0) // +5 лет

	card := &models.Card{
		UserID:     userID,
		CardNumber: cardNumber,
		OperatorID: operatorID,
		IssueDate:  issueDate,
		ExpiryDate: expiryDate,
		Balance:    balance,
	}

	return r.Create(card)
}

func (r *CardRepository) generateCardNumber(operatorID int) string {
	var prefix string
	switch operatorID {
	case 1: // Visa
		prefix = "4"
	case 2: // Mastercard
		prefix = "5"
	case 3: // American Express
		prefix = "3"
	case 4: // Maestro
		prefix = "6"
	case 5: // МИР
		prefix = "2"
	default:
		prefix = "4"
	}

	// Генерируем остальные 15 цифр
	cardNumber := prefix
	for i := 0; i < 15; i++ {
		cardNumber += fmt.Sprintf("%d", rand.Intn(10))
	}

	// Форматируем как XXXX-XXXX-XXXX-XXXX
	return fmt.Sprintf("%s-%s-%s-%s",
		cardNumber[0:4], cardNumber[4:8], cardNumber[8:12], cardNumber[12:16])
}

func (r *CardRepository) cardNumberExists(cardNumber string) bool {
	query := `SELECT EXISTS(SELECT 1 FROM bank_cards WHERE card_number = $1)`
	var exists bool
	r.db.QueryRow(query, cardNumber).Scan(&exists)
	return exists
}
