CREATE TABLE transaction_history (
    id TEXT PRIMARY KEY,
    voucher_id INTEGER NOT NULL,
    user_id INTEGER NOT NULL, 
    points_redeemed INTEGER,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);
