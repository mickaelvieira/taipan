ALTER TABLE `bookmarks` DROP KEY `url_idx`;
ALTER TABLE `bookmarks` ADD `hash` VARCHAR(100) COLLATE utf8mb4_unicode_520_ci NOT NULL AFTER `url`;
ALTER TABLE `bookmarks` ADD CONSTRAINT `hash_idx` UNIQUE KEY (`hash`);
UPDATE `bookmarks` SET `hash` = SHA2(`url`, 256);
