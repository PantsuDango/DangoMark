CREATE DATABASE `dango_mark`;

USE `dango_mark`;

CREATE TABLE `user` (
    `id` int(11) NOT NULL AUTO_INCREMENT,
    `user` varchar(32) NOT NULL COMMENT '昵称',
    `password` varchar(255) NOT NULL COMMENT '密码',
    `ip` varchar(32) NOT NULL COMMENT 'ip地址',
    `total` varchar(32) NOT NULL DEFAULT 1 COMMENT '登录次数',
    `createtime` datetime NOT NULL COMMENT '注册时间',
    `lastupdate` datetime NOT NULL COMMENT '最新登录时间',
    PRIMARY KEY (`id`),
    UNIQUE KEY `user` (`user`),
    KEY `ip` (`ip`)
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8mb4 COMMENT='用户信息表';


CREATE TABLE `image_data` (
    `id` int(11) NOT NULL AUTO_INCREMENT,
    `url` varchar(128) NOT NULL COMMENT '图片链接',
    `language` varchar(16) DEFAULT NULL COMMENT '语种: CH-中文, ENG-英语, JAP-日文, KOR-韩语',
    `suggestion` text DEFAULT NULL COMMENT '预标注建议',
    `mark_result` text DEFAULT NULL COMMENT '标注结果',
    `quality_result` text DEFAULT NULL COMMENT '复检结果',
    `coordinate_json` text DEFAULT NULL COMMENT '文字坐标json字符串',
    `status` tinyint(4) NOT NULL DEFAULT 0 COMMENT '数据状态: 0-未标注, 1-已标注未复检, 2-已标注已复检, 3-已用于训练, 4-无意义数据',
    `createtime` datetime NOT NULL COMMENT '创建时间',
    `lastupdate` datetime NOT NULL COMMENT '修改时间',
    PRIMARY KEY (`id`),
    UNIQUE KEY (`url`),
    KEY (`status`),
    KEY (`language`, `status`)
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8mb4 COMMENT='图片标注数据表';