CREATE TABLE `Users`
(
    id   bigint auto_increment,
    name varchar(255) NOT NULL,
    PRIMARY KEY (`id`)
);

INSERT INTO `Users` (`name`)
VALUES ('abc'),('def'),('xyz');