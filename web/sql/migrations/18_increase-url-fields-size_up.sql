ALTER TABLE `documents` MODIFY `url` varchar(768) COLLATE utf8mb4_unicode_520_ci NOT NULL;
ALTER TABLE `bot_logs` MODIFY `request_uri` varchar(768) COLLATE utf8mb4_unicode_520_ci NOT NULL;
