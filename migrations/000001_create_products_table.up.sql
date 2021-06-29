CREATE TABLE products (
    id SERIAL PRIMARY KEY,
    title VARCHAR(255),    
    description VARCHAR(255),
    rating INTEGER,
    image VARCHAR(155),
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMPTZ
);