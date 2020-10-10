
-- +migrate Up
DROP TABLE IF EXISTS `users`;
CREATE TABLE `users`
(
    id         bigint unsigned AUTO_INCREMENT COMMENT 'ユーザーID',
    name       varchar(255) NOT NULL COMMENT 'ユーザー名',
    email      varchar(255) NOT NULL COMMENT 'メールアドレス',
    password   varchar(255) NOT NULL COMMENT 'パスワード',
    created_at datetime DEFAULT CURRENT_TIMESTAMP COMMENT '作成日時',
    updated_at datetime DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新日時',
    PRIMARY KEY (`id`),
    UNIQUE KEY `uq_email` (`email`),
    KEY `idx_name` (`name`)
) ENGINE = INNODB
  CHARSET = utf8mb4 COMMENT ='ユーザー';

INSERT INTO users(`name`, `email`, `password`)
VALUES ('test', 'test@shinjiezumi.work', '$2a$10$LvsHcEMnYFFe1taM1sSr1eLjoZ740O7o5M5aFB05RuI6yY0mHa.1u');

-- +migrate Down
DROP TABLE `users`;