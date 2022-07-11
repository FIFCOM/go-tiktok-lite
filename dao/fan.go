package dao

// Follow 对应 Follows 表
type Follow struct {
	Id    int64
	Name  string
	focus []User // 关注列表
}

// Follower 对应 Followers 表
type Follower struct {
	Id   int64
	Name string
	fans []User // 粉丝列表
}

// GetUserfocus 通过用户id返回用户的关注列表
func GetUserfocus(id int64) ([]User, error) {
	user := Follow{}
	err := DB.Where("id = ?", id).First(&user).Error
	Handle(err)
	return user.focus, err
}

// GetUserfans 通过用户id返回用户的粉丝列表
func GetUserfans(id int64) ([]User, error) {
	user := Follower{}
	err := DB.Where("id = ?", id).First(&user).Error
	Handle(err)
	return user.fans, err
}
