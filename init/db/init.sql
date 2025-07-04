CREATE DATABASE card_service_db;
\c card_service_db;

-- Создание таблиц
CREATE TABLE card_operators (
                                id SERIAL PRIMARY KEY,
                                name VARCHAR(50) NOT NULL UNIQUE,
                                code VARCHAR(10) NOT NULL UNIQUE,
                                created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE users (
                       id SERIAL PRIMARY KEY,
                       first_name VARCHAR(30) NOT NULL,
                       last_name VARCHAR(30) NOT NULL,
                       email VARCHAR(150) UNIQUE NOT NULL,
                       phone VARCHAR(20),
                       birth_date DATE,
                       created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

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

-- Добавление операторов карт
INSERT INTO card_operators (name, code) VALUES
                                            ('Visa', 'VISA'),
                                            ('Mastercard', 'MC'),
                                            ('American Express', 'AMEX'),
                                            ('Maestro', 'MAESTRO'),
                                            ('МИР', 'MIR'),
                                            ('UnionPay', 'UP'),
                                            ('JCB', 'JCB'),
                                            ('Diners Club', 'DC');

-- Тестовые данные
INSERT INTO users (first_name, last_name, email, phone, birth_date) VALUES
                                                                        ('Иван', 'Иванов', 'ivan.ivanov@test.com', '+7-900-123-45-67', '1990-01-15'),
                                                                        ('Мария', 'Петрова', 'maria.petrova@test.com', '+7-900-765-43-21', '1985-05-20'),
                                                                        ('Алексей', 'Сидоров', 'alexey.sidorov@test.com', '+7-900-555-77-88', '1992-12-03');

INSERT INTO bank_cards (user_id, card_number, operator_id, issue_date, expiry_date, balance) VALUES
                                                                                                 (1, '4111-1111-1111-1111', 1, '2022-01-15', '2027-01-31', 15000.50),
                                                                                                 (1, '5555-5555-5555-4444', 2, '2023-03-10', '2028-03-31', 50000.00),
                                                                                                 (2, '2200-1234-5678-9012', 5, '2021-06-20', '2026-06-30', 25000.75);