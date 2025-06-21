CREATE TABLE IF NOT EXISTS users (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    username VARCHAR(255) UNIQUE NOT NULL,
    password_hash VARCHAR(255) NOT NULL
);

-- インデックスを作成
CREATE INDEX IF NOT EXISTS idx_users_username ON users(username);

-- 注意: 既存のrandomテーブルのuser_idはVARCHAR型のため、
-- 外部キー制約は後でuser_idカラムをUUID型に変更してから追加する予定 