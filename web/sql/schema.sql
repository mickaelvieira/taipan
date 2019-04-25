SET NAMES utf8mb4;
SET foreign_key_checks = 0;
SET sql_mode = 'NO_AUTO_VALUE_ON_ZERO';

-- Database

DROP DATABASE IF EXISTS `taipan`;
CREATE DATABASE `taipan` CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_520_ci;

USE `taipan`;

-- bookmarks

DROP TABLE IF EXISTS `bookmarks`;
CREATE TABLE `bookmarks` (
  `id` char(36) COLLATE utf8mb4_unicode_520_ci NOT NULL,
  `url` varchar(255) COLLATE utf8mb4_unicode_520_ci NOT NULL,
  `hash` varchar(100) COLLATE utf8mb4_unicode_520_ci NOT NULL,
  `title` text COLLATE utf8mb4_unicode_520_ci NOT NULL DEFAULT "",
  `description` text COLLATE utf8mb4_unicode_520_ci NOT NULL DEFAULT "",
  `charset` varchar(255) COLLATE utf8mb4_unicode_520_ci NOT NULL DEFAULT "",
  `canonical_url` varchar(255) COLLATE utf8mb4_unicode_520_ci NOT NULL DEFAULT "",
  `status` varchar(50) COLLATE utf8mb4_unicode_520_ci NOT NULL,
  `created_at` datetime NOT NULL,
  `updated_at` datetime NOT NULL,
  `image_cache_key` varchar(255) COLLATE utf8mb4_unicode_520_ci DEFAULT NULL,
  `html_cache_key` varchar(255) COLLATE utf8mb4_unicode_520_ci DEFAULT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `hash_idx` (`hash`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_520_ci;

-- bookmarks_history

DROP TABLE IF EXISTS `bookmarks_history`;
CREATE TABLE `bookmarks_history` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `bookmark_id` char(36) COLLATE utf8mb4_unicode_520_ci NOT NULL,
  `response_status_code` varchar(255) COLLATE utf8mb4_unicode_520_ci NOT NULL,
  `response_reason_phrase` varchar(255) COLLATE utf8mb4_unicode_520_ci NOT NULL,
  `response_headers` longtext COLLATE utf8mb4_unicode_520_ci NOT NULL,
  `request_uri` varchar(255) COLLATE utf8mb4_unicode_520_ci NOT NULL,
  `request_method` varchar(255) COLLATE utf8mb4_unicode_520_ci NOT NULL,
  `request_headers` longtext COLLATE utf8mb4_unicode_520_ci NOT NULL,
  `created_at` datetime NOT NULL,
  PRIMARY KEY (`id`),
  KEY `idx_bookmark_id` (`bookmark_id`),
  CONSTRAINT `bookmarks_history_fk_bookmark_id` FOREIGN KEY (`bookmark_id`) REFERENCES `bookmarks` (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=1655 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_520_ci;

-- feeds

DROP TABLE IF EXISTS `feeds`;
CREATE TABLE `feeds` (
  `id` char(36) COLLATE utf8mb4_unicode_520_ci NOT NULL,
  `url` varchar(255) COLLATE utf8mb4_unicode_520_ci NOT NULL,
  `title` varchar(255) COLLATE utf8mb4_unicode_520_ci NOT NULL,
  `hash` varchar(100) COLLATE utf8mb4_unicode_520_ci NOT NULL,
  `type` varchar(50) COLLATE utf8mb4_unicode_520_ci NOT NULL,
  `status` varchar(50) COLLATE utf8mb4_unicode_520_ci NOT NULL,
  `created_at` datetime NOT NULL,
  `updated_at` datetime NOT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `hash_idx` (`hash`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_520_ci;

-- feeds_history

DROP TABLE IF EXISTS `feeds_history`;
CREATE TABLE `feeds_history` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `feed_id` char(36) COLLATE utf8mb4_unicode_520_ci NOT NULL,
  `response_status_code` varchar(255) COLLATE utf8mb4_unicode_520_ci NOT NULL,
  `response_reason_phrase` varchar(255) COLLATE utf8mb4_unicode_520_ci NOT NULL,
  `response_headers` longtext COLLATE utf8mb4_unicode_520_ci NOT NULL,
  `request_uri` varchar(255) COLLATE utf8mb4_unicode_520_ci NOT NULL,
  `request_method` varchar(255) COLLATE utf8mb4_unicode_520_ci NOT NULL,
  `request_headers` longtext COLLATE utf8mb4_unicode_520_ci NOT NULL,
  `created_at` datetime NOT NULL,
  PRIMARY KEY (`id`),
  KEY `idx_feed_id` (`feed_id`),
  CONSTRAINT `feeds_history_fk_feed_id` FOREIGN KEY (`feed_id`) REFERENCES `feeds` (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_520_ci;

-- users

DROP TABLE IF EXISTS `users`;
CREATE TABLE `users` (
  `id` char(36) COLLATE utf8mb4_unicode_520_ci NOT NULL,
  `username` varchar(255) COLLATE utf8mb4_unicode_520_ci NOT NULL DEFAULT "",
  `firstname` varchar(255) COLLATE utf8mb4_unicode_520_ci NOT NULL DEFAULT "",
  `lastname` varchar(255) COLLATE utf8mb4_unicode_520_ci NOT NULL DEFAULT "",
  `password` varchar(255) COLLATE utf8mb4_unicode_520_ci NOT NULL,
  `status` smallint(6) NOT NULL,
  `created_at` datetime NOT NULL,
  `updated_at` datetime NOT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_520_ci;

-- users_bookmarks

DROP TABLE IF EXISTS `users_bookmarks`;
CREATE TABLE `users_bookmarks` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `user_id` char(36) COLLATE utf8mb4_unicode_520_ci NOT NULL,
  `bookmark_id` char(36) COLLATE utf8mb4_unicode_520_ci NOT NULL,
  `added_at` datetime NOT NULL,
  `accessed_at` datetime DEFAULT NULL,
  `marked_as_read` tinyint(1) NOT NULL,
  `linked` tinyint(1) NOT NULL DEFAULT 1,
  PRIMARY KEY (`id`),
  KEY `idx_user_id` (`user_id`),
  KEY `idx_bookmark_id` (`bookmark_id`),
  CONSTRAINT `users_bookmarks_fk_bookmark_id` FOREIGN KEY (`bookmark_id`) REFERENCES `bookmarks` (`id`),
  CONSTRAINT `users_bookmarks_fk_user_id` FOREIGN KEY (`user_id`) REFERENCES `users` (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=4256 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_520_ci;

-- users_emails

DROP TABLE IF EXISTS `users_emails`;
CREATE TABLE `users_emails` (
  `id` char(36) COLLATE utf8mb4_unicode_520_ci NOT NULL,
  `user_id` char(36) COLLATE utf8mb4_unicode_520_ci NOT NULL,
  `email` varchar(255) COLLATE utf8mb4_unicode_520_ci NOT NULL,
  `primary` tinyint(1) NOT NULL,
  `created_at` datetime NOT NULL,
  `updated_at` datetime NOT NULL,
  PRIMARY KEY (`id`),
  KEY `idx_user_id` (`user_id`),
  CONSTRAINT `users_emails_fk_user_id` FOREIGN KEY (`user_id`) REFERENCES `users` (`id`) ON DELETE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_520_ci;

-- users_feeds

DROP TABLE IF EXISTS `users_feeds`;
CREATE TABLE `users_feeds` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `user_id` char(36) COLLATE utf8mb4_unicode_520_ci NOT NULL,
  `feed_id` char(36) COLLATE utf8mb4_unicode_520_ci NOT NULL,
  `added_at` datetime NOT NULL,
  PRIMARY KEY (`id`),
  KEY `idx_user_id` (`user_id`),
  KEY `idx_feed_id` (`feed_id`),
  CONSTRAINT `users_feeds_fk_feed_id` FOREIGN KEY (`feed_id`) REFERENCES `feeds` (`id`),
  CONSTRAINT `users_feeds_fk_user_id` FOREIGN KEY (`user_id`) REFERENCES `users` (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_520_ci;

-- users_tokens

DROP TABLE IF EXISTS `users_tokens`;
CREATE TABLE `users_tokens` (
  `id` char(36) COLLATE utf8mb4_unicode_520_ci NOT NULL,
  `user_id` char(36) COLLATE utf8mb4_unicode_520_ci NOT NULL,
  `user_agent` text COLLATE utf8mb4_unicode_520_ci NOT NULL,
  `user_ip` text COLLATE utf8mb4_unicode_520_ci NOT NULL,
  `macaroon` text COLLATE utf8mb4_unicode_520_ci NOT NULL,
  `status` varchar(50) COLLATE utf8mb4_unicode_520_ci NOT NULL,
  `created_at` datetime NOT NULL,
  `expires_at` datetime NOT NULL,
  `accessed_at` datetime NOT NULL,
  PRIMARY KEY (`id`),
  KEY `idx_user_id` (`user_id`),
  CONSTRAINT `users_tokens_fk_user_id` FOREIGN KEY (`user_id`) REFERENCES `users` (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_520_ci;

