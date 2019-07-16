DROP TABLE IF EXISTS `subscriptions`;
CREATE TABLE `subscriptions` (
  `user_id` int(11) NOT NULL,
  `source_id` int(11) NOT NULL,
  `subscribed` tinyint(1) NOT NULL DEFAULT 0,
  `created_at` datetime NOT NULL,
  `updated_at` datetime NOT NULL,
  PRIMARY KEY (`user_id`,`source_id`),
  KEY `idx_user_id` (`user_id`),
  KEY `idx_source_id` (`source_id`),
  CONSTRAINT `subscriptions_fk_source_id` FOREIGN KEY (`source_id`) REFERENCES `syndication` (`id`),
  CONSTRAINT `subscriptions_fk_user_id` FOREIGN KEY (`user_id`) REFERENCES `users` (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_520_ci;

INSERT INTO `subscriptions` (user_id, created_at, updated_at, subscribed, source_id) SELECT '1', `created_at`, `updated_at`, `paused`, `id` FROM `syndication`;

UPDATE `subscriptions` SET subscribed = 2 WHERE subscribed = 0;
UPDATE `subscriptions` SET subscribed = 0 WHERE subscribed = 1;
UPDATE `subscriptions` SET subscribed = 1 WHERE subscribed = 2;
