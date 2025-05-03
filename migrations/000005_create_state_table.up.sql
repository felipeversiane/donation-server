CREATE TABLE IF NOT EXISTS state (
    id UUID PRIMARY KEY,
    name VARCHAR(100) NOT NULL,
    uf VARCHAR(10) NOT NULL,
    country_id UUID REFERENCES country(id) ON DELETE SET NULL,
    created_at TIMESTAMP NOT NULL DEFAULT now(),
    updated_at TIMESTAMP NOT NULL DEFAULT now()
);

CREATE INDEX IF NOT EXISTS idx_state_name ON state(name);
CREATE INDEX IF NOT EXISTS idx_state_uf ON state(uf);
CREATE INDEX IF NOT EXISTS idx_country_id ON state(country_id);
