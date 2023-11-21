CREATE TYPE AccountType AS ENUM ('savings', 'current');

CREATE TABLE Accounts
(
    id SERIAL PRIMARY KEY,
    user_id INT REFERENCES Users(id) NOT NULL,
    account_number VARCHAR(255) NOT NULL,
    account_type AccountType NOT NULL,
    balance DECIMAL(10, 2) NOT NULL,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

CREATE TRIGGER set_timestamp
BEFORE UPDATE ON Accounts
FOR EACH ROW
EXECUTE PROCEDURE trigger_set_timestamp();