# 用户表
drop table if exists users;
create table users(
    id int not null auto_increment comment '自增主键',
    name varchar(255) not null comment '用户名',
    password varchar(255) not null comment '用户密码',
    primary key(id)
) engine=InnoDB default charset=utf8 comment '用户表';
# 测试用户：账号aaa，密码123456; 账号bbb，密码123456; 账号ccc，密码123456
insert into users(name,password) values('aaa','27f944594a18380df0bc6c26e678eb478d7c336edad4c17067e413c3b13cc3d2');
insert into users(name,password) values('bbb','27f944594a18380df0bc6c26e678eb478d7c336edad4c17067e413c3b13cc3d2');
insert into users(name,password) values('ccc','27f944594a18380df0bc6c26e678eb478d7c336edad4c17067e413c3b13cc3d2');

# follows 关注表
drop table if exists follows;
create table follows(
    id int not null auto_increment comment '自增主键',
    user_id int not null comment '当前用户Id',
    focus_id int not null comment '关注者的Id',
    primary key(id)
) engine=InnoDB default charset=utf8 comment '关注表';
# 测试关注：aaa关注ccc; bbb关注ccc;
insert into follows(user_id,focus_id) values(1,3);
insert into follows(user_id,focus_id) values(2,3);

# followers 粉丝表
drop table if exists followers;
create table followers(
    id int not null auto_increment comment '自增主键',
    user_id int not null comment '当前用户Id',
    fans_id int not null comment '粉丝的Id',
    primary key(id)
) engine=InnoDB default charset=utf8 comment '粉丝表';
# 测试粉丝：ccc的粉丝有aaa和bbb;
insert into followers(user_id,fans_id) values(3,1);
insert into followers(user_id,fans_id) values(3,2);