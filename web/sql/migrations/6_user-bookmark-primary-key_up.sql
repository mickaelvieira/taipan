ALTER TABLE `users_bookmarks` DROP PRIMARY KEY, MODIFY `id` INT(11) NULL;
ALTER TABLE `users_bookmarks` ADD PRIMARY KEY (`user_id`, `bookmark_id`);
ALTER TABLE `users_bookmarks` DROP COLUMN `id`;
