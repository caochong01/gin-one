CREATE TABLE `diary` (
  `id` int NOT NULL AUTO_INCREMENT COMMENT '主键ID',
  `diary_title` varchar(200) NOT NULL COMMENT '标题',
  `diary_detail` text COMMENT '内容',
  `createdAt` datetime NOT NULL COMMENT '创建日期',
  `updatedAt` datetime DEFAULT NULL COMMENT '修改日期',
  `deletedAt` datetime DEFAULT NULL COMMENT '删除日期',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci COMMENT='日记表'



CREATE TABLE `comment` (
  `id` int NOT NULL AUTO_INCREMENT,
  `diary_id` int NOT NULL,
  `comment_content` text CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL,
  `createdAt` datetime NOT NULL COMMENT '创建日期',
  `updatedAt` datetime DEFAULT NULL COMMENT '修改日期',
  `deletedAt` datetime DEFAULT NULL COMMENT '删除日期',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci COMMENT='评论表'



CREATE TABLE `praise` (
  `id` int NOT NULL AUTO_INCREMENT,
  `diary_id` int NOT NULL COMMENT '文章ID',
  `praise_num` int NOT NULL DEFAULT '0' COMMENT '点赞数',
  `createdAt` datetime NOT NULL COMMENT '创建日期',
  `updatedAt` datetime DEFAULT NULL COMMENT '修改日期',
  `deletedAt` datetime DEFAULT NULL COMMENT '删除日期',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci COMMENT='点赞表'
