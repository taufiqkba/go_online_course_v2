create table products_categories(
    id INT NOT NULL AUTO_INCREMENT,
    name VARCHAR(255) NOT NULL,
    image VARCHAR(255) NOT NULL,
    created_by INT NULL,
    updated_by INT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NULL,
    deleted_at TIMESTAMP NULL,
    PRIMARY KEY(id),
    INDEX idx_product_categories_created_by(created_by),
    INDEX idx_product_categories_updated_by(updated_by),
    CONSTRAINT FK_products_categories_created_by FOREIGN KEY (created_by) REFERENCES admins(id) ON DELETE SET NULL,
    CONSTRAINT FK_products_categories_updated_by FOREIGN KEY (updated_by) REFERENCES admins(id) ON DELETE SET NULL
) ENGINE = INNODB AUTO_INCREMENT = 1 DEFAULT CHARSET = utf8;