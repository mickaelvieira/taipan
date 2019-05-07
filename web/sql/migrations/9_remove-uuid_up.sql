-- Drop all foreign keys

ALTER TABLE `feeds_history` DROP FOREIGN KEY `feeds_history_fk_feed_id`;
ALTER TABLE `bookmarks_history` DROP FOREIGN KEY `bookmarks_history_fk_bookmark_id`;
ALTER TABLE `users_bookmarks` DROP FOREIGN KEY `users_bookmarks_fk_bookmark_id`;
ALTER TABLE `users_bookmarks` DROP FOREIGN KEY `users_bookmarks_fk_user_id`;
ALTER TABLE `users_emails` DROP FOREIGN KEY `users_emails_fk_user_id`;
ALTER TABLE `users_feeds` DROP FOREIGN KEY `users_feeds_fk_feed_id`;
ALTER TABLE `users_feeds` DROP FOREIGN KEY `users_feeds_fk_user_id`;
ALTER TABLE `users_tokens` DROP FOREIGN KEY `users_tokens_fk_user_id`;

-- Drop unnecessary tables

DROP TABLE IF EXISTS `users_tokens`;
DROP TABLE IF EXISTS `users_emails`;

-- Bookmarks and Users UUIDs

ALTER TABLE `bookmarks` DROP PRIMARY KEY;
ALTER TABLE `bookmarks` ADD `old_id` CHAR(36) COLLATE utf8mb4_unicode_520_ci NOT NULL;
UPDATE `bookmarks` SET `old_id` = `id`;
ALTER TABLE `bookmarks` DROP COLUMN `id`;
ALTER TABLE `bookmarks` ADD `id` INT(11) NOT NULL AUTO_INCREMENT FIRST, ADD PRIMARY KEY (`id`);

ALTER TABLE `users` DROP PRIMARY KEY;
ALTER TABLE `users` ADD `old_id` CHAR(36) COLLATE utf8mb4_unicode_520_ci NOT NULL;
UPDATE `users` SET `old_id` = `id`;
ALTER TABLE `users` DROP COLUMN `id`;
ALTER TABLE `users` ADD `id` INT(11) NOT NULL AUTO_INCREMENT FIRST, ADD PRIMARY KEY (`id`);

ALTER TABLE `users_bookmarks` DROP PRIMARY KEY;
ALTER TABLE `users_bookmarks` ADD `old_user_id` CHAR(36) COLLATE utf8mb4_unicode_520_ci NOT NULL;
ALTER TABLE `users_bookmarks` ADD `old_bookmark_id` CHAR(36) COLLATE utf8mb4_unicode_520_ci NOT NULL;
UPDATE `users_bookmarks` SET `old_user_id` = `user_id`;
UPDATE `users_bookmarks` SET `old_bookmark_id` = `bookmark_id`;

ALTER TABLE `users_bookmarks` DROP COLUMN `user_id`;
ALTER TABLE `users_bookmarks` DROP COLUMN `bookmark_id`;
ALTER TABLE `users_bookmarks` ADD `user_id` INT(11) NOT NULL FIRST;
ALTER TABLE `users_bookmarks` ADD `bookmark_id` INT(11) NOT NULL AFTER `user_id`;

UPDATE `users_bookmarks` INNER JOIN `users` ON `users`.`old_id` = `users_bookmarks`.`old_user_id` SET `user_id` = `id`;
UPDATE `users_bookmarks` INNER JOIN `bookmarks` ON `bookmarks`.`old_id` = `users_bookmarks`.`old_bookmark_id` SET `bookmark_id` = `id`;

