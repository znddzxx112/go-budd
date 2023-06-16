use `budd`;

DROP TABLE IF EXISTS `user`;
CREATE TABLE `user`(
    `id` int(10) unsigned NOT NULL AUTO_INCREMENT COMMENT '主键',
    `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    `account` varchar(100) NOT NULL DEFAULT '' COMMENT '账户名称',
    `true_name` varchar(10) NOT NULL DEFAULT '' COMMENT '真实姓名',
    `mobile` varchar(20) NOT NULL DEFAULT '' COMMENT '手机号码',
    `password` varchar(255) NOT NULL DEFAULT '' COMMENT '登录密码',
    `status` tinyint(3) NOT NULL DEFAULT '1' COMMENT '1:可用 2：已删除',
    `last_login` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '上次登录时间',
    `this_login` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '本次登录时间',
    `create_user_id` int(10) NOT NULL DEFAULT '0' COMMENT '创建本用户的用户id',
    PRIMARY KEY (`id`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=10001 DEFAULT CHARSET=utf8mb4 ROW_FORMAT=DYNAMIC;

DROP TABLE IF EXISTS `graph_valid`;
CREATE TABLE `graph_valid`(
   `id` int(10) unsigned NOT NULL AUTO_INCREMENT COMMENT '主键',
   `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
   `code` varchar(20) NOT NULL DEFAULT '' COMMENT '图形验证码',
   `token` varchar(100) NOT NULL DEFAULT '' COMMENT '凭证',
   PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8mb4;


