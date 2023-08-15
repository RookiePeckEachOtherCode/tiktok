SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- comments: table
DROP TABLE IF EXISTS `comments`;
CREATE TABLE `comments` (
  `id` bigint NOT NULL AUTO_INCREMENT COMMENT '评论ID',
  `user_info_id` bigint DEFAULT NULL COMMENT '用户信息ID',
  `video_id` bigint DEFAULT NULL COMMENT '视频ID',
  `content` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT '评论内容',
  `created_at` datetime(3) DEFAULT NULL COMMENT '评论创建时间',
  PRIMARY KEY (`id`),
  KEY `fk_videos_comments` (`video_id`),
  KEY `fk_user_infos_comments` (`user_info_id`),
  CONSTRAINT `fk_user_infos_comments` FOREIGN KEY (`user_info_id`) REFERENCES `user_infos` (`id`),
  CONSTRAINT `fk_videos_comments` FOREIGN KEY (`video_id`) REFERENCES `videos` (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=10 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- user_favor_videos: table
DROP TABLE IF EXISTS `user_favor_videos`;
CREATE TABLE `user_favor_videos` (
  `user_info_id` bigint NOT NULL COMMENT '用户信息ID',
  `video_id` bigint NOT NULL COMMENT '视频ID',
  PRIMARY KEY (`user_info_id`,`video_id`),
  KEY `fk_user_favor_videos_video` (`video_id`),
  CONSTRAINT `fk_user_favor_videos_user_info` FOREIGN KEY (`user_info_id`) REFERENCES `user_infos` (`id`),
  CONSTRAINT `fk_user_favor_videos_video` FOREIGN KEY (`video_id`) REFERENCES `videos` (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- user_infos: table
DROP TABLE IF EXISTS `user_infos`;
CREATE TABLE `user_infos` (
  `id` bigint NOT NULL AUTO_INCREMENT COMMENT '用户信息ID',
  `name` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT '用户名',
  `follow_count` bigint DEFAULT NULL  COMMENT '关注数',
  `follower_count` bigint DEFAULT NULL COMMENT '粉丝数',
  `is_follow` tinyint(1) DEFAULT NULL COMMENT '是否关注',
  `avatar` varchar(200) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT '头像地址',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=5 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- user_logins: table
DROP TABLE IF EXISTS `user_logins`;
CREATE TABLE `user_logins` (
  `id` bigint NOT NULL AUTO_INCREMENT COMMENT '用户登录ID',
  `user_info_id` bigint DEFAULT NULL COMMENT '用户信息ID',
  `username` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL COMMENT '用户名',
  `password` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL COMMENT '密码',
  PRIMARY KEY (`id`,`username`),
  KEY `fk_user_infos_user` (`user_info_id`),
  CONSTRAINT `fk_user_infos_user` FOREIGN KEY (`user_info_id`) REFERENCES `user_infos` (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=5 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- user_relations: table
DROP TABLE IF EXISTS `user_relations`;
CREATE TABLE `user_relations` (
  `user_info_id` bigint NOT NULL COMMENT '用户信息ID',
  `follow_id` bigint NOT NULL COMMENT '关注用户信息ID',
  PRIMARY KEY (`user_info_id`,`follow_id`),
  KEY `fk_user_relations_follows` (`follow_id`),
  CONSTRAINT `fk_user_relations_follows` FOREIGN KEY (`follow_id`) REFERENCES `user_infos` (`id`),
  CONSTRAINT `fk_user_relations_user_info` FOREIGN KEY (`user_info_id`) REFERENCES `user_infos` (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- videos: table
DROP TABLE IF EXISTS `videos`;
CREATE TABLE `videos` (
  `id` bigint NOT NULL AUTO_INCREMENT COMMENT '视频ID',
  `user_info_id` bigint DEFAULT NULL COMMENT '用户信息ID',
  `play_url` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT '播放链接',
  `cover_url` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT '封面链接',
  `favorite_count` bigint DEFAULT NULL COMMENT '收藏数',
  `comment_count` bigint DEFAULT NULL COMMENT '评论数',
  `is_favorite` tinyint(1) DEFAULT NULL COMMENT '是否收藏',
  `title` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT '视频标题',
  `created_at` datetime(3) DEFAULT NULL COMMENT '视频创建时间',
  `updated_at` datetime(3) DEFAULT NULL COMMENT '视频更新时间',
  PRIMARY KEY (`id`),
  KEY `fk_user_infos_videos` (`user_info_id`),
  CONSTRAINT `fk_user_infos_videos` FOREIGN KEY (`user_info_id`) REFERENCES `user_infos` (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=6 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;
-- chat_records: table
CREATE TABLE `chat_records` (
  `id` bigint NOT NULL AUTO_INCREMENT COMMENT '聊天记录id',
  `user_id` bigint NOT NULL COMMENT '用户id',
  `to_user_id` bigint NOT NULL COMMENT '目标用户id',
  `content` text COMMENT '聊天内容',
  `created_at`  bigint  DEFAULT NULL COMMENT '时间戳',
  PRIMARY KEY (`id`),
  KEY `idx_user_id` (`user_id`),
  KEY `idx_target_id` (`to_user_id`),
  CONSTRAINT `fk_user_id` FOREIGN KEY (`user_id`) REFERENCES `user_infos`(`id`),
  CONSTRAINT `fk_to_user_id` FOREIGN KEY (`to_user_id`) REFERENCES `user_infos`(`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;