ALTER TABLE `users_bookmarks` ADD PRIMARY KEY (`user_id`, `bookmark_id`);
ALTER TABLE `users_bookmarks` ADD CONSTRAINT `users_bookmarks_fk_bookmark_id` FOREIGN KEY (`bookmark_id`) REFERENCES `bookmarks` (`id`) ON DELETE CASCADE ON UPDATE CASCADE;
ALTER TABLE `users_bookmarks` ADD CONSTRAINT `users_bookmarks_fk_user_id` FOREIGN KEY (`user_id`) REFERENCES `users` (`id`) ON DELETE CASCADE ON UPDATE CASCADE;

ALTER TABLE `bookmarks_history` ADD `old_bookmark_id` CHAR(36) COLLATE utf8mb4_unicode_520_ci NOT NULL;
UPDATE `bookmarks_history` SET `old_bookmark_id` = `bookmark_id`;
ALTER TABLE `bookmarks_history` DROP COLUMN `bookmark_id`;
ALTER TABLE `bookmarks_history` ADD `bookmark_id` INT(11) NOT NULL AFTER `id`;
UPDATE `bookmarks_history` INNER JOIN `bookmarks` ON `bookmarks`.`old_id` = `bookmarks_history`.`old_bookmark_id` SET `bookmark_id` = `bookmarks`.`id`;
ALTER TABLE `bookmarks_history` ADD CONSTRAINT `bookmarks_history_fk_bookmark_id` FOREIGN KEY (`bookmark_id`) REFERENCES `bookmarks` (`id`) ON DELETE CASCADE ON UPDATE CASCADE;

-- Feeds and Users UUIDs

ALTER TABLE `feeds` DROP PRIMARY KEY;
ALTER TABLE `feeds` ADD `old_id` CHAR(36) COLLATE utf8mb4_unicode_520_ci NOT NULL;
UPDATE `feeds` SET `old_id` = `id`;
ALTER TABLE `feeds` DROP COLUMN `id`;
ALTER TABLE `feeds` ADD `id` INT(11) NOT NULL AUTO_INCREMENT FIRST, ADD PRIMARY KEY (`id`);

DROP TABLE IF EXISTS `users_feeds`;
CREATE TABLE `users_feeds` (
  `user_id` INT(11) NOT NULL,
  `feed_id` INT(11) NOT NULL,
  `added_at` datetime NOT NULL,
  PRIMARY KEY (`user_id`, `feed_id`),
  KEY `idx_user_id` (`user_id`),
  KEY `idx_feed_id` (`feed_id`),
  CONSTRAINT `users_feeds_fk_feed_id` FOREIGN KEY (`feed_id`) REFERENCES `feeds` (`id`) ON DELETE CASCADE ON UPDATE CASCADE,
  CONSTRAINT `users_feeds_fk_user_id` FOREIGN KEY (`user_id`) REFERENCES `users` (`id`) ON DELETE CASCADE ON UPDATE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_520_ci;

INSERT INTO `users_feeds` (`user_id`, `feed_id`, `added_at`) SELECT 1, id, created_at FROM `feeds`;

DROP TABLE IF EXISTS `feeds_history`;
CREATE TABLE `feeds_history` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `feed_id` INT(11) NOT NULL,
  `response_status_code` varchar(255) COLLATE utf8mb4_unicode_520_ci NOT NULL,
  `response_reason_phrase` varchar(255) COLLATE utf8mb4_unicode_520_ci NOT NULL,
  `response_headers` longtext COLLATE utf8mb4_unicode_520_ci NOT NULL,
  `request_uri` varchar(255) COLLATE utf8mb4_unicode_520_ci NOT NULL,
  `request_method` varchar(255) COLLATE utf8mb4_unicode_520_ci NOT NULL,
  `request_headers` longtext COLLATE utf8mb4_unicode_520_ci NOT NULL,
  `created_at` datetime NOT NULL,
  PRIMARY KEY (`id`),
  KEY `idx_feed_id` (`feed_id`),
  CONSTRAINT `feeds_history_fk_feed_id` FOREIGN KEY (`feed_id`) REFERENCES `feeds` (`id`) ON DELETE CASCADE ON UPDATE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_520_ci;
