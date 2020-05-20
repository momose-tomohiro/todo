--todoÉeÅ[ÉuÉãçÏê¨
DROP TABLE IF EXISTS `todo`.`trn_todo`;
CREATE TABLE `todo`.`trn_todo` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
  `name` varchar(255) COLLATE utf8mb4_bin NULL DEFAULT NULL,
  `priority` varchar(1) COLLATE utf8mb4_bin NULL DEFAULT NULL,
  `time_limit` varchar(255) COLLATE utf8mb4_bin NULL DEFAULT NULL,
  `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `deleted_at` datetime NULL DEFAULT NULL,
PRIMARY KEY (`id`)
)ENGINE="InnoDB" DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin;
