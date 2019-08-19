DROP TABLE IF EXISTS `users_emails_confirm`;
CREATE TABLE `users_emails_confirm` (
  `token` varchar(64) COLLATE utf8mb4_unicode_520_ci NOT NULL,
  `user_id` int(11) NOT NULL,
  `email_id` int(11) NOT NULL,
  `used` tinyint(1) NOT NULL DEFAULT 0,
  `expired_at` datetime NOT NULL,
  `created_at` datetime NOT NULL,
  `used_at` datetime NULL,
  PRIMARY KEY (`token`, `user_id`, `email_id`),
  UNIQUE KEY `token_idx` (`token`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_520_ci;

ALTER TABLE `users_emails` ADD COLUMN `confirmed_at` datetime NULL;
