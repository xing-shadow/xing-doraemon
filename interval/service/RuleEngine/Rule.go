/*
 * @Time : 2020/10/23 14:49
 * @Author : wangyl
 * @File : Rule.go
 * @Software: GoLand
 */
package RuleEngine

import (
	"gopkg.in/yaml.v2"
	"strconv"
	"strings"
)

// M is map
type M map[string]interface{}

// S is slice
type S []interface{}

// Prom ...
type Prom struct {
	ID  int64
	URL string
}

type Rules []Rule

type Rule struct {
	ID          int64             `json:"id"`
	PromID      int64             `json:"prom_id"`
	Expr        string            `json:"expr"`
	Op          string            `json:"op"`
	Value       string            `json:"value"`
	For         string            `json:"for"`
	Labels      map[string]string `json:"labels"`
	Summary     string            `json:"summary"`
	Description string            `json:"description"`
}

type PromRules struct {
	Prom  Prom
	Rules Rules
}

func (r Rules) Content() ([]byte, error) {
	rules := S{}
	for _, rule := range r {
		rules = append(rules, M{
			"alert":  strconv.FormatInt(rule.ID, 10),
			"expr":   strings.Join([]string{rule.Expr, rule.Op, rule.Value}, " "),
			"for":    rule.For,
			"labels": rule.Labels,
			"annotations": M{
				"rule_id":     strconv.FormatInt(rule.ID, 10),
				"prom_id":     strconv.FormatInt(rule.PromID, 10),
				"summary":     rule.Summary,
				"description": rule.Description,
			},
		})
	}
	result := M{
		"groups": S{
			M{
				"name":  "ruleengine",
				"rules": rules,
			},
		},
	}
	return yaml.Marshal(result)
}
