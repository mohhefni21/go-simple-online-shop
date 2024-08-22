CREATE TABLE IF NOT EXISTS transactions(
    id SERIAL PRIMARY KEY,
    user_public_id VARCHAR(100) NOT NULL,
    product_id VARCHAR(50) NOT NULL,
    product_price INTEGER NOT NULL,
    amount INTEGER NOT NULL,
    sub_total INTEGER NOT NULL,
    platform_fee INTEGER NOT NULL,
    grand_total INTEGER NOT NULL,
    status INTEGER NOT NULL,
    product_snapshot JSONB NOT NULL,
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP DEFAULT NOW()
);