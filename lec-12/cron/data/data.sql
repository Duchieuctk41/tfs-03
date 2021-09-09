CREATE DATABASE meow;
USE meow;
CREATE TABLE `order` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
  `created_at` datetime NOT NULL,
  `email` text CHARACTER SET latin1,
  `customer_name` varchar(100) DEFAULT NULL,
  `thankyou_email_sent` tinyint(1),
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=0 DEFAULT CHARSET=utf8 COLLATE=utf8_unicode_ci;

INSERT INTO meow.order(created_at, email, customer_name, thankyou_email_sent) 
VALUES("2021-09-09 16:05:00", "duchieu@gmail.com", "hieu hoc code", 0);

INSERT INTO meow.order(created_at, email, customer_name, thankyou_email_sent)
VALUES("2021-09-09 16:09:00", "hieudeptrai@gmail.com", "khong ten", 0);

INSERT INTO meow.order(created_at, email, customer_name, thankyou_email_sent)
VALUES("2021-09-09 16:010:00", "hieuhoccode@gmail.com", "no name", 0);

SELECT * FROM meow.order;