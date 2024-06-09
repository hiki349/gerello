CREATE TABLE IF NOT EXISTS projects (
    id UUID PRIMARY KEY,
    title VARCHAR(255) NOT NULL,
    description VARCHAR(1000),
    created_at timestamp NOT NULL,
    updated_at timestamp
);
