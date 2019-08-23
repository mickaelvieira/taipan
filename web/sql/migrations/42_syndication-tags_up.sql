DROP TABLE IF EXISTS `syndication_tags`;
CREATE TABLE `syndication_tags` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `label` varchar(255) COLLATE utf8mb4_unicode_520_ci NOT NULL,
  `created_at` datetime NOT NULL,
  `updated_at` datetime NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_520_ci;

DROP TABLE IF EXISTS `syndication_tags_relation`;
CREATE TABLE `syndication_tags_relation` (
  `source_id` int(11) NOT NULL,
  `tag_id` int(11) NOT NULL,
  PRIMARY KEY (`source_id`, `tag_id`),
  CONSTRAINT `syndication_tags_relation_fk_source_id` FOREIGN KEY (`source_id`) REFERENCES `syndication` (`id`) ON DELETE CASCADE ON UPDATE CASCADE,
  CONSTRAINT `syndication_tags_relation_fk_tag_id` FOREIGN KEY (`tag_id`) REFERENCES `syndication_tags` (`id`) ON DELETE CASCADE ON UPDATE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_520_ci;
