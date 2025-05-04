CREATE TABLE IF NOT EXISTS address (
    id UUID PRIMARY KEY NOT NULL,
    zip_code VARCHAR(20) NOT NULL,
    neighborhood VARCHAR(100) NOT NULL,
    street VARCHAR(255) NOT NULL,
    number VARCHAR(20),
    complement VARCHAR(255),
    user_id UUID UNIQUE REFERENCES users(id) ON DELETE CASCADE,
    city_id UUID NOT NULL REFERENCES city(id) ON DELETE RESTRICT,
    created_at TIMESTAMP NOT NULL DEFAULT now(),
    updated_at TIMESTAMP NOT NULL DEFAULT now()
);

CREATE INDEX IF NOT EXISTS idx_address_zip_code ON address(zip_code);
CREATE INDEX IF NOT EXISTS idx_address_city_id ON address(city_id);