USE `bocchi`;

DROP TABLE IF EXISTS `users`;
DROP TABLE IF EXISTS `codes`;
DROP TABLE IF EXISTS `problems`;
DROP TABLE IF EXISTS `testcases`;

CREATE TABLE `users` (
  `id` varchar(36) NOT NULL,
  `name` varchar(32) NOT NULL,
  `score` bigint NOT NULL DEFAULT 0,
  `created_at` bigint NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` bigint NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `deleted_at` bigint DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

CREATE TABLE `codes` (
  `id` varchar(36) NOT NULL,
  `user_id` varchar(36) NOT NULL,
  `problem_id` varchar(36) NOT NULL,
  `code` text NOT NULL,
  `answer` text DEFAULT NULL,
  `created_at` bigint NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` bigint NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `deleted_at` bigint DEFAULT NULL,
  PRIMARY KEY (`id`, `created_at`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

CREATE TABLE `problems` (
  `id` varchar(36) NOT NULL,
  `creater_id` varchar(36) NOT NULL,
  `score` bigint NOT NULL,
  `title` varchar(64) NOT NULL,
  `created_at` bigint NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` bigint NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `deleted_at` bigint DEFAULT NULL,
  PRIMARY KEY(`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

CREATE TABLE `testcases` (
  `id` varchar(36) NOT NULL,
  `problem_id` varchar(36) NOT NULL,
  `stdin` text NOT NULL,
  `stdout` text NOT NULL,
  PRIMARY KEY(`problem_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;