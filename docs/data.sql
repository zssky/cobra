CREATE TABLE `template` (
`id` int(11) PRIMARY KEY AUTOINCREMENT,
`name` varchar(32) NOT NULL,
`type` INTEGER NOT NULL DEFAULT 0 ,
`content` TEXT NOT NULL,
`remark` TEXT DEFAULT NULL
);
