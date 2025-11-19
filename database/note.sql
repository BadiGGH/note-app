DROP TABLE IF EXISTS note;
CREATE TABLE note (
  id          INT AUTO_INCREMENT NOT NULL,
  title       VARCHAR(128) NOT NULL,
  author      VARCHAR(128) NOT NULL,
  body        LONGTEXT NOT NULL,
  created_at  DATETIME NOT NULL,
  modified_at DATETIME,
  PRIMARY KEY (`id`)
);