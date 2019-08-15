DROP TABLE IF EXISTS `password_reset`;
CREATE TABLE `password_reset` (
  `token` varchar(64) COLLATE utf8mb4_unicode_520_ci NOT NULL,
  `user_id` int(11) NOT NULL,
  `used` tinyint(1) NOT NULL DEFAULT 0,
  `expired_at` datetime NOT NULL,
  `created_at` datetime NOT NULL,
  `used_at` datetime NULL,
  PRIMARY KEY (`token`,`user_id`),
  UNIQUE KEY `token_idx` (`token`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_520_ci;
