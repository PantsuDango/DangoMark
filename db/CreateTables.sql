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