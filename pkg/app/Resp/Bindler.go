/*
 * @Time : 2020/10/20 16:21
 * @Author : wangyl
 * @File : Bindler.go
 * @Software: GoLand
 */
package Resp

import (
	"errors"
	"reflect"
	"strings"
)

type (
	// Binder is the interface that wraps the Bind method.
	Binder interface {
		Bind(i interface{}, c Context) error
	}

	// DefaultBinder is the default implementation of the Binder interface.
	DefaultBinder struct{}

	// BindUnmarshaler is the interface used to wrap the UnmarshalParam method.
	// Types that don't implement this, but do implement encoding.TextUnmarshaler
	// will use that interface instead.
	BindUnmarshaler interface {
		// UnmarshalParam decodes and assigns a value from an form or query param.
		UnmarshalParam(param string) error
	}
)

// Bind implements the `Binder#Bind` function.
func (b *DefaultBinder) Bind(i interface{}, c Context) (err error) {
	req := c.Request
	params := map[string][]string{}
	for _, param := range c.Params {
		params[param.Key] = []string{param.Value}
	}

}

func (b *DefaultBinder) bindData(ptr interface{}, data map[string][]string, tag string) error {
	if ptr == nil || len(data) == 0 {
		return nil
	}

	typ := reflect.TypeOf(ptr).Elem()
	val := reflect.ValueOf(ptr).Elem()

	//Map
	if typ.Kind() == reflect.Map {
		for k, v := range data {
			val.SetMapIndex(reflect.ValueOf(k), reflect.ValueOf(v))
		}
		return nil
	}

	//!struct
	if typ.Kind() != reflect.Struct {
		return errors.New("binding element must be a struct")
	}
	for i := 0; i < typ.NumField(); i++ {
		typeField := typ.Field(i)
		structField := val.Field(i)
		if !structField.CanSet() {
			continue
		}
		structFieldKind := structField.Kind()
		inputFieldName := typeField.Tag.Get(tag)
		if inputFieldName == "" {
			inputFieldName = typeField.Name
			// If tag is nil, we inspect if the field is a struct.
			if _, ok := structField.Addr().Interface().(BindUnmarshaler); !ok && structFieldKind == reflect.Struct {
				if err := b.bindData(structField.Addr().Interface(), data, tag); err != nil {
					return err
				}
				continue
			}
		}
		inputValue, exists := data[inputFieldName]
		if !exists {
			// Go json.Unmarshal supports case insensitive binding.  However the
			// url params are bound case sensitive which is inconsistent.  To
			// fix this we must check all of the map values in a
			// case-insensitive search.
			for k, v := range data {
				if strings.EqualFold(k, inputFieldName) {
					inputValue = v
					exists = true
					break
				}
			}
		}
		if !exists {
			continue
		}
		// Call this first, in case we're dealing with an alias to an array type
	}
}
