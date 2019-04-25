-- I need to remove duplicate entries first
-- select id, url, hash, count(id) as total from bookmarks group by url having total > 1;

ALTER TABLE `bookmarks` ADD CONSTRAINT `url_idx` UNIQUE KEY (`url`);
ALTER TABLE `bookmarks` DROP KEY `hash_idx`;
ALTER TABLE `bookmarks` DROP COLUMN `hash`;