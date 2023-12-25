// @Author zhangjiaozhu 2023/10/14 18:29:00
package common

import (
	"fmt"
	"reflect"
	"strconv"
	"time"
)

// DataToStructByTagSql 根据结构体中sqL标签映射数据到结构体中并且转换类型
func DataToStructByTagSql(data map[string]string,obj interface{})  {
	elem := reflect.ValueOf(obj).Elem()
	for i := 0 ; i < elem.NumField(); i++{
		// 获取sql对应值
		value := data[elem.Type().Field(i).Tag.Get("sql")]
		// 获取对应字段的名称
		name := elem.Type().Field(i).Name
		// 获取对应字段的类型
		structFieldType := elem.Field(i).Type()
		// 获取变量类型，也可直接写“string类型"
		val := reflect.ValueOf(value)
		var err error
		if structFieldType != val.Type(){
			// 类型转换
			val , err = TypeConversion(value,structFieldType.Name())
			if err != nil {

			}
		}
		// 设置类型值
		elem.FieldByName(name).Set(val)
	}
}
func TypeConversion(value string,ntype string)(reflect.Value,error)  {
	if ntype == "string"{
		return reflect.ValueOf(value),nil
	}else if ntype == "time.Time"{
		t,err := time.ParseInLocation("2006-01-02 15:04:05",value,time.Local)
		return reflect.ValueOf(t),err
	}else if ntype == "Time"{
		t,err := time.ParseInLocation("2006-01-02 15:04:05",value,time.Local)
		return reflect.ValueOf(t),err
	}else if ntype == "int" {
		i ,err := strconv.Atoi(value)
		return reflect.ValueOf(i),err
	}else if ntype == "int8" {
		i, err := strconv.ParseInt(value, 10, 8)
		if err != nil {
			return reflect.ValueOf(int8(0)), err
		}
		return reflect.ValueOf(int8(i)), nil
	} else if ntype == "int32" {
		i, err := strconv.ParseInt(value, 10, 32)
		if err != nil {
			return reflect.ValueOf(int32(0)), err
		}
		return reflect.ValueOf(int32(i)), nil
	} else if ntype == "int64" {
		i, err := strconv.ParseInt(value, 10, 64)
		if err != nil {
			return reflect.ValueOf(int64(0)), err
		}
		return reflect.ValueOf(i), nil
	}else if ntype == "float32" {
		f, err := strconv.ParseFloat(value, 32)
		if err != nil {
			return reflect.ValueOf(float32(0.0)), err
		}
		return reflect.ValueOf(float32(f)), nil
	} else if ntype == "float64" {
		f, err := strconv.ParseFloat(value, 64)
		if err != nil {
			return reflect.ValueOf(float64(0.0)), err
		}
		return reflect.ValueOf(f), nil
	} else if ntype == "bool" {
		b, err := strconv.ParseBool(value)
		if err != nil {
			return reflect.ValueOf(false), err
		}
		return reflect.ValueOf(b), nil
	}
	return reflect.Value{}, fmt.Errorf("不支持的类型转换: %s", ntype)
}