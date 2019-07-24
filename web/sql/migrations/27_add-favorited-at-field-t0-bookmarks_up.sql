ALTER TABLE `bookmarks` ADD COLUMN `favorited_at` DATETIME DEFAULT NULL AFTER `added_at`;
UPDATE `bookmarks` SET `favorited_at` = `updated_at` WHERE favorite = 1;
