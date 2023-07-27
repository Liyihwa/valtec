DROP TABLE IF EXISTS `HERO`;

CREATE TABLE `HERO` (
                        `id` INT PRIMARY KEY AUTO_INCREMENT,
                        `name` VARCHAR(31),
                        `avatar` VARCHAR(255),
                        `c` VARCHAR(255),
                        `q` VARCHAR(255),
                        `e` VARCHAR(255),
                        `x` VARCHAR(255)
);

DROP TABLE IF EXISTS `MAP`;
CREATE TABLE `MAP` (
                       `id` INT PRIMARY KEY AUTO_INCREMENT,
                       `name` VARCHAR(31),
                       `url` VARCHAR(255),
                       `avatar` VARCHAR(255)
);


DROP TABLE IF EXISTS `POSITION`;
CREATE TABLE `POSITION` (
                            `id` INT PRIMARY KEY AUTO_INCREMENT,
                            `deleted_at` TIMESTAMP NULL,
                            `hero_id` INT,
                            `skill` VARCHAR(255),
                            `map_id` INT,
                            `stand_x` FLOAT,
                            `stand_y` FLOAT,
                            `put_x` FLOAT,
                            `put_y` FLOAT,
                            `like` INT,
                            `dislike` INT,
                            `description` TEXT
);


DROP TABLE IF EXISTS `USER`;
CREATE TABLE `USER` (
                        `id` INT PRIMARY KEY AUTO_INCREMENT,
                        `password` BINARY(32),
                        `email` VARCHAR(255),
                        `captcha` CHAR(32)
);

SELECT * FROM `MAP`;