CREATE TABLE `book`
(
  `book_id` varchar(255) NOT NULL COMMENT 'stu_id',
  `name` varchar(255) NOT NULL COMMENT 'name',
  `count` INTEGER (255) NOT NULL COMMENT 'password',
  PRIMARY KEY(`book_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;
