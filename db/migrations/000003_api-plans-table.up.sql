CREATE TABLE api_plans (
    id UUID PRIMARY KEY,
    api_id UUID REFERENCES apis(id),
    name VARCHAR(255) NOT NULL,
    value_cents INTEGER,
    request_limit INTEGER NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);
