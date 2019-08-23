ALTER TABLE `syndication` ADD FULLTEXT INDEX `search_syndication_idx` (`title`);
ALTER TABLE `documents` ADD FULLTEXT INDEX `search_documents_idx` (`title`, `description`);
