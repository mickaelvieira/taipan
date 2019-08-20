ALTER TABLE `password_reset` ADD CONSTRAINT `password_reset_fk_user_id` FOREIGN KEY (`user_id`) REFERENCES `users` (`id`) ON DELETE CASCADE ON UPDATE CASCADE;
