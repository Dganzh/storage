CREATE TABLE `tb_file` (
    `id` int(11) NOT NULL AUTO_INCREMENT,
    `sha1` char(40) NOT NULL DEFAULT '' COMMENT '文件hash',
    `name` varchar(256) NOT NULL DEFAULT '' COMMENT '文件名',
    `size` bigint(20) DEFAULT '0' COMMENT '文件大小',
    `path` varchar(1024) NOT NULL DEFAULT '' COMMENT '文件存储位置',
    `create_at` datetime default NOW() COMMENT '创建日期',
    `update_at` datetime default NOW() on update current_timestamp() COMMENT '更新日期',
    `status` int(11) NOT NULL DEFAULT '0' COMMENT '状态(可用/禁用/已删除等状态)',
    `ext1` int(11) DEFAULT  '0' COMMENT '备用字段',
    `ext2` text COMMENT '备用字段2',
    PRIMARY KEY (`id`),
    UNIQUE KEY `idx_file_hash` (`sha1`),
    KEY `idx_status` (`status`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;


CREATE TABLE `tb_user` (
    `id` int(11) NOT NULL AUTO_INCREMENT,
    `user_name` varchar(64) NOT NULL DEFAULT '' COMMENT '用户名',
    `user_pwd` varchar(256) NOT NULL DEFAULT '' COMMENT '用户密码',
    `email` varchar(64) DEFAULT '' COMMENT '邮箱',
    `phone` varchar(128) DEFAULT '' COMMENT '手机号',
    `email_validated` tinyint(1) DEFAULT 0 COMMENT '邮箱是否已验证',
    `phone_validated` tinyint(1) DEFAULT 0 COMMENT '手机号是否已验证',
    `signup_at` datetime DEFAULT CURRENT_TIMESTAMP COMMENT '注册日期',
    `last_active` datetime DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '最后活跃时间戳',
    `profile` text COMMENT '用户属性',
    `status` int(11) NOT NULL DEFAULT '0' COMMENT '账户状态（启用/禁用/锁定/标记删除)',
    PRIMARY KEY (`id`),
    UNIQUE KEY `idx_phone` (`phone`),
    KEY `idx_status` (`status`)
) ENGINE=InnoDB AUTO_INCREMENT=5 DEFAULT CHARSET=utf8mb4;

