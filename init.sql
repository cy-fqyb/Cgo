CREATE TABLE
    users (
        id VARCHAR(36) PRIMARY KEY COMMENT '用户ID',
        name VARCHAR(20) not NULL COMMENT '用户名',
        email VARCHAR(32) NOT NULL COMMENT '邮箱',
        password VARCHAR(32) NOT NULL COMMENT '密码',
        sex VARCHAR(2) NOT NULL COMMENT '性比',
        avatar VARCHAR(100) COMMENT '头像',
        status INT DEFAULT '1' COMMENT '状态，1为正常，0为禁用',
        synopsis VARCHAR(255) DEFAULT '这个人很懒啥都没写' COMMENT '简介',
        create_time DATETIME COMMENT '创建时间',
        update_time DATETIME COMMENT '更新时间'
    );

DROP TABLE IF EXISTS `user_friend`;

CREATE TABLE
    `user_friend` (
        `id` varchar(36) NOT NULL COMMENT '主键',
        `user_id` varchar(36) NOT NULL COMMENT '用户ID',
        `friend_id` varchar(36) NOT NULL COMMENT '好友ID',
        `create_time` datetime DEFAULT NULL COMMENT '创建时间',
        `update_time` datetime DEFAULT NULL COMMENT '更新时间',
        PRIMARY KEY (`id`)
    ) ENGINE = InnoDB DEFAULT CHARSET = utf8mb4 COMMENT = '用户好友表';

DROP TABLE IF EXISTS `msg`;

CREATE TABLE
    `msg`(
        `id` VARCHAR(36) NOT NULL COMMENT '主键',
        `from_id` VARCHAR(36) NOT NULL COMMENT '发送者ID',
        `to_id` VARCHAR(36) NOT NULL COMMENT '接收者ID',
        `msg` VARCHAR(255) NOT NULL COMMENT '消息内容',
        `type` INT NOT NULL COMMENT '消息类型，1为文本，2为图片，3为文件',
        `create_time` DATETIME COMMENT '创建时间',
        `update_time` DATETIME COMMENT '更新时间',
        PRIMARY KEY (`id`)
    ) ENGINE = InnoDB DEFAULT CHARSET = utf8mb4 COMMENT = '消息表';

DROP TABLE IF EXISTS `user_room`;

CREATE TABLE
    `user_room`(
        `id` VARCHAR(36) NOT NULL COMMENT '主键',
        `user_id` VARCHAR(36) NOT NULL COMMENT '用户ID',
        `room_id` VARCHAR(36) NOT NULL COMMENT '房间ID',
        `create_time` DATETIME COMMENT '创建时间',
        `update_time` DATETIME COMMENT '更新时间',
        PRIMARY KEY (`id`)
    ) ENGINE = InnoDB DEFAULT CHARSET = utf8mb4 COMMENT = '用户房间表';

DROP TABLE IF EXISTS `room`;

CREATE Table
    `room`(
        `id` VARCHAR(36) NOT NULL COMMENT '主键',
        `name` VARCHAR(20) NOT NULL COMMENT '房间名称',
        `avatar` VARCHAR(100) COMMENT '房间头像',
        `master_id` VARCHAR(36) NOT NULL COMMENT '房主ID',
        `create_time` DATETIME COMMENT '创建时间',
        `update_time` DATETIME COMMENT '更新时间',
        PRIMARY KEY (`id`)
    ) ENGINE = InnoDB DEFAULT CHARSET = utf8mb4 COMMENT = '房间表';

DROP TABLE IF EXISTS `apply`;

CREATE TABLE
    `apply` (
        `id` varchar(36) NOT NULL COMMENT '主键',
        `user_id` varchar(36) NOT NULL COMMENT '用户ID',
        `apply_id` varchar(36) NOT NULL COMMENT '申请ID',
        `create_time` datetime DEFAULT NULL COMMENT '创建时间',
        `update_time` datetime DEFAULT NULL COMMENT '更新时间',
        PRIMARY KEY (`id`)
    ) ENGINE = InnoDB DEFAULT CHARSET = utf8mb4 COMMENT = '申请表';