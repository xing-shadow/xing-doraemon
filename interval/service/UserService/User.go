/*
 * @Time : 2020/11/13 17:32
 * @Author : wangyl
 * @File : UserService.go
 * @Software: GoLand
 */
package UserService

import (
	"errors"
	"github.com/jinzhu/gorm"
	"xing-doraemon/interval/model/db"
	"xing-doraemon/interval/model/view"
	"xing-doraemon/pkg/Auth/JwtAuth"
	"xing-doraemon/pkg/Utils"
)

func Login(req view.LoginReq) (resp view.LoginResp, err error) {
	var user db.User
	err = opt.DB.Where("name=?", req.Username).First(&user).Error
	if err != nil {
		return view.LoginResp{}, err
	}
	md5Password := Utils.Md5ToHex([]byte(req.Password))
	if md5Password != user.Password {
		return view.LoginResp{}, errors.New("username or password error")
	}
	var token string
	if token, err = JwtAuth.GenerateToken(user.ID, user.Name); err != nil {
		return view.LoginResp{}, errors.New("generate token fail: " + err.Error())
	} else {
		resp.Token = token
		return
	}
}

func UserList(req view.UserListReq) (resp view.UserListResp, err error) {
	var page, pageSize, offset uint
	var count int
	var users []db.User
	if req.Page <= 0 {
		page = 1
	} else {
		page = req.Page
	}
	if req.PageSize <= 0 {
		req.PageSize = 1000
	} else {
		pageSize = req.PageSize
	}
	offset = (page - 1) * pageSize
	err = opt.DB.Offset(offset).Limit(pageSize).Find(&users).Error
	if err != nil {
		return view.UserListResp{}, err
	}
	err = opt.DB.Table(db.User{}.TableName()).Count(&count).Error
	if err != nil {
		return view.UserListResp{}, err
	}
	for _, user := range users {
		resp.UserList = append(resp.UserList, view.UserItem{
			Id:       user.ID,
			UserName: user.Name,
		})
	}
	resp.Total = count
	resp.CurrentPage = int(page)
	return
}

func CreateUser(req view.UserCreateReq) (err error) {
	var user db.User
	err = opt.DB.Where("name=?", req.UserName).First(&user).Error
	if err != nil && !gorm.IsRecordNotFoundError(err) {
		return err
	}
	if user.ID > 0 {
		return errors.New("The user already exists ")
	}
	md5Password := Utils.Md5ToHex([]byte(req.Password))
	err = opt.Create(&db.User{
		Name:     req.UserName,
		Password: md5Password,
	}).Error
	return
}

func UpdateUser(req view.UserUpdateReq) (err error) {
	var user db.User
	err = opt.DB.Where("id=?", req.Id).First(&user).Error
	if err != nil {
		return err
	}
	err = opt.DB.Model(&db.User{}).Where("id=?", req.Id).Updates(&db.User{
		Name:     req.UserName,
		Password: req.Password,
	}).Error
	return
}

func DeleteUser(req view.UserDeleteReq) (err error) {
	var user db.User
	err = opt.DB.Where("id=?", req.Id).First(&user).Error
	if err != nil {
		return err
	}
	err = opt.DB.Where("id=?", req.Id).Delete(&user).Error
	return
}
