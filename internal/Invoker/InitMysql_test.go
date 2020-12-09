/*
@Time : 2020/9/3 18:01
@Author : wangyl
@File : InitMysql_test.go.go
@Software: GoLand
*/
package Invoker

import (
	"database/sql/driver"
	"encoding/json"
	"fmt"
	"testing"

	"github.com/jinzhu/gorm"

	"xing-doraemon/global"
)

func TestQuery(t *testing.T) {
	global.GetAlterGatewayConfig().Mysql.DBType = "mysql"
	global.GetAlterGatewayConfig().Mysql.DBUser = "root"
	global.GetAlterGatewayConfig().Mysql.DBPasswd = "1"
	global.GetAlterGatewayConfig().Mysql.DBTns = "192.168.0.71:3306"
	global.GetAlterGatewayConfig().Mysql.DBName = "doraemon" //doraemon
	global.GetAlterGatewayConfig().Mysql.DBLoc = "Asia%2FShanghai"
	if err := Init(); err != nil {
		t.Fatal(err)
	}
	//var results model.Users
	//doraemonMysql.Table(results.TableName()).Select("id,name").Order("id ASC").Find(&results)
	//now := time.Now().Format("2006-01-02 15:04:05")
	//doraemonMysql.Model(&model.Alerts{}).Where("status=? AND confirmed_before<?",1,now).Update("status",2)
	//var results []*model.Users
	//doraemonMysql.Select("id,name").Where("name = ?","admin").Find(&results)
	//doraemonMysql.Model(&model.Hosts{}).Update(map[string]string{
	//	"hostname":"localhost:9091",
	//}).Where("hostname = ?","localhost:9090")
	var method GrpcServiceMethod
	var err error
	err = doraemonMysql.Where("id = ?", "10").
		Preload("Service", func(db *gorm.DB) *gorm.DB {
			return db.Preload("Proto")
		}).
		First(&method).Error
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(method)
}

type (
	GrpcServiceMethod struct {
		gorm.Model

		ServiceID     uint        `gorm:"column:service_id;"`
		Name          string      `gorm:"column:name;not null;"`
		MethodComment string      `gorm:"column:method_comment;"`
		InputName     string      `gorm:"column:input_name;not null;"`
		InputType     ProtoFields `gorm:"column:input_type;not null;type:json"`  // 入参类型描述
		OutputType    ProtoFields `gorm:"column:output_type;not null;type:json"` // 返回值类型描述
		OutputName    string      `gorm:"column:output_name;not null;"`

		Service   GrpcProtoService `gorm:"foreignKey:ServiceID"`
		TestCases []GrpcTestCase   `gorm:"foreignKey:MethodID"`
	}

	GrpcProtoService struct {
		gorm.Model

		ProtoID uint
		Name    string

		Proto GrpcProto `gorm:"foreignKey:ProtoID"`
	}
	GrpcProto struct {
		gorm.Model

		AppName     string
		FileName    string
		PackageName string

		Services []GrpcProtoService `gorm:"foreignKey:ProtoID"`
	}
	// GRPC 测试用例
	GrpcTestCase struct {
		gorm.Model
		MethodID uint
		Uid      uint
		Name     string
		Input    string        `gorm:"type:longtext;"`
		Metadata ProtoMetadata `gorm:"type:longtext;"`
		Script   string        `gorm:"type:longtext;"`

		Method GrpcServiceMethod `gorm:"foreignKey:MethodID"`
	}
	ProtoMetadata []struct {
		Key         string `json:"key"`
		Value       string `json:"value"`
		Description string `json:"description"`
	}
	// map: 字段名 => 类型描述
	ProtoFields map[string]ProtoField
	ProtoField  struct {
		JsonName    string      `json:"json_name"`
		Type        int32       `json:"type"`
		Label       int32       `json:"label"`
		Number      int32       `json:"number"`
		IsRepeated  bool        `json:"is_repeated"`
		MessageType ProtoFields `json:"message_type"`
		Comment     string      `json:"comment"`
	}
)

func (c ProtoFields) Value() (driver.Value, error) {
	b, err := json.Marshal(c)
	return string(b), err
}

func (c *ProtoFields) Scan(input interface{}) error {
	return json.Unmarshal(input.([]byte), c)
}

func (m ProtoMetadata) Value() (driver.Value, error) {
	b, err := json.Marshal(m)
	return string(b), err
}

func (m *ProtoMetadata) Scan(input interface{}) error {
	return json.Unmarshal(input.([]byte), m)
}

func (GrpcProto) TableName() string {
	return "grpc_proto"
}

func (GrpcProtoService) TableName() string {
	return "grpc_proto_service"
}

func (GrpcServiceMethod) TableName() string {
	return "grpc_service_method"
}

func (GrpcTestCase) TableName() string {
	return "grpc_test_case"
}
