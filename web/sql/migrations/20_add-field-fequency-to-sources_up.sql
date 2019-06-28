ALTER TABLE `feeds` ADD COLUMN `frequency` VARCHAR(10) NOT NULL DEFAULT "hourly" AFTER `paused`;
