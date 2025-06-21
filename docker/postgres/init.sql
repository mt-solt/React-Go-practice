-- randomuserとしてテーブルを作成
CREATE TABLE IF NOT EXISTS random (
    id UUID PRIMARY KEY,
    random_val BIGINT NOT NULL,
    user_id VARCHAR(255) NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP NOT NULL DEFAULT NOW()
);

-- インデックスの作成
CREATE INDEX IF NOT EXISTS idx_random_user_id ON random(user_id);
CREATE INDEX IF NOT EXISTS idx_random_created_at ON random(created_at);

-- randomuserにテーブルへの権限を付与
GRANT ALL PRIVILEGES ON TABLE random TO randomuser;
GRANT ALL PRIVILEGES ON ALL SEQUENCES IN SCHEMA public TO randomuser;

-- テストデータの挿入
INSERT INTO random (id, random_val, user_id, created_at, updated_at) VALUES
    ('550e8400-e29b-41d4-a716-446655440001', 12345, 'user1', NOW(), NOW()),
    ('550e8400-e29b-41d4-a716-446655440002', 67890, 'user1', NOW(), NOW()),
    ('550e8400-e29b-41d4-a716-446655440003', 11111, 'user2', NOW(), NOW()),
    ('550e8400-e29b-41d4-a716-446655440004', 22222, 'user2', NOW(), NOW()),
    ('550e8400-e29b-41d4-a716-446655440005', 33333, 'user3', NOW(), NOW())
ON CONFLICT (id) DO NOTHING; 
