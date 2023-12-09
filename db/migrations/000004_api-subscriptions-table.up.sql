CREATE TABLE api_subscriptions (
    id UUID PRIMARY KEY,
    user_id UUID REFERENCES users(id),
    api_plan_id UUID REFERENCES api_plans(id),
    api_name VARCHAR(255),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);
