CREATE TABLE transaction_history (
    id VARCHAR(50) PRIMARY KEY,
    voucher_id INT NOT NULL,
    user_id INT NOT NULL,
    points_redeemed INT,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    FOREIGN KEY (voucher_id) REFERENCES vouchers(id) ON DELETE CASCADE,
    FOREIGN KEY (user_id) REFERENCES customers(id) ON DELETE CASCADE
);