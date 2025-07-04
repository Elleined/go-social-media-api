CREATE TABLE IF NOT EXISTS provider_type (
    id BIGINT UNSIGNED PRIMARY KEY AUTO_INCREMENT,
    name VARCHAR(25) NOT NULL UNIQUE
);

INSERT INTO provider_type(name)
VALUES
    ("LOCAL"),
    ("MICROSOFT"),
    ("FACEBOOK"),
    ("GOOGLE");

CREATE TABLE IF NOT EXISTS user_social (
    id BIGINT UNSIGNED PRIMARY KEY AUTO_INCREMENT,
    created_at DATETIME NOT NULL DEFAULT NOW(),
    provider_id VARCHAR(50) NOT NULL UNIQUE,

    user_id BIGINT UNSIGNED NOT NULL,
    provider_type_id BIGINT UNSIGNED NOT NULL,
    FOREIGN KEY (user_id) REFERENCES user(id),
    FOREIGN KEY (provider_type_id) REFERENCES provider_type(id)
);

CREATE UNIQUE INDEX idx_provider_id ON user_social(provider_id);