CREATE TABLE IF NOT EXISTS `users`
(
    `id`         varchar(255) NOT NULL COMMENT '主键ID',
    `username`   varchar(255) NOT NULL COMMENT '用户名',
    `password`   varchar(255) NOT NULL COMMENT '密码',
    `nickname`   varchar(100) NOT NULL COMMENT '昵称',
    `avatar`     varchar(255)          DEFAULT NULL COMMENT '头像',
    `email`      varchar(255) NOT NULL COMMENT '邮箱',
    `phone`      varchar(100) NOT NULL COMMENT '电话',
    `created_at` datetime     NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    `updated_at` datetime     NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
    `deleted_at` datetime              DEFAULT NULL COMMENT '删除时间',
    PRIMARY KEY (`id`),
    UNIQUE KEY `idx_username` (`username`),
    KEY `idx_email` (`email`),
    KEY `idx_phone` (`phone`),
    KEY `idx_createdAt` (`created_at`)
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4;