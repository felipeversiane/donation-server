CREATE TABLE IF NOT EXISTS document (
    id UUID PRIMARY KEY NOT NULL,
    value VARCHAR(50) NOT NULL UNIQUE,
    user_id UUID UNIQUE REFERENCES users(id) ON DELETE CASCADE,
    created_at TIMESTAMP NOT NULL DEFAULT now(),
    updated_at TIMESTAMP NOT NULL DEFAULT now()
);

CREATE INDEX IF NOT EXISTS idx_document_value ON document(value);
