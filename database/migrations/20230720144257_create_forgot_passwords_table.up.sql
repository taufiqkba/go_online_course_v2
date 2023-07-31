create table forgot_passwords(
    id INT NOT NULL AUTO_INCREMENT,
    user_id INT NULL,
    valid boolean NOT NULL DEFAULT 1,
    code VARCHAR(255) NOT NULL,
    expired_at TIMESTAMP NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NULL,
    deleted_at TIMESTAMP NULL,
    PRIMARY KEY (id),
    INDEX idx_forgot_passwords_user_id(user_id),
    CONSTRAINT FK_forgot_passwords_user_id FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE SET NULL
) ENGINE = INNODB AUTO_INCREMENT = 1 DEFAULT CHARSET = utf8;