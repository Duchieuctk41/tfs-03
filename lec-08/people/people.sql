-- DROP DATABASE people;
-- CREATE DATABASE people;
USE people;

CREATE TABLE `classes` (
  `id` varchar(10) PRIMARY KEY,
  `name` varchar(30) NOT NULL,
  `term` varchar(20) NOT NULL,
  `description` text(1000)
);

CREATE TABLE `teachers` (
  `id` int PRIMARY KEY AUTO_INCREMENT,
  `email` varchar(45) UNIQUE NOT NULL,
  `password` varchar(8) NOT NULL,
  `name` varchar(30) NOT NULL,
  `phone` varchar(12),
  `status` boolean DEFAULT true
);

CREATE TABLE `students` (
  `id` int PRIMARY KEY AUTO_INCREMENT,
  `email` varchar(45) UNIQUE NOT NULL,
  `password` varchar(8) NOT NULL,
  `name` varchar(30) NOT NULL,
  `class_id` varchar(10),
  `phone` varchar(12),
  `status` boolean DEFAULT true,
  `score` tinyint(1)
);

ALTER TABLE `students` ADD FOREIGN KEY (`class_id`) REFERENCES `classes` (`id`);

CREATE TABLE `calendars` (
  `id` int PRIMARY KEY AUTO_INCREMENT,
  `teacher_id` int ,
  `class_id` varchar(10) NOT NULL,
  `time_start` datetime NOT NULL,
  `time_end` datetime NOT NULL
);

ALTER TABLE `calendars` ADD FOREIGN KEY (`class_id`) REFERENCES `classes` (`id`);

ALTER TABLE `calendars` ADD FOREIGN KEY (`teacher_id`) REFERENCES `teachers` (`id`);