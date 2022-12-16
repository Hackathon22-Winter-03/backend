USE `backend+hackathon22-winter-03`;

DROP TABLE IF EXISTS `users`;
DROP TABLE IF EXISTS `codes`;
DROP TABLE IF EXISTS `problems`;
DROP TABLE IF EXISTS `testcases`;

CREATE TABLE `users` (
  `id` varchar(36) NOT NULL,
  `name` varchar(32) NOT NULL,
  `comment` text NOT NULL,
  `score` bigint NOT NULL DEFAULT 0,
  `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `deleted_at` datetime DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

CREATE TABLE `codes` (
  `id` varchar(36) NOT NULL,
  `user_id` varchar(36) NOT NULL,
  `problem_id` varchar(36) NOT NULL,
  `code` text NOT NULL,
  `answer` text DEFAULT NULL,
  `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `deleted_at` datetime DEFAULT NULL,
  PRIMARY KEY (`id`, `created_at`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

CREATE TABLE `problems` (
  `id` varchar(36) NOT NULL,
  `creator_id` varchar(36) NOT NULL,
  `score` bigint NOT NULL,
  `title` varchar(64) NOT NULL,
  `text` text NOT NULL,
  `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `deleted_at` datetime DEFAULT NULL,
  PRIMARY KEY(`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

CREATE TABLE `testcases` (
  `id` varchar(36) NOT NULL,
  `problem_id` varchar(36) NOT NULL,
  `stdin` text NOT NULL,
  `stdout` text NOT NULL,
  PRIMARY KEY(`problem_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;