create database if not exists demo_dashboard;

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



