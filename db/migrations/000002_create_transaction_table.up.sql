CREATE TABLE transactions (
    id UUID PRIMARY KEY,
    source_account JSONB NOT NULL,
    destination_account JSONB NOT NULL,
    amount NUMERIC(20, 2) NOT NULL,
    status VARCHAR(50) NOT NULL,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NOT NULL
);
