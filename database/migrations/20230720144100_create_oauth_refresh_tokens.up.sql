create table oauth_refresh_tokens(
    id INT NOT NULL AUTO_INCREMENT,
    oauth_access_token_id INT NULL,
    user_id INT NOT NULL,
    token VARCHAR(255) NULL,
    expired_at TIMESTAMP NULL,
    created_by INT NULL,
    updated_by INT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NULL,
    deleted_at TIMESTAMP NULL,
    PRIMARY KEY (id),
    UNIQUE KEY oauth_refresh_tokens_token_unique(token),
    INDEX idx_oauth_refresh_tokens_oauth_access_token_id(oauth_access_token_id),
    INDEX idx_oauth_refresh_tokens_token(token),
    INDEX idx_oauth_refresh_tokens_created_by(created_by),
    INDEX idx_oauth_refresh_tokens_updated_by(updated_by),
    CONSTRAINT FK_oauth_refresh_tokens_oauth_access_token_id FOREIGN KEY (oauth_access_token_id) REFERENCES oauth_access_tokens(id) ON DELETE SET NULL
) ENGINE = INNODB AUTO_INCREMENT = 1 DEFAULT CHARSET = utf8;