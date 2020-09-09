/*
@Time : 2020/8/25 10:38
@Author : wangyl
@File : group.go
@Software: GoLand
*/
package model

import (
	"encoding/json"
	"github.com/pkg/errors"
	"io/ioutil"
	"strings"
	"time"

	"github.com/jinzhu/gorm"

	"xing-doraemon/gobal"
	"xing-doraemon/pkg/common"
)

type Groups struct {
	Id   int64  `gorm:"AUTO_INCREMENT" json:"id,omitempty"`
	Name string `gorm:"unique;size:255" json:"name"`
	User string `gorm:"size:1023" json:"user"`
}

type HttpRes struct {
	Status int    `json:"status"`
	Msg    string `json:"msg"`
	Data   []struct {
		ID      string `json:"id"`
		Title   string `json:"title"`
		Mobile  string `json:"mobile"`
		Email   string `json:"email"`
		AddTime string `json:"add_time"`
		Account string `json:"account"`
	} `json:"data"`
}

func (Groups) TableName() string {
	return "group"
}

func SendAlertsFor(db *gorm.DB, VUG *common.ValidUserGroup) []string {
	var userList []string
	if VUG.User != "" {
		userList = strings.Split(VUG.User, ",")
	}
	if VUG.Group != "" {
		var groups []*Groups
		db.Table(Groups{}.TableName()).Where("name in (?)", strings.Split(VUG.Group, ",")).Find(&groups)
		for _, v := range groups {
			userList = append(userList, strings.Split(v.User, ",")...)
		}
	}
	if VUG.DutyGroup != "" {
		date := time.Now().Format("2006-1-2")
		idList := strings.Split(VUG.DutyGroup, ",")
		for _, id := range idList {
			res, _ := common.HttpGet(gobal.GetAlterGatewayConfig().Send.DutyGroupUrl, map[string]string{"teamId": id, "day": date}, nil)
			info := HttpRes{}
			jsonDataFromHttp, _ := ioutil.ReadAll(res.Body)
			json.Unmarshal(jsonDataFromHttp, &info)
			for _, i := range info.Data {
				userList = append(userList, i.Account)
			}
		}
	}
	hashMap := map[string]bool{}
	for _, name := range userList {
		hashMap[name] = true
	}
	res := []string{}
	for key := range hashMap {
		res = append(res, key)
	}
	return res
}

func (g *Groups) GetAll(db *gorm.DB) []Groups {
	groups := []Groups{}
	db.Table(g.TableName()).Find(&groups)
	return groups
}

func (g *Groups) AddGroup(db *gorm.DB) error {
	err := db.Create(&g).Error
	return errors.Wrap(err, "database insert error")
}

func (g *Groups) UpdateGroup(db *gorm.DB) error {
	err := db.Model(&Groups{}).Where("id=?", g.Id).Updates(g).Error
	return errors.Wrap(err, "database update error")
}

func (g *Groups) DeleteGroup(db *gorm.DB, id string) error {
	err := db.Where("id=?", g.Id).Delete(&Groups{}).Error
	return errors.Wrap(err, "database delete error")
}
