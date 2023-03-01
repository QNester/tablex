package tablex

import (
	"errors"
	"reflect"

	"github.com/jedib0t/go-pretty/v6/table"
)

var ObjectIsNotStruct = errors.New("tablex_error: received object is not a structure")

type tablexInfo struct {
	isCollection bool
	reflType     reflect.Type

	emptyVal string

	headers []interface{}
	fields  map[structFieldName]*fieldInfo
}

func newTablexInfo(obj interface{}, emptyVal string) (*tablexInfo, error) {
	var isCollection bool

	reflType := getReflectType(obj)

	if reflType.Kind() == reflect.Slice || reflType.Kind() == reflect.Array {
		isCollection = true
		reflType = reflType.Elem()
	}

	if reflType.Kind() != reflect.Struct {
		return nil, ObjectIsNotStruct
	}

	tInfo := &tablexInfo{
		isCollection: isCollection,
		reflType:     reflType,
		emptyVal:     emptyVal,
	}

	tInfo.loadFields()
	tInfo.loadHeaders()

	return tInfo, nil
}

func (tx *tablexInfo) loadFields() {
	tx.fields = make(map[structFieldName]*fieldInfo)
	lastIdx := 0

	for i := 0; i < tx.reflType.NumField(); i++ {
		field := tx.reflType.Field(i)
		lastIdx = tx.loadFieldInfo(field, lastIdx, nil)
	}
}

func (tx *tablexInfo) loadHeaders() {
	headers := make([]interface{}, len(tx.fields))
	for _, fi := range tx.fields {
		headers[fi.idxInTable] = fi.columnHeader
	}

	tx.headers = headers
}

func (tx *tablexInfo) rowsForObject(obj interface{}) []table.Row {
	if !tx.isCollection {
		row := tx.rowForObject(obj)
		return []table.Row{row}
	}

	reflValue := reflect.ValueOf(obj)
	rows := make([]table.Row, reflValue.Len())

	for i := 0; i < reflValue.Len(); i++ {
		collectionElem := reflValue.Index(i).Interface()
		rows[i] = tx.rowForObject(collectionElem)
	}

	return rows
}

func (tx *tablexInfo) rowForObject(obj interface{}) table.Row {
	reflValue := reflect.ValueOf(obj)
	values := make([]interface{}, len(tx.headers))

	for _, fi := range tx.fields {
		values[fi.idxInTable] = fi.getFieldValue(reflValue, tx.emptyVal)
	}

	return values
}
