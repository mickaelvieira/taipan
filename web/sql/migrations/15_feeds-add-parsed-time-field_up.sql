ALTER TABLE `feeds` ADD COLUMN `parsed_at` datetime NULL AFTER `updated_at`;


ALTER TABLE `bookmarks` CHANGE COLUMN `marked_as_read` `marked_as_read` TINYINT(1) NOT NULL DEFAULT 0;