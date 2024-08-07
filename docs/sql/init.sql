INSERT IGNORE INTO roles(`id`, `name`, `identify`, `created_at`, `updated_at`, `deleted_at`) VALUES('cqp4dpe8h0jqqf2u0cs0', '系统管理员', 'admin', NOW(), NOW(), NULL);
INSERT IGNORE INTO roles(`id`, `name`, `identify`, `created_at`, `updated_at`, `deleted_at`) VALUES('cqp4ec68h0jr3qbc3pug', '普通用户', 'normal', NOW(), NOW(), NULL);

INSERT IGNORE INTO users (`id`, `username`, `nickname`, `password`, `salt`, `avatar`, `email`, `phone`, `role_id`, `created_at`, `updated_at`, `deleted_at`) 
VALUES ('cqp4hku8h0jk8mqc74e0', 'godfrey', '葛孚雷', '602B6BD0236F381FEB1F925683A41A99', 'waU2NWCdNjb36AAFmYFyVTUJOp9DQ1Hc', NULL, 'godfrey@example.com', '18888888888', 'cqp4dpe8h0jqqf2u0cs0', NOW(), NOW(), NULL);