CREATE TABLE `user`
(
  `stu_id` varchar(255) NOT NULL COMMENT 'stu_id',
  `name` varchar(255) NOT NULL COMMENT 'name',
  `password` varchar(255) NOT NULL COMMENT 'password',
  `gender` varchar(255) NOT NULL COMMENT 'gender',
  PRIMARY KEY(`stu_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;
