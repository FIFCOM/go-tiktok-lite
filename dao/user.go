package dao

// User 是用户结构体，对应users表
type User struct {
	Id       int64  // 自增主键，用户ID
	Name     string // 用户名
	Password string // 密码
}

// GetUserById 由ID获取用户结构体
func GetUserById(id int64) (User, error) {
	user := User{}
	err := DB.Where("id = ?", id).First(&user).Error
	Handle(err)
	return user, err
}

// GetUserByName 由用户名获取用户结构体
func GetUserByName(name string) (User, error) {
	user := User{}
	err := DB.Where("name = ?", name).First(&user).Error
	return user, err
}

// InsertUser 插入数据
func InsertUser(user *User) error {
	err := DB.Create(user).Error
	Handle(err)
	return err
}
