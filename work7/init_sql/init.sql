set names utf8;

create database test_db;
user test_db;

create table product
(
    product_id int auto_increment comment '主键'
        primary key,
    name       varchar(255) null comment '商品名称',
    price      bigint       null comment '价格',
    stock      int          null comment '库存'
)
    comment '商品表' charset = utf8mb4;

INSERT INTO test_db.product (product_id, name, price, stock) VALUES (1, '衣服', 1000, 100);
INSERT INTO test_db.product (product_id, name, price, stock) VALUES (2, '玩具', 2000, 100);

create table trade
(
    trade_no    int auto_increment comment '交易编号'
        primary key,
    user_id     int    null comment '用户主键',
    product_id  int    null comment '商品主键',
    product_num int    null comment '商品数量',
    price       int    null comment '商品价格',
    cost        bigint null comment '商品成交价格',
    Date        date   null comment '成交日期'
)
    comment '交易表' charset = utf8mb4;

INSERT INTO test_db.trade (trade_no, user_id, product_id, product_num, price, cost, Date) VALUES (1, 1, 1, 1, 1000, 1000, '2022-03-13');


create table users
(
    id       int auto_increment comment '主键'
        primary key,
    name     varchar(255) null comment '用户姓名',
    email    varchar(255) null comment '邮箱',
    age      int          null comment '年龄',
    birthday datetime     null comment '生日'
)
    comment '用户表' charset = utf8mb4;

INSERT INTO test_db.users (id, name, email, age, birthday) VALUES (1, '张三', 'zhangsan@qq.com', 18, '2003-10-24 00:00:00');
INSERT INTO test_db.users (id, name, email, age, birthday) VALUES (2, '李四', 'lisi@qq.com', 17, '2004-10-10 00:00:00');