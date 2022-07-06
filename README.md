# go-tiktok-lite
 Simple version of tiktok backend implemented in golang

# APIs

| Method | Path | Description |
| :----- | ---- | ----------- |
| GET    | /douyin/feed | 无需登录，返回视频列表 |
| POST   | /douyin/user/register | 注册接口，提供用户名和密码 |
| POST   | /douyin/user/login | 登录接口，提供用户名和密码，返回用户id和token |
| POST   | /douyin/publish/action | 视频投稿 |
| GET   | /douyin/publish/list | 登录用户发布的视频列表 |
| POST | /douyin/relation/action | 登陆用户或其他用户进行关注或取消关注|
| GET | /douyin/relation/follow/list | 登陆用户关注的所有用户列表 |
| GET | /douyin/relation/follower/list | 所有关注登陆用户的粉丝列表 |
| POST | /douyin/favorite/action/ | 登录用户对视频的点赞和取消点赞操作。 |
| GET | /douyin/favorite/list/ | 登录用户的所有点赞视频 |
| POST | /douyin/comment/action/ | 评论操作 |
| GET | /douyin/comment/list/ | 查看视频的所有评论，按发布时间排序 |

