CREATE TABLE IF NOT EXISTS documents (
    id UUID PRIMARY KEY,
    type VARCHAR(20) NOT NULL,
    value VARCHAR(20) NOT NULL UNIQUE,
    created_at TIMESTAMP NOT NULL DEFAULT now(),
    updated_at TIMESTAMP NOT NULL DEFAULT now(),

    CONSTRAINT chk_document_type CHECK (type IN ('cpf', 'cnpj'))
);


CREATE UNIQUE INDEX IF NOT EXISTS idx_documents_value ON documents(value);
CREATE INDEX IF NOT EXISTS idx_documents_type ON documents(type);
