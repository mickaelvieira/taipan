ALTER TABLE `bot_logs` ADD COLUMN `failed` TINYINT(1) NOT NULL DEFAULT 0 AFTER `id`;
ALTER TABLE `bot_logs` ADD COLUMN `failure_reason` VARCHAR(255) COLLATE utf8mb4_unicode_520_ci NOT NULL DEFAULT "" AFTER `failed`;
