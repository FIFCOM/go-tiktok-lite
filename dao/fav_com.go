package dao

type UserFavorite struct {
	UserId  int
	VideoId int
}

func InsertUserFavorite(favorite UserFavorite) error {
	err := DB.Create(favorite).Error
	Handle(err)
	return err
}

//func DeleteUserFavorite(favorite UserFavorite) error {
//	err :=
//
//	return err
//}
