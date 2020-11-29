package psql

import (
	"github.com/nizonglonggit/userapisrv/pkg/define"
)

func RegisterUser(user define.User) (err error) {
	db := ExampleDB.Begin()
	err = ExampleDB.Table(define.TableUserUser).Create(&user).Error
	if err != nil {
		db.Rollback()
		return
	}
	db.Commit()
	return
}

func ListAllUsers() (users []define.User, err error) {
	users = []define.User{}
	err = ExampleDB.Table(define.TableUserUser).Find(&users).Error
	return
}
