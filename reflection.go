package tablex

import (
	"reflect"
)

type tablexTags map[string]string
type structFieldName string
type tableFieldName string

const (
	tablexTag         = "tablex"
	tablexHeaderField = "header"
)

func getReflectType(obj interface{}) reflect.Type {
	reflectType := reflect.TypeOf(obj)
	if reflect.ValueOf(obj).Kind() == reflect.Ptr {
		reflectType = reflect.Indirect(reflect.ValueOf(obj)).Type()
	}

	return reflectType
}

func isStruct(field reflect.StructField) bool {
	fieldOriginalKind := field.Type.Kind()
	if fieldOriginalKind == reflect.Ptr {
		fieldOriginalKind = field.Type.Elem().Kind()
	}

	return fieldOriginalKind == reflect.Struct
}
