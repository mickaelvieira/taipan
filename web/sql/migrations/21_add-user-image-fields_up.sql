ALTER TABLE `users` ADD `image_name` VARCHAR(255) COLLATE utf8mb4_unicode_520_ci NOT NULL DEFAULT "" AFTER `status`;
ALTER TABLE `users` ADD `image_width` INT(11) DEFAULT 0 AFTER `image_name`;
ALTER TABLE `users` ADD `image_height` INT(11) DEFAULT 0 AFTER `image_width`;
ALTER TABLE `users` ADD `image_format` VARCHAR(50) COLLATE utf8mb4_unicode_520_ci NOT NULL DEFAULT "" AFTER `image_height`;
