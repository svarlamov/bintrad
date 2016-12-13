CREATE DATABASE IF NOT EXISTS `bintraddev`;
USE bintraddev;
CREATE TABLE IF NOT EXISTS `user` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `username` varchar(255) NOT NULL,
  `first_name` varchar(255) NOT NULL,
  `last_name` varchar(255) NOT NULL,
  `middle_name` varchar(255) DEFAULT "",
  `grade` int DEFAULT -1,
  `gender` char(1),
  `passkey` varchar(255) DEFAULT "",
  `passkey_salt` varchar(255) DEFAULT "",
  `student_id` BIGINT(16) DEFAULT -1,
  `email` varchar(255) NOT NULL UNIQUE,
  `starting_balance` double NOT NULL,
  `created_at` datetime,
  `updated_at` datetime,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;
CREATE TABLE IF NOT EXISTS `access_token` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `user_id` bigint(20) NOT NULL,
  `token` varchar(255) NOT NULL,
  `created_at` datetime,
  `updated_at` datetime,
  PRIMARY KEY (`id`),
  FOREIGN KEY (`user_id`) REFERENCES user(id)
) ENGINE=InnoDB DEFAULT CHARSET=latin1;
CREATE TABLE IF NOT EXISTS `ticker` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `ticker` varchar(255) NOT NULL,
  `name` varchar(255) NOT NULL,
  `type` varchar(255) NOT NULL,
  `created_at` datetime,
  `updated_at` datetime,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;
CREATE TABLE IF NOT EXISTS `ticker_data` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `ticker_id` bigint(20) NOT NULL,
  `opens_at` datetime NOT NULL,
  `closes_at` datetime NOT NULL,
  `period` int NOT NULL,
  `open` double NOT NULL,
  `high` double NOT NULL,
  `low` double NOT NULL,
  `close` double NOT NULL,
  `volume` double NOT NULL,
  `created_at` datetime,
  `updated_at` datetime,
  PRIMARY KEY (`id`),
  FOREIGN KEY (`ticker_id`) REFERENCES ticker(id)
) ENGINE=InnoDB DEFAULT CHARSET=latin1;
CREATE TABLE IF NOT EXISTS `contract_session` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `user_id` bigint(20) NOT NULL,
  `ticker_id` bigint(20) NOT NULL,
  `price` double NOT NULL,
  `ttl` int NOT NULL,
  `period` int NOT NULL,
  `data_start` datetime,
  `data_end` datetime,
  `final_tick_id` bigint(20) NOT NULL,
  `created_at` datetime,
  `is_closed` boolean,
  `closed_at` datetime,
  `updated_at` datetime,
  PRIMARY KEY (`id`),
  FOREIGN KEY (`user_id`) REFERENCES user(id),
  FOREIGN KEY (`final_tick_id`) REFERENCES ticker_data(id),
  FOREIGN KEY (`ticker_id`) REFERENCES ticker(id)
) ENGINE=InnoDB DEFAULT CHARSET=latin1;
CREATE TABLE IF NOT EXISTS `contract` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `user_id` bigint(20) NOT NULL,
  `contract_session_id` bigint(20) NOT NULL,
  `ticker_id` bigint(20) NOT NULL,
  `bet` double NOT NULL,
  `price` double NOT NULL,
  `is_bullish` boolean NOT NULL,
  `is_correct` boolean NOT NULL,
  `return` double NOT NULL,
  `created_at` datetime,
  `updated_at` datetime,
  PRIMARY KEY (`id`),
  FOREIGN KEY (`user_id`) REFERENCES user(id),
  FOREIGN KEY (`contract_session_id`) REFERENCES contract_session(id),
  FOREIGN KEY (`ticker_id`) REFERENCES ticker(id)
) ENGINE=InnoDB DEFAULT CHARSET=latin1;
