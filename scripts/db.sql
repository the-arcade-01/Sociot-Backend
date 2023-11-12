-- ------------------
-- Database `josu`
-- ------------------

CREATE DATABASE IF NOT EXISTS `sociot`;

-- ------------------
-- Table `users`
-- ------------------

CREATE TABLE IF NOT EXISTS `sociot`.`users` (
    `userId` INT NOT NULL AUTO_INCREMENT PRIMARY KEY,
    `userName` VARCHAR(255) NOT NULL UNIQUE,
    `email` VARCHAR(255) NOT NULL UNIQUE,
    `password` VARCHAR(255) NOT NULL,
    `createdAt` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `updatedAt` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
);

-- ------------------
-- Table `posts`
-- ------------------

CREATE TABLE IF NOT EXISTS `sociot`.`posts` (
    `userId` INT NOT NULL,
    `postId` INT NOT NULL AUTO_INCREMENT PRIMARY KEY,
    `content` VARCHAR(500) NOT NULL,
    `createdAt` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `updatedAt` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    FOREIGN KEY (userId) REFERENCES `sociot`.`users`(userId)
)