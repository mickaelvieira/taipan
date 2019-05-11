RENAME TABLE `bookmarks` TO `documents`;
RENAME TABLE `users_bookmarks` TO `bookmarks`;

ALTER TABLE `bookmarks` CHANGE COLUMN `bookmark_id` `document_id` int(11) NOT NULL;
