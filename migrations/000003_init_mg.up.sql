CREATE TABLE IF NOT EXISTS refresh_token (
    id BIGINT UNSIGNED PRIMARY KEY AUTO_INCREMENT,
    created_at DATETIME NOT NULL DEFAULT NOW(),
    token CHAR(36) NOT NULL UNIQUE,
    expires_at DATETIME NOT NULL,
    revoked_at DATETIME DEFAULT NULL,
    user_id BIGINT UNSIGNED NOT NULL,
    FOREIGN KEY (user_id) REFERENCES user(id)
);

CREATE INDEX idx_created_at ON refresh_token(created_at);
CREATE UNIQUE INDEX token_idx ON refresh_token(token);
