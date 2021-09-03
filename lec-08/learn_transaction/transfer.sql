CREATE DATABASE people;

USE paypay;

CREATE TABLE senders(
user_id varchar(10) PRIMARY KEY,
total float
);

CREATE TABLE receivers(
user_id varchar(10) PRIMARY KEY,
total float
);

SELECT * FROM paypay.senders;
SELECT * FROM paypay.receivers;

UPDATE senders SET total = 1000 WHERE user_id = 'user_A';

DROP TABLE send, receive;

INSERT INTO senders(user_id, total) VALUES('user_A', 1000);

INSERT INTO receivers(user_id, total) VALUES('user_B', 1000);
