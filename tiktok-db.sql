# users 用户表
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
    focus_id int not null comment '关注的Id',
    primary key(id)
) engine=InnoDB default charset=utf8 comment '关注表';
# 测试关注：aaa关注ccc; bbb关注ccc;
insert into follows(user_id,focus_id) values(1,3);
insert into follows(user_id,focus_id) values(2,3);

# videos 视频表
drop table if exists videos;
create table videos(
    id int not null auto_increment comment '自增主键',
    user_id int not null comment '用户Id',
    title varchar(255) not null comment '视频标题',
    video_url varchar(255) not null comment '视频地址',
    cover_url varchar(255) not null comment '封面地址',
    publish_time datetime not null comment '发布时间',
    primary key(id)
) engine=InnoDB default charset=utf8 comment '视频表';
# 测试视频1：aaa发布了"测试视频1",
# 视频地址为："https://api.fifcom.cn/vid/1.mp4",
# 封面地址为："https://www.fifcom.cn/avatar/",
# 发布时间为："2022-01-01 00:00:00";
insert into videos(user_id,title,video_url,cover_url,publish_time)
values(1,'测试视频1','1.mp4','img/raw/512.png','2022-01-01 00:00:00');
# 测试视频2：bbb发布了"测试视频2",
# 视频地址为："https://api.fifcom.cn/vid/2.mp4",
# 封面地址为："https://www.fifcom.cn/avatar/",
# 发布时间为："2022-02-01 00:00:00";
insert into videos(user_id,title,video_url,cover_url,publish_time)
values(2,'测试视频2','2.mp4','img/raw/512.png','2022-02-01 00:00:00');
# 测试视频3：ccc发布了"测试视频3",
# 视频地址为："https://api.fifcom.cn/vid/3.mp4",
# 封面地址为："https://www.fifcom.cn/avatar/",
# 发布时间为："2022-03-01 00:00:00";
insert into videos(user_id,title,video_url,cover_url,publish_time)
values(3,'测试视频3','3.mp4','img/raw/512.png','2022-03-01 00:00:00');
# 测试视频4：aaa发布了"测试视频4",
# 视频地址为："https://api.fifcom.cn/vid/4.mp4",
# 封面地址为："https://www.fifcom.cn/avatar/",
# 发布时间为："2022-04-01 00:00:00";
insert into videos(user_id,title,video_url,cover_url,publish_time)
values(1,'测试视频4','4.mp4','img/raw/512.png','2022-04-01 00:00:00');