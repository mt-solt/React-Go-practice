CREATE TABLE IF NOT EXISTS random (
    id UUID PRIMARY KEY,
    random_val BIGINT NOT NULL,
    user_id VARCHAR(255) NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP NOT NULL DEFAULT NOW()
);

-- インデックスを作成
CREATE INDEX IF NOT EXISTS idx_random_user_id ON random(user_id);
CREATE INDEX IF NOT EXISTS idx_random_created_at ON random(created_at); 