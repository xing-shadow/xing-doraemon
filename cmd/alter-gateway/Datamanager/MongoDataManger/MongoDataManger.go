/*
@Time : 2020/7/20 10:11
@Author : wangyl
@File : MongoDataManger.go
@Software: GoLand
*/
package MongoDataManger

import (
	"net/url"

	"iPublic/MongoModular"
)

var  MongoUrl string

var MongoClient MongoModular.MongoDBServ

func InitDataMongoClient() error {
	urlParse,err := url.Parse(MongoUrl)
	if err != nil {
		return err
	}
	if err := MongoModular.GetMongoDBHandlerWithURL(MongoUrl,&MongoClient); err != nil {
		return err
	}else {
		MongoClient.URLS = []string{urlParse.Host}
	}
	return nil
}