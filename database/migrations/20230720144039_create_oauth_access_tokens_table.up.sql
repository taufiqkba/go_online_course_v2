create table oauth_access_token (
    id INT NOT NULL AUTO_INCREMENT,
    oauth_client_id INT NULL,
    user_id INT NOT NULL,
    token VARCHAR(255) NULL,
    scope VARCHAR(255) NULL,
    expired_at TIMESTAMP NULL,
    created_by INT NULL,
    updated_by INT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NULL,
    deleted_at TIMESTAMP NULL,
    PRIMARY KEY (id),
    UNIQUE KEY oauth_access_tokens_token_unique(token),
    INDEX idx_oauth_access_tokens_token(token),
    INDEX idx_oauth_access_tokens_oauth_client_id(oauth_client_id),
    INDEX idx_oauth_access_tokens_oauth_created_by(created_by),
    INDEX idx_oauth_access_tokens_oauth_updated_by(updated_by),
    CONSTRAINT FK_oauth_access_tokens_oauth_client_id FOREIGN KEY (oauth_client_id) REFERENCES oauth_clients(id) ON DELETE SET NULL
) ENGINE = INNODB AUTO_INCREMENT = 1 DEFAULT CHARSET = utf8;