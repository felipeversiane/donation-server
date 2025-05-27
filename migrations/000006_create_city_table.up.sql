CREATE TABLE IF NOT EXISTS city (
    id UUID PRIMARY KEY NOT NULL,
    name VARCHAR(100) NOT NULL,
    state_id UUID REFERENCES state(id) ON DELETE SET NULL,
    created_at TIMESTAMP NOT NULL DEFAULT now(),
    updated_at TIMESTAMP NOT NULL DEFAULT now()
);

CREATE INDEX IF NOT EXISTS idx_city_name ON city(name);
CREATE INDEX IF NOT EXISTS idx_city_state_id ON city(state_id);
