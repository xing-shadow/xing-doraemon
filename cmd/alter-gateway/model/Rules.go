/*
@Time : 2020/7/20 11:52
@Author : wangyl
@File : Rules.go
@Software: GoLand
*/
package model

import (
	"gopkg.in/mgo.v2/bson"

	"xing-doraemon/cmd/alter-gateway/Datamanager/MongoDataManger"
)

type Rules struct {
	Id          int64  `orm:"column(id);auto" json:"id,omitempty" bson:"Id"`
	Expr        string `orm:"column(expr);size(1023)" json:"expr" bson:"Expr"`
	Op          string `orm:"column(op);size(31)" json:"op" bson:"Op"`
	Value       string `orm:"column(value);size(1023)" json:"op" bson:"Value"`
	For         string `orm:"column(for);size(1023)" json:"for" bson:"For"`
	Summary     string `orm:"column(summary);size(1023)" json:"summary"`
	Description string `orm:"column(description);size(1023)" json:"description" bson:"Description"`
	Prom        *Proms `orm:"rel(fk)" json:"prom_id" bson:"Prom"`
	Plan        *Plans `orm:"rel(fk)" json:"plan_id" bson:"Plan"`
	//Labels      []*Labels `orm:"rel(m2m);rel_through(alert-gateway/models.RuleLabels)" json:"omitempty"`
}

func (r *Rules) CollectionName() string {
	return "Rule"
}

func (r *Rules) Sort() string {
	return "+Id"
}

func (r *Rules) Get(prom string, id string) []Rules {
	rules := []Rules{}
	baseFilter := []interface{}{bson.M{"Prom":prom}}
	baseFilter = append(baseFilter,bson.M{"Id":id})
	filter := bson.M{"$and":baseFilter}
	MongoDataManger.MongoClient.FindAll(r.CollectionName(),filter,r.Sort(),0,0,&rules)
	return rules
}