-- Active: 1754455246108@@localhost@5432@card_service_db
-- Active: 1754455246108@@localhost@5432@postgres
DROP DATABASE IF EXISTS card_service_db;
CREATE DATABASE card_service_db;
\c card_service_db;

-- Table for card operators (CardOperator)
CREATE TABLE card_operators (
    id SERIAL PRIMARY KEY,
    name VARCHAR(50) NOT NULL,
    code VARCHAR(10) NOT NULL UNIQUE
);

-- Table for users (User)
CREATE TABLE users (
    id SERIAL PRIMARY KEY,
    first_name VARCHAR(30) NOT NULL,
    last_name VARCHAR(30) NOT NULL,
    email VARCHAR(150) UNIQUE NOT NULL,
    phone VARCHAR(20),
    birth_date DATE,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- Table for bank cards (BankCard)
CREATE TABLE bank_cards (
    id SERIAL PRIMARY KEY,
    user_id INTEGER NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    card_number VARCHAR(19) NOT NULL UNIQUE,
    operator_id INTEGER NOT NULL REFERENCES card_operators(id),
    issue_date DATE NOT NULL,
    expiry_date DATE NOT NULL,
    is_active BOOLEAN DEFAULT true,
    balance DECIMAL(15,2) DEFAULT 0.00,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- Table for transactions (Transaction)
CREATE TABLE transactions (
    id SERIAL PRIMARY KEY,
    from_card_id INTEGER NOT NULL REFERENCES bank_cards(id) ON DELETE CASCADE,
    to_card_id INTEGER NOT NULL REFERENCES bank_cards(id) ON DELETE CASCADE,
    amount DECIMAL(15,2) NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- Test data for card operators
INSERT INTO card_operators (name, code) VALUES
    ('Visa', 'VISA'),
    ('MasterCard', 'MC'),
    ('МИР', 'MIR'),
    ('American Express', 'AMEX'),
    ('UnionPay', 'UPAY');

-- Test data for users
INSERT INTO users (first_name, last_name, email, phone, birth_date) VALUES
    ('Jhon', 'Smith', 'jhon.smith@test.com', '+7-900-123-45-67', '1990-01-15'),
    ('Carl', 'Johnson', 'carl.johnson@test.com', '+7-900-765-43-21', '1985-05-20'),
    ('Alex', 'Parker', 'alex.parker@test.com', '+7-900-555-77-88', '1992-12-03');

-- Test data for bank cards
INSERT INTO bank_cards (user_id, card_number, operator_id, issue_date, expiry_date, balance) VALUES
    (1, '4111111111111111', 1, '2022-01-15', '2027-01-31', 15000.50),
    (1, '5555555555554444', 2, '2023-03-10', '2028-03-31', 50000.00),
    (2, '2200123456789012', 3, '2021-06-20', '2026-06-30', 25000.75),
    (3, '3782822463100051', 4, '2024-02-01', '2029-02-01', 100000.00);

-- Test data for transactions
INSERT INTO transactions (from_card_id, to_card_id, amount) VALUES
    (1, 2, 1000.00),
    (2, 3, 500.50),
    (3, 1, 750.75),
    (1, 2, 2000.00),
    (2, 3, 300.00),
    (3, 1, 1500.00);