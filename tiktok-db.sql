# 用户表
drop table if exists users;
create table users(
    id int not null auto_increment comment '自增主键',
    name varchar(255) not null comment '用户名',
    password varchar(255) not null comment '用户密码',
    primary key(id)
) engine=InnoDB default charset=utf8 comment '用户表';
# 测试用户：账号test，密码test123123
insert into users(name,password) values('test','bebd18a77a23f20aab5efdd643ac100a5f1fc1a267d6b15add5b231bf25a75e8');

# follows 关注表
drop table if exists follows;
create table follows(
    id int not null auto_increment comment '自增主键',
    user_id int not null comment '当前用户Id',
    focus_id int not null comment '关注者的Id',
    primary key(id)
) engine=InnoDB default charset=utf8 comment '关注表';

# followers 粉丝表
drop table if exists followers;
create table followers(
    id int not null auto_increment comment '自增主键',
    user_id int not null comment '当前用户Id',
    fans_id int not null comment '粉丝的Id',
    primary key(id)
) engine=InnoDB default charset=utf8 comment '粉丝表';
