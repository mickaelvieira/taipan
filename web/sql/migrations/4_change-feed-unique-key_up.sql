-- I need to remove duplicate entries first
-- select id, url, hash, count(id) as total from feeds group by url having total > 1;

ALTER TABLE `feeds` ADD CONSTRAINT `url_idx` UNIQUE KEY (`url`);
ALTER TABLE `feeds` DROP KEY `hash_idx`;
ALTER TABLE `feeds` DROP COLUMN `hash`;