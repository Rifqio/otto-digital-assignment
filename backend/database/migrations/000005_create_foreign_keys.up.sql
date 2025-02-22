PRAGMA foreign_keys = ON;

ALTER TABLE vouchers 
ADD CONSTRAINT fk_vouchers_brand 
FOREIGN KEY (brand_id) REFERENCES brands(id) ON DELETE CASCADE;

ALTER TABLE transaction_history 
ADD CONSTRAINT fk_transaction_voucher 
FOREIGN KEY (voucher_id) REFERENCES vouchers(id) ON DELETE CASCADE;

ALTER TABLE transaction_history 
ADD CONSTRAINT fk_transaction_customer 
FOREIGN KEY (user_id) REFERENCES customers(id) ON DELETE CASCADE;
