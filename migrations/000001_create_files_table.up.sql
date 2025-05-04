CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE IF NOT EXISTS files (
    id UUID PRIMARY KEY NOT NULL,
    name VARCHAR(255) NOT NULL,  
    url TEXT NOT NULL,          
    type VARCHAR(50) NOT NULL,   
    created_at TIMESTAMP NOT NULL DEFAULT now(),
    updated_at TIMESTAMP NOT NULL DEFAULT now()
);

CREATE INDEX IF NOT EXISTS idx_files_type ON files(type);
