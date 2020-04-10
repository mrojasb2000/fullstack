CREATE TABLE `users` (
`id` int unsigned AUTO_INCREMENT,
`nickname` longtext,
`email` longtext,
`password` longtext,
`created_at` DATETIME NULL DEFAULT CURRENT_TIMESTAMP,
`updated_at` DATETIME NULL DEFAULT CURRENT_TIMESTAMP, 
PRIMARY KEY (`id`)
)  