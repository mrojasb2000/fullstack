CREATE TABLE `posts` (
`id` bigint unsigned AUTO_INCREMENT,
`title` varchar(255) NOT NULL UNIQUE,
`content` varchar(255) NOT NULL,`author_id` int unsigned NOT NULL,
`created_at` DATETIME NULL DEFAULT CURRENT_TIMESTAMP,
`updated_at` DATETIME NULL DEFAULT CURRENT_TIMESTAMP, 
PRIMARY KEY (`id`)
)