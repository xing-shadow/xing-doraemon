/*
@Time : 2020/8/25 11:23
@Author : wangyl
@File : algorithms.go
@Software: GoLand
*/
package common

import (
	"strings"
)

func CalculateReversePolishNotation(labelmap map[string]string, expression string) bool {
	stack := []bool{}
	exp := strings.Split(expression, ` `)
	for i := 0; i < len(exp); i++ {
		if exp[i] == "&" {
			var v1, v2 bool
			v2, stack = stack[len(stack)-1], stack[:len(stack)-1]
			v1, stack = stack[len(stack)-1], stack[:len(stack)-1]
			stack = append(stack, v1 && v2)
		} else if exp[i] == "|" {
			var v1, v2 bool
			v2, stack = stack[len(stack)-1], stack[:len(stack)-1]
			v1, stack = stack[len(stack)-1], stack[:len(stack)-1]
			stack = append(stack, v1 || v2)
		} else {
			if strings.Contains(exp[i], "!") {
				val := strings.Split(exp[i], "!=")
				if _, ok := labelmap[val[0]]; ok {
					if labelmap[val[0]] == val[1] {
						stack = append(stack, false)
					} else {
						stack = append(stack, true)
					}
				} else {
					return false
				}
			} else {
				val := strings.Split(exp[i], "=")
				if _, ok := labelmap[val[0]]; ok {
					if labelmap[val[0]] == val[1] {
						stack = append(stack, true)
					} else {
						stack = append(stack, false)
					}
				} else {
					return false
				}
			}
		}
	}
	return stack[len(stack)-1]
}
