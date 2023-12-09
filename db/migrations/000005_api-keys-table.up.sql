CREATE TABLE api_keys (
    id UUID PRIMARY KEY,
    api_subscription_id UUID REFERENCES api_subscriptions(id),
    user_id UUID REFERENCES users(id),
    key VARCHAR(64) NOT NULL,
    valid BOOLEAN NOT NULL,
    quota_used INT NOT NULL,
    renewal_date DATE DEFAULT (CURRENT_DATE + INTERVAL '1 month'),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);
