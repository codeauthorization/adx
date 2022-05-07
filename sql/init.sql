drop table `demander_info`;
CREATE TABLE `demander_info`
(
    `id` bigint unsigned NOT NULL AUTO_INCREMENT COMMENT '主键',
    `name` varchar(50) NOT NULL  COMMENT '广告主名称',
    `username` varchar(50) NOT NULL  COMMENT '账户',
    `password` varchar(50) NOT NULL  COMMENT '密码',
    `create_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    `modify_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '修改时间',
    PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci COMMENT='广告主信息';