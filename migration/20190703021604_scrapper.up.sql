CREATE DATABASE IF NOT EXISTS `scrapper`;

USE `scrapper`;

DROP TABLE IF EXISTS `archieve`;

CREATE TABLE IF NOT EXISTS `archieve` (
  `archieve_id` char(36) NOT NULL,
  `code` varchar(255) NOT NULL,
  `link` varchar(255) NOT NULL,
  `published` varchar(255) NOT NULL,
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,

  PRIMARY KEY(`archieve_id`),

  CONSTRAINT `UN01_ARCHIEVE`
    UNIQUE KEY `UN_code`(`code`)
);

DROP TABLE IF EXISTS `journal`;

CREATE TABLE IF NOT EXISTS `journal` (
  `journal_id` char(36) NOT NULL,
  `archieve_id` char(36) NOT NULL,
  `title` varchar(255) NOT NULL,
  `authors` varchar(255) NOT NULL,
  `abstract` text NOT NULL,
  `link` varchar(255) NOT NULL,
  `pdf_link` varchar(255) NOT NULL,
  `published` varchar(255) NOT NULL,
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,

  PRIMARY KEY(`journal_id`),

  CONSTRAINT `FK01_JOURNAL`
    FOREIGN KEY `fk_archieve_id`(`archieve_id`)
    REFERENCES `archieve`(`archieve_id`)
    ON DELETE CASCADE
    ON UPDATE CASCADE
    
);

