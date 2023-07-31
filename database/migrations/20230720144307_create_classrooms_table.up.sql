create table class_rooms(
    id INT NOT NULL AUTO_INCREMENT,
    user_id INT NULL,
    product_id INT NULL,
    created_by INT NULL,
    updated_by INT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NULL,
    deleted_at TIMESTAMP NULL,
    PRIMARY KEY (id),
    INDEX idx_class_rooms_user_id (user_id),
    INDEX idx_class_rooms_product_id (product_id),
    INDEX idx_class_rooms_created_by (created_by),
    INDEX idx_class_rooms_upadted_by (updated_by),
    CONSTRAINT FK_class_rooms_user_id FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE SET NULL,
    CONSTRAINT FK_class_rooms_product_id FOREIGN KEY (product_id) REFERENCES products(id) ON DELETE SET NULL,
    CONSTRAINT FK_class_rooms_created_by FOREIGN KEY (created_by) REFERENCES users(id) ON DELETE SET NULL,
    CONSTRAINT FK_class_rooms_updated_by FOREIGN KEY (updated_by) REFERENCES users(id) ON DELETE SET NULL
) ENGINE = INNODB AUTO_INCREMENT = 1 DEFAULT CHARSET = utf8;