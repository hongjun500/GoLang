recordings

create table if not exists album
(
    id     int auto_increment
    primary key,
    title  varchar(128)  not null,
    artist varchar(255)  not null,
    price  decimal(5, 2) not null
    );

create table if not exists tag
(
    id          int unsigned auto_increment
    primary key,
    name        varchar(100)     default ''  null comment '标签名称',
    created_on  datetime                     null comment '创建时间',
    created_by  varchar(100)     default ''  null comment '创建人',
    modified_on datetime                     null comment '修改时间',
    modified_by varchar(100)     default ''  null comment '修改人',
    deleted_on  datetime                     null,
    state       tinyint unsigned default '1' null comment '状态 0为禁用、1为启用'
    )
    comment '文章标签管理' charset = utf8;

