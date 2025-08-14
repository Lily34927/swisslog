package utils

import (
	"fmt"
	"reflect"
	"log"
)


// 打印结构体的 key 和 value
func StructToMap(obj interface{}) (map[string]interface{}, error) {
	// 1. 检查输入是否为结构体
	t := reflect.TypeOf(obj)
	// 如果传入的是指针，获取其指向的类型
	if t.Kind() == reflect.Ptr {
		t = t.Elem()
	}
	if t.Kind() != reflect.Struct {
		return nil, fmt.Errorf("obj is not a struct or a pointer to a struct, but %s", t.Kind())
	}

	// 2. 获取结构体的值
	v := reflect.ValueOf(obj)
	// 如果传入的是指针，获取其指向的值
	if v.Kind() == reflect.Ptr {
		v = v.Elem()
	}
	
	data := make(map[string]interface{})

	// 3. 遍历结构体的字段
	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		value := v.Field(i).Interface()

		// 4. 获取 json tag 作为 map 的 key
		name := field.Tag.Get("json")
		if name == "" {
			name = field.Name // 如果没有 json tag，使用字段名
		}

		data[name] = value
	}

	for key, value := range data{
		log.Println("key:", key, ", value:", value)
	}
	
	return data, nil
}