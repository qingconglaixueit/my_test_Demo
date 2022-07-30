CREATE TABLE `tenant_info`
(
  `id` varchar(255) NOT NULL COMMENT 'tenant id',
  `name` varchar(255) NOT NULL COMMENT 'tenant name',
  `addr` varchar(255) NOT NULL COMMENT 'tenant addr',
  PRIMARY KEY(`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;
