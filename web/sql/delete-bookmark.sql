USE `taipan`;

-- SELECT @id := id from `bookmarks` where url = "" LIMIT 1;
DELETE FROM `bookmarks_history` where bookmark_id = @id;
DELETE FROM `users_bookmarks` where bookmark_id = @id;
DELETE FROM `bookmarks` where id = @id;