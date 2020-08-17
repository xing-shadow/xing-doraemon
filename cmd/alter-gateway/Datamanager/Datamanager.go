/*
@Time : 2020/7/20 10:13
@Author : wangyl
@File : Datamanager.go
@Software: GoLand
*/
package Datamanager

import (
	"xing-doraemon/cmd/alter-gateway/Datamanager/MongoDataManger"
)

func Init() error {
	return MongoDataManger.InitDataMongoClient()
}
