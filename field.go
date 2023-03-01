package tablex

import (
	"fmt"
	"reflect"
	"strings"
)

const nestedHeadersSeparator = " / "

type fieldInfo struct {
	isStruct bool

	idxInTable      int
	columnHeader    tableFieldName
	structFieldName structFieldName
	tags            tablexTags
}

func (tx *tablexInfo) loadFieldInfo(field reflect.StructField, lastIndex int, parentFI *fieldInfo) int {
	fieldType, parsedTags := getTablexTags(field)
	if fieldType == nil {
		return lastIndex
	}

	resultFI := &fieldInfo{
		columnHeader:    tableFieldName(parsedTags[tablexHeaderField]),
		tags:            parsedTags,
		structFieldName: structFieldName(field.Name),
	}

	if parentFI != nil {
		resultFI.columnHeader = tableFieldName(
			fmt.Sprintf("%s%s%s", parentFI.columnHeader, nestedHeadersSeparator, resultFI.columnHeader),
		)
		resultFI.structFieldName = structFieldName(fmt.Sprintf("%s.%s", parentFI.structFieldName, field.Name))
	}

	if isStruct(field) {
		resultFI.isStruct = true

		for i := 0; i < fieldType.NumField(); i++ {
			nestedField := fieldType.Field(i)
			lastIndex = tx.loadFieldInfo(nestedField, lastIndex, resultFI)
		}

		return lastIndex
	}

	resultFI.idxInTable = lastIndex
	tx.fields[resultFI.structFieldName] = resultFI

	return resultFI.idxInTable + 1
}

func (fi *fieldInfo) getFieldValue(reflValue reflect.Value, emptyVal string) interface{} {
	fieldNames := strings.Split(string(fi.structFieldName), ".")
	for _, fName := range fieldNames {
		if reflValue.Kind() == reflect.Ptr {
			if reflValue.IsNil() {
				return emptyVal
			}

			reflValue = reflValue.Elem().FieldByName(fName)
			continue
		}

		reflValue = reflValue.FieldByName(fName)
	}

	if !reflValue.IsValid() {
		return emptyVal
	}

	return reflValue.Interface()
}
