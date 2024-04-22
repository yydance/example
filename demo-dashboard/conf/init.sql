CREATE DATABASE IF NOT EXISTS demo_dashboard;
USE demo_dashboard;

/*
CREATE TABLE IF NOT EXISTS plugin_config (
`name`   JSON      NOT NULL   DEFAULT ''       COMMENT '插件模版名称',
`ids`   varchar(32)      NOT NULL   DEFAULT ''       COMMENT '插件列表，以逗号分割多个插件ID',
`desc`   varchar(128)         DEFAULT ''       COMMENT '插件模版描述',
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

CREATE TABLE IF NOT EXISTS plugins_list (
`id`   int   PRIMARY KEY   NOT NULL      AUTO_INCREMENT    COMMENT '自增ID',
`name`   varchar(32)      NOT NULL   DEFAULT ''       COMMENT '插件名称',
KEY `id_idx` (id)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

CREATE TABLE IF NOT EXISTS plugins (
`id`   int   PRIMARY KEY   NOT NULL      AUTO_INCREMENT    COMMENT '自增ID',
`route_plugin`   JSON      NOT NULL   DEFAULT ''       COMMENT '插件配置信息',
KEY `id_idx` (id)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

CREATE TABLE IF NOT EXISTS global_rules (
`name`   JSON      NOT NULL   DEFAULT ''       COMMENT '全局插件配置信息',
`id`   int   PRIMARY KEY   NOT NULL      AUTO_INCREMENT    COMMENT '自增ID',
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
*/

CREATE TABLE IF NOT EXISTS upstream (
    `id` varchar(32) NOT NULL DEFAULT '' COMMENT '唯一ID',
    `name` varchar(64) NOT NULL DEFAULT '' COMMENT '名称',
    `type` varchar(32)  NOT NULL DEFAULT '' COMMENT '上游类型',
    `desc` varchar(128) DEFAULT '' COMMENT '描述',
    `create_time` varchar(16) NOT NULL   COMMENT '创建时间',
    `update_time` varchar(16) NOT NULL   COMMENT '修改时间',
    KEY `id_idx` (id)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

CREATE TABLE IF NOT EXISTS service (
    `id` varchar(32) NOT NULL DEFAULT '' COMMENT '唯一ID',
    `name` varchar(64) NOT NULL DEFAULT '' COMMENT '名称',
    `desc` varchar(128) DEFAULT '',
    `create_time` varchar(16) NOT NULL   COMMENT '创建时间',
    `update_time` varchar(16) NOT NULL   COMMENT '修改时间',
    KEY `id_idx` (id)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

CREATE TABLE IF NOT EXISTS route (
    `id` varchar(32) NOT NULL COMMENT '唯一ID',
    `name` varchar(64) NOT NULL COMMENT '名称',
    `hosts` varchar(128) NOT NULL COMMENT '域名列表，以逗号分割',
    `uris` varchar(128) NOT NULL COMMENT '路径，以逗号分割',
    `status` tinyint NOT NULL DEFAULT 0 COMMENT '状态',
    `create_time` varchar(16) NOT NULL   COMMENT '创建时间',
    `update_time` varchar(16) NOT NULL   COMMENT '修改时间',
    KEY `id_idx` (id)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

