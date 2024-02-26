CREATE TABLE IF NOT EXISTS orders (
    id VARCHAR(26) PRIMARY KEY,
    price FLOAT,
    tax FLOAT,
    final_price FLOAT
);