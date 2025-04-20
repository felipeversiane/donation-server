CREATE TABLE IF NOT EXISTS addresses (
    id UUID PRIMARY KEY,
    user_id UUID NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    country VARCHAR(100),
    zip_code VARCHAR(20),
    state VARCHAR(100),
    city VARCHAR(100),
    neighborhood VARCHAR(100),
    street VARCHAR(255),
    number VARCHAR(20),
    complement VARCHAR(255)
);
