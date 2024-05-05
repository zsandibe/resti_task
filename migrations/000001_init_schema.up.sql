CREATE TABLE accounts
(
    id         SERIAL PRIMARY KEY,
    username   VARCHAR(255) NOT NULL UNIQUE,
    password   VARCHAR(255) NOT NULL,
    balance     DECIMAL NOT NULL,   
    created_at TIMESTAMP    NOT NULL DEFAULT NOW()
);

CREATE TABLE transactions
(
    id SERIAL PRIMARY KEY,
    amount DECIMAL NOT NULL,
    account_id INT REFERENCES accounts(id),
    "group" TEXT NOT NULL,
    account2_id INT,
    created_at TIMESTAMP NOT NULL DEFAULT NOW()
);


CREATE INDEX idx_account_id ON transactions(account_id);
CREATE INDEX idx_account2_id ON transactions(account2_id);

