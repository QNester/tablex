package tablex

import (
	"reflect"
	"strings"
)

func getTablexTags(field reflect.StructField) (reflect.Type, tablexTags) {
	fieldType := field.Type
	if fieldType.Kind() == reflect.Ptr {
		fieldType = fieldType.Elem()
	}

	tagVal := field.Tag.Get(tablexTag)
	if tagVal == "" || tagVal == "-" {
		return nil, nil
	}

	parsedTags := parseTablexTags(tagVal)
	headerVal, ok := parsedTags[tablexHeaderField]
	if !ok || headerVal == "" {
		return nil, nil
	}

	return fieldType, parsedTags
}

func parseTablexTags(rawTagVal string) tablexTags {
	if rawTagVal == "" {
		return tablexTags{}
	}

	pairs := strings.Split(rawTagVal, ",")
	tags := make(tablexTags)
	for _, pair := range pairs {
		splittedPair := strings.Split(pair, ":")
		if len(splittedPair) != 2 {
			continue
		}

		tags[splittedPair[0]] = splittedPair[1]
	}

	return tags
}
