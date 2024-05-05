CREATE TABLE accounts
(
    id         SERIAL PRIMARY KEY,
    name   VARCHAR(255) NOT NULL,
    balance     DECIMAL(15,2) NOT NULL,   
    created_at TIMESTAMP    NOT NULL
);

CREATE TABLE transactions
(
    id SERIAL PRIMARY KEY,
    value DECIMAL(15,2) NOT NULL,
    account_id INT REFERENCES accounts(id),
    "group" TEXT NOT NULL,
    account2_id INT,
    created_at TIMESTAMP NOT NULL
);


CREATE INDEX idx_account_id ON transactions(account_id);
CREATE INDEX idx_account2_id ON transactions(account2_id);

