package tools

import (
    "fmt"
    "reflect"
)

// ToMap 结构体转为Map[string]interface{}
func ToMap(in interface{}, tagName string) (map[string]interface{}, error){
    out := make(map[string]interface{})
    v := reflect.ValueOf(in)
    if v.Kind() == reflect.Ptr {
        // 指针转换
        v = v.Elem()
    }

    if v.Kind() != reflect.Struct {  // 非结构体返回错误提示
        return nil, fmt.Errorf("only accepts struct or struct pointer; got %T", v)
    }

    t := v.Type()
    // 遍历结构体字段
    // 指定tagName值为map中key;字段值为map中value
    for i := 0; i < v.NumField(); i++ {
        if tagValue := t.Field(i).Tag.Get(tagName); tagValue != "" {
            out[tagValue] = v.Field(i).Interface()
        }
    }
    return out, nil
}

// 将结构体转为单层map
func ToDouMap(in interface{}, tag string) (map[string]interface{}, error) {

    // 当前函数只接收struct类型
    v := reflect.ValueOf(in)
    if v.Kind() == reflect.Ptr { // 结构体指针
        v = v.Elem()
    }
    if v.Kind() != reflect.Struct {
        return nil, fmt.Errorf("ToMap only accepts struct or struct pointer; got %T", v)
    }

    out := make(map[string]interface{})
    queue := make([]interface{}, 0, 1)
    // 初始化待反射处理的数据
    queue = append(queue, in)

    for len(queue) > 0 {
        v := reflect.ValueOf(queue[0])
        if v.Kind() == reflect.Ptr { // 结构体指针
            v = v.Elem()
        }
        queue = queue[1:]
        t := v.Type()
        for i := 0; i < v.NumField(); i++ {
            // 判断值类型 是否嵌套struct
            vi := v.Field(i)
            if vi.Kind() == reflect.Ptr { // 内嵌指针
                vi = vi.Elem()
                if vi.Kind() == reflect.Struct { // 结构体
                    queue = append(queue, vi.Interface())
                } else {
                    ti := t.Field(i)
                    if tagValue := ti.Tag.Get(tag); tagValue != "" {
                        // 存入map
                        out[tagValue] = vi.Interface()
                    }
                }
                break
            }
            if vi.Kind() == reflect.Struct { // 内嵌结构体
                queue = append(queue, vi.Interface())
                break
            }
            // 一般字段
            ti := t.Field(i)
            if tagValue := ti.Tag.Get(tag); tagValue != "" {
                // 存入map
                out[tagValue] = vi.Interface()
            }
        }
    }
    return out, nil
}
