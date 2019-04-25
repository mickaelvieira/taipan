USE `taipan`;

UPDATE `bookmarks` SET `title` = "" WHERE `title` IS NULL;
UPDATE `bookmarks` SET `description` = "" WHERE `description` IS NULL;
UPDATE `bookmarks` SET `charset` = "" WHERE `charset` IS NULL;
UPDATE `bookmarks` SET `canonical_url` = "" WHERE `canonical_url` IS NULL;

UPDATE `users` SET `username` = "" WHERE `username` IS NULL;
UPDATE `users` SET `firstname` = "" WHERE `firstname` IS NULL;
UPDATE `users` SET `lastname` = "" WHERE `lastname` IS NULL;
