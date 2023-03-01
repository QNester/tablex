package tablex

import (
	"reflect"
	"testing"
)

type getTablexTagsTestStruct struct {
	fieldWithHeader                  string  `tablex:"header:good_header"`
	fieldWithHeaderPtr               *string `tablex:"header:ptr_header"`
	emptyField                       string  `tablex:""`
	emptyHeader                      string  `tablex:"header:"`
	skipMark                         string  `tablex:"-"`
	fieldWithAnotherTagWithoutHeader string  `tablex:"another_tag:good_header"`
	fieldWithAnotherTagWithHeader    string  `tablex:"header:good,another_tag:another"`
}

func Test_getTablexTags(t *testing.T) {
	obj := getTablexTagsTestStruct{}
	reflType := reflect.TypeOf(obj)
	type args struct {
		field reflect.StructField
	}
	tests := []struct {
		name  string
		args  args
		want  reflect.Type
		want1 tablexTags
	}{
		{
			name:  "field with header",
			args:  args{field: reflType.Field(0)},
			want:  reflType.Field(0).Type,
			want1: map[string]string{"header": "good_header"},
		},
		{
			name:  "field with header as ptr",
			args:  args{field: reflType.Field(1)},
			want:  reflType.Field(1).Type.Elem(),
			want1: map[string]string{"header": "ptr_header"},
		},
		{
			name:  "with empty tablex val",
			args:  args{field: reflType.Field(2)},
			want:  nil,
			want1: nil,
		},
		{
			name:  "with empty header",
			args:  args{field: reflType.Field(3)},
			want:  nil,
			want1: nil,
		},
		{
			name:  "with skip tablex mark",
			args:  args{field: reflType.Field(4)},
			want:  nil,
			want1: nil,
		},
		{
			name:  "field with another tag without header",
			args:  args{field: reflType.Field(5)},
			want:  nil,
			want1: nil,
		},
		{
			name:  "field with another tag with header",
			args:  args{field: reflType.Field(6)},
			want:  reflType.Field(6).Type,
			want1: map[string]string{"header": "good", "another_tag": "another"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := getTablexTags(tt.args.field)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("getTablexTags() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("getTablexTags() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}
