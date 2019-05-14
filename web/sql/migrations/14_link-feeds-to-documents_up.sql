ALTER TABLE `feeds` ADD COLUMN `document_id` int(11) NOT NULL AFTER `id`;
ALTER TABLE `feeds` ADD CONSTRAINT `feeds_fk_document_id` FOREIGN KEY (`document_id`) REFERENCES `documents` (`id`) ON DELETE CASCADE ON UPDATE CASCADE;
