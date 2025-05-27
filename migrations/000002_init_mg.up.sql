CREATE UNIQUE INDEX idx_email ON user(email);
CREATE INDEX idx_first_last_name ON user (last_name, first_name);

CREATE UNIQUE INDEX idx_name ON emoji(name);

CREATE INDEX idx_create_at ON post(created_at);
CREATE INDEX idx_create_at ON comment(created_at);
CREATE INDEX idx_create_at ON post_reaction(created_at);
CREATE INDEX idx_create_at ON comment_reaction(created_at);