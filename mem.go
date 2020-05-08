package mem

import (
	"reflect"
)

func GetMemoryConsumption(variable interface{}, isRecursed ...bool) int {
	var getValue bool
	if len(isRecursed) > 0 {
		getValue = isRecursed[0]
	}
	totalSize := 0
	var value reflect.Value
	if getValue {
		value = variable.(reflect.Value)
	} else {
		if reflect.ValueOf(variable).Kind() == reflect.Ptr {
			value = reflect.ValueOf(variable).Elem()
		} else {
			value = reflect.ValueOf(variable)
		}
	}
	if value.Kind() == reflect.Map {
		iter := value.MapRange()
		for iter.Next() {
			totalSize += GetMemoryConsumption(iter.Key(), true)
			totalSize += GetMemoryConsumption(iter.Value(), true)
		}
	} else if value.Kind() == reflect.Slice {
		for i := 0; i < value.Len(); i++ {
			totalSize += GetMemoryConsumption(value.Index(i), true)
		}
	} else if value.Kind() == reflect.Struct {
		for i := 1; i < value.NumField(); i++ {
			totalSize += GetMemoryConsumption(value.Field(i), true)
		}
	} else {
		totalSize += int(reflect.TypeOf(value).Size())
	}
	return totalSize
}
