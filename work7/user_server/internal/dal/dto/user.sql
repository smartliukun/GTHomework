use test;

CREATE TABLE `users` (
 `id` int(11) NOT NULL AUTO_INCREMENT,
 `name` varchar(255) DEFAULT NULL,
 `email` varchar(255)  DEFAULT NULL,
 `age` int(11) DEFAULT NULL,
 `birthday` datetime DEFAULT NULL,
 PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=3 DEFAULT CHARSET=utf8mb4;


INSERT INTO `test`.`users` (`id`, `name`, `email`, `age`, `birthday`) VALUES ('1', '张三', 'zhangsan@qq.com', '18', '2003-10-24');
INSERT INTO `test`.`users` (`id`, `name`, `email`, `age`, `birthday`) VALUES ('2', '李四', 'lisi@qq.com', '17', '2004-10-10');
