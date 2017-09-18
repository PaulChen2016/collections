package collections

import (
	"errors"
	"reflect"
)

// MapFunc 模拟函数式的map方法
type MapFunc func(v interface{}) interface{}

// FilterFunc 模拟函数式发Filter方法
type FilterFunc func(v interface{}) bool

// ReduceFunc 模拟函数式的Reduce方法
type ReduceFunc func(v1, v2 interface{}) interface{}

// ForEachFunc 模拟forEach方法
type ForEachFunc func(v interface{})

// EveryFunc 模拟every方法
type EveryFunc func(v interface{}) bool

// SomeFunc 模拟some方法
type SomeFunc func(v interface{}) bool

// GetOneFunc 查询方法
type GetOneFunc func(v interface{}) bool

// JasonSlice 结构体，方法的承载对象
type JasonSlice struct {
	s []interface{}
}

// NewSlice 实例化函数承载对象
func NewSlice(s []interface{}) *JasonSlice {
	instance := new(JasonSlice)
	instance.s = s
	return instance
}

// NewInstance 通过 []* 类型初始化承载对象
func NewInstance(s interface{}) (instance *JasonSlice, err error) {
	instance = new(JasonSlice)
	_value := reflect.ValueOf(s)
	if _value.Kind() != reflect.Slice {
		return nil, errors.New("InterfaceSlice() given a non-slice type")
	}
	instance.s = make([]interface{}, _value.Len())
	for i := 0; i < _value.Len(); i++ {
		instance.s[i] = _value.Index(i).Interface()
	}
	return instance, nil
}

// GetData 方法 ，用于运算后获取结果
func (j *JasonSlice) GetData() []interface{} {
	return j.s
}

// Map 方法
func (j *JasonSlice) Map(mapf MapFunc) *JasonSlice {
	newJasonSlice := make([]interface{}, 0, len(j.s))
	for _, v := range j.s {
		newJasonSlice = append(newJasonSlice, mapf(v))
	}
	return &JasonSlice{s: newJasonSlice}
}

// Filter 方法
func (j *JasonSlice) Filter(filter FilterFunc) *JasonSlice {
	newJasonSlice := make([]interface{}, 0, len(j.s))
	for _, v := range j.s {
		if filter(v) {
			newJasonSlice = append(newJasonSlice, v)
		}
	}
	return &JasonSlice{s: newJasonSlice}
}

// 普通的 Reduce 方法
func (j *JasonSlice) Reduce(init interface{}, reducer ReduceFunc) interface{} {
	reduceValue := init
	for _, v := range j.s {
		reduceValue = reducer(reduceValue, v)
	}
	return reduceValue
}

// Reduce2Slice 运算的结果依然是slice，可以继续使用slice的函数方法
func (j *JasonSlice) Reduce2Slice(init interface{}, reducer ReduceFunc) *JasonSlice {
	reduceValue := init
	for _, v := range j.s {
		reduceValue = reducer(reduceValue, v)
	}
	_value := reflect.ValueOf(reduceValue)
	if _value.Kind() != reflect.Slice {
		panic("init given a non-slice type")
	}
	r := make([]interface{}, _value.Len())
	for i := 0; i < _value.Len(); i++ {
		r[i] = _value.Index(i).Interface()
	}
	return &JasonSlice{s: r}
}

// ForEach 方法
func (j *JasonSlice) ForEach(forEacher ForEachFunc) {
	for _, v := range j.s {
		forEacher(v)
	}
}

// Every 方法
func (j *JasonSlice) Every(everyer EveryFunc) bool {
	for _, v := range j.s {
		if !everyer(v) {
			return false
		}
	}
	return true
}

// Some 方法
func (j *JasonSlice) Some(somer SomeFunc) bool {
	for _, v := range j.s {
		if somer(v) {
			return true
		}
	}
	return false
}

// GetOne 方法
func (j *JasonSlice) GetOne(getOner GetOneFunc) interface{} {
	for _, v := range j.s {
		if getOner(v) {
			return v
		}
	}
	return nil
}
