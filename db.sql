create database DouYin;
use DouYin;
DROP TABLE IF EXISTS `users`;
CREATE TABLE `users` (
                         `id` int unsigned NOT NULL AUTO_INCREMENT,
                         `created_at` datetime DEFAULT current_timestamp,
                         `updated_at` datetime On update current_timestamp default NULL,
                         `extension` varchar(255) DEFAULT NULL,
                         `username` varchar(128) NOT NULL,
                         `password` varchar(64) NOT NULL,
                         PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8mb3;