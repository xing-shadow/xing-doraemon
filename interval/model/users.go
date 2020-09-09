/*
@Time : 2020/8/25 10:39
@Author : wangyl
@File : users.go
@Software: GoLand
*/
package model

import (
	"crypto/md5"
	"fmt"

	"github.com/jinzhu/gorm"
	"github.com/pkg/errors"

	"xing-doraemon/pkg/common"
)

type Users struct {
	Id       int64  `gorm:"AUTO_INCREMENT" json:"id,omitempty"`
	Name     string `gorm:"unique;size:255" json:"name"`
	Password string `gorm:"size:1023" json:"password,omitempty"`
}

func (u Users) TableName() string {
	return "users"
}

func (u *Users) CheckUser(db *gorm.DB, userinfo common.AuthModel) (*common.AuthModel, error) {
	var results Users
	if err := db.Table(u.TableName()).Select("id,password").
		Where("name=?", userinfo.Username).First(&results).Error; err != nil {
		return nil, err
	}
	hash := fmt.Sprintf("%x", md5.Sum([]byte(userinfo.Password)))
	if hash != results.Password {
		return nil, errors.Errorf("invalid password")
	} else {
		return &userinfo, nil
	}
}

func (u *Users) GetAll(db *gorm.DB) []Users {
	var result []Users
	db.Table(u.TableName()).Select("id,name").Order("id ASC").Find(&result)
	return result
}

func (u *Users) AddUser(db *gorm.DB) error {
	u.Password = fmt.Sprintf("%x", md5.Sum([]byte("123456")))
	return db.Table(u.TableName()).Create(&u).Error
}

func (u *Users) UpdatePassword(db *gorm.DB, name string, oldPassword string, newPassword string) error {
	var user []Users
	if err := db.Select(u.TableName()).Select("id,name,password").
		Where("name=?", name).Find(&user).Error; err != nil {
		return err
	} else {
		if len(user) > 0 {
			if user[0].Password == fmt.Sprintf("%x", md5.Sum([]byte(oldPassword))) {
				updateErr := db.Model(&Users{}).Update(map[string]interface{}{
					"password": fmt.Sprintf("%x", md5.Sum([]byte(newPassword))),
				}).Where("name=?", name).Error
				return errors.Wrap(updateErr, "database update error")
			} else {
				return errors.Errorf("wrong password")
			}
		} else {
			return errors.Errorf("user is not exist")
		}
	}
}

func (u *Users) DeleteUsers(db *gorm.DB, id string) error {
	if err := db.Where("id=?", id).Delete(&Users{}).Error; err != nil {
		return errors.Wrap(err, "database delete error")
	} else {
		return err
	}
}
