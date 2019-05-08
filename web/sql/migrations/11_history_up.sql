RENAME TABLE `bookmarks_history` TO `bot_logs`;
DROP TABLE `feeds_history`;
ALTER TABLE `bot_logs` ADD COLUMN `checksum` BINARY(32) NULL AFTER `id`;
ALTER TABLE `bot_logs` ADD COLUMN `content_type` VARCHAR(100) NOT NULL DEFAULT '' AFTER `checksum`;
ALTER TABLE `bot_logs` MODIFY `bookmark_id` INT(11) NULL;
ALTER TABLE `bot_logs` DROP FOREIGN KEY `bookmarks_history_fk_bookmark_id`;
ALTER TABLE `bot_logs` ADD INDEX `request_uri_idx` (`request_uri`);
UPDATE `bot_logs` SET content_type = "text/html";