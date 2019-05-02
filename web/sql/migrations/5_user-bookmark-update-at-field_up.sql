ALTER TABLE `users_bookmarks` ADD `updated_at` datetime NOT NULL AFTER `added_at`;
UPDATE `users_bookmarks` SET `updated_at` = `added_at`;