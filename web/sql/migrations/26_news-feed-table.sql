DROP TABLE IF EXISTS `newsfeed`;
CREATE TABLE `newsfeed` (
  `user_id` int(11) NOT NULL,
  `document_id` int(11) NOT NULL,
  `created_at` datetime NOT NULL,
  PRIMARY KEY (`user_id`,`document_id`),
  KEY `idx_user_id` (`user_id`),
  KEY `idx_document_id` (`document_id`),
  CONSTRAINT `subscriptions_fk_document_id` FOREIGN KEY (`document_id`) REFERENCES `documents` (`id`),
  CONSTRAINT `subscriptions_fk_user_id` FOREIGN KEY (`user_id`) REFERENCES `users` (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_520_ci;
