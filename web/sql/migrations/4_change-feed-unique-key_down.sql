ALTER TABLE `feeds` DROP KEY `url_idx`;
ALTER TABLE `feeds` ADD `hash` VARCHAR(100) COLLATE utf8mb4_unicode_520_ci NOT NULL AFTER `url`;
ALTER TABLE `feeds` ADD CONSTRAINT `hash_idx` UNIQUE KEY (`hash`);
UPDATE `feeds` SET `hash` = SHA2(`url`, 256);
