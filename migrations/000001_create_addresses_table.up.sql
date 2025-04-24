CREATE TABLE IF NOT EXISTS addresses (
    id UUID PRIMARY KEY,
    country VARCHAR(100),
    zip_code VARCHAR(20),
    state VARCHAR(100),
    city VARCHAR(100),
    neighborhood VARCHAR(100),
    street VARCHAR(255),
    number VARCHAR(20),
    complement VARCHAR(255)
    created_at TIMESTAMP NOT NULL DEFAULT now(),
    updated_at TIMESTAMP NOT NULL DEFAULT now()
);

CREATE INDEX IF NOT EXISTS idx_addresses_country ON addresses(country);
CREATE INDEX IF NOT EXISTS idx_addresses_state ON addresses(state);
CREATE INDEX IF NOT EXISTS idx_addresses_city ON addresses(city);
CREATE INDEX IF NOT EXISTS idx_addresses_zip_code ON addresses(zip_code);