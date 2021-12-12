create table banner
(
    banner_id   bigint auto_increment comment 'BannerID' primary key,
    banner_name varchar(255) default '' not null comment 'Banner名称',
) comment 'Banner表';