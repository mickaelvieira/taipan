ALTER TABLE `bookmarks` MODIFY `old_id` CHAR(36) COLLATE utf8mb4_unicode_520_ci NOT NULL;
ALTER TABLE `feeds` MODIFY `old_id` CHAR(36) COLLATE utf8mb4_unicode_520_ci NOT NULL;
ALTER TABLE `users` MODIFY `old_id` CHAR(36) COLLATE utf8mb4_unicode_520_ci NOT NULL;
ALTER TABLE `users_bookmarks` MODIFY `old_bookmark_id` CHAR(36) COLLATE utf8mb4_unicode_520_ci NOT NULL;
ALTER TABLE `bookmarks_history` MODIFY `old_bookmark_id` CHAR(36) COLLATE utf8mb4_unicode_520_ci NOT NULL;
ALTER TABLE `users_bookmarks` MODIFY `old_user_id` CHAR(36) COLLATE utf8mb4_unicode_520_ci NOT NULL;
