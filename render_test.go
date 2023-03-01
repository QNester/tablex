package tablex

import (
	"reflect"
	"testing"

	mocks "github.com/QNester/tablex/mock"
	"github.com/jedib0t/go-pretty/v6/table"
)

type renderNestedTestStruct struct {
	UserID int `tablex:"header:user_id"`
}

type rendererTestStruct struct {
	ID           int                     `tablex:"header:id"`
	User         *renderNestedTestStruct `tablex:"header:user"`
	Name         string                  `tablex:"header:name"`
	Email        *string                 `tablex:"header:email"`
	NoHeader     string                  `tablex:"header:"`
	NoTablex     string
	privateField int
}

func Test_tableRenderer_Render(t *testing.T) {
	type fields struct {
		writer  Writer
		options RendererOptions
	}
	email := "email"
	writer := mocks.NewWriterMock(t)
	headers := table.Row{tableFieldName("id"), tableFieldName("user / user_id"), tableFieldName("name"), tableFieldName("email")}
	type args struct {
		obj interface{}
	}
	tests := []struct {
		name   string
		fields fields
		expect func()
		args   args
		err    error
	}{
		{
			name: "default renderer received alone object",
			fields: fields{
				writer: writer,
				options: RendererOptions{
					EmptyValue: DefaultEmptyValue,
					Format:     DefaultRenderFormat,
				},
			},
			args: args{
				obj: rendererTestStruct{
					ID: 1,
					User: &renderNestedTestStruct{
						UserID: 1,
					},
					Name:         "test",
					NoHeader:     "no_header",
					NoTablex:     "no_tablex",
					privateField: 1,
				},
			},
			expect: func() {
				writer.On("AppendHeader", headers).Return()
				writer.On("AppendRows", []table.Row{{1, 1, "test", DefaultEmptyValue}}).Return()
				writer.On("Render").Return("")
			},
		},
		{
			name: "default renderer received alone object pointer",
			fields: fields{
				writer: writer,
				options: RendererOptions{
					EmptyValue: DefaultEmptyValue,
					Format:     DefaultRenderFormat,
				},
			},
			args: args{
				obj: &rendererTestStruct{
					ID: 1,
					User: &renderNestedTestStruct{
						UserID: 1,
					},
					Name:         "test",
					NoHeader:     "no_header",
					NoTablex:     "no_tablex",
					privateField: 1,
				},
			},
			expect: func() {
				writer.On("AppendHeader", headers).Return()
				writer.On("AppendRows", []table.Row{{1, 1, "test", DefaultEmptyValue}}).Return()
				writer.On("Render").Return("")
			},
		},
		{
			name: "default renderer alone object with null pointer",
			fields: fields{
				writer: writer,
				options: RendererOptions{
					EmptyValue: DefaultEmptyValue,
					Format:     DefaultRenderFormat,
				},
			},
			args: args{
				obj: rendererTestStruct{
					ID:   1,
					User: nil,
					Name: "name",
				},
			},
			expect: func() {
				writer.On("AppendHeader", headers).Return()
				writer.On("AppendRows", []table.Row{{1, DefaultEmptyValue, "name", DefaultEmptyValue}}).Return()
				writer.On("Render").Return("")
			},
		},
		{
			name: "special emptyValue renderer alone object with null pointer",
			fields: fields{
				writer: writer,
				options: RendererOptions{
					EmptyValue: "ooops, it's empty!",
					Format:     DefaultRenderFormat,
				},
			},
			args: args{
				obj: rendererTestStruct{
					ID:    1,
					User:  nil,
					Name:  "name",
					Email: &email,
				},
			},
			expect: func() {
				writer.On("AppendHeader", headers).Return()
				writer.On("AppendRows", []table.Row{{1, "ooops, it's empty!", "name", &email}}).Return()
				writer.On("Render").Return("")
			},
		},
		{
			name: "default renderer received objects collections",
			fields: fields{
				writer: writer,
				options: RendererOptions{
					EmptyValue: DefaultEmptyValue,
					Format:     DefaultRenderFormat,
				},
			},
			args: args{
				obj: []rendererTestStruct{
					{
						ID: 1,
						User: &renderNestedTestStruct{
							UserID: 1,
						},
						Name: "test",
					},
					{
						ID: 2,
						User: &renderNestedTestStruct{
							UserID: 2,
						},
						Name: "test_2",
					},
				},
			},
			expect: func() {
				writer.On("AppendHeader", headers).Return()
				rows := []table.Row{{1, 1, "test", DefaultEmptyValue}, {2, 2, "test_2", DefaultEmptyValue}}
				writer.On("AppendRows", rows).Return()
				writer.On("Render").Return("")
			},
		},
		{
			name: "CSV format renderer received alone object",
			fields: fields{
				writer: writer,
				options: RendererOptions{
					EmptyValue: DefaultEmptyValue,
					Format:     RenderFormatCSV,
				},
			},
			args: args{
				obj: rendererTestStruct{
					ID: 1,
					User: &renderNestedTestStruct{
						UserID: 1,
					},
					Name: "test",
				},
			},
			expect: func() {
				writer.On("AppendHeader", headers).Return()
				writer.On("AppendRows", []table.Row{{1, 1, "test", DefaultEmptyValue}}).Return()
				writer.On("RenderCSV").Return("")
			},
		},
		{
			name: "HTML format renderer received alone object",
			fields: fields{
				writer: writer,
				options: RendererOptions{
					EmptyValue: DefaultEmptyValue,
					Format:     RenderFormatHTML,
				},
			},
			args: args{
				obj: rendererTestStruct{
					ID: 1,
					User: &renderNestedTestStruct{
						UserID: 1,
					},
					Name: "test",
				},
			},
			expect: func() {
				writer.On("AppendHeader", headers).Return()
				writer.On("AppendRows", []table.Row{{1, 1, "test", DefaultEmptyValue}}).Return()
				writer.On("RenderHTML").Return("")
			},
		},
		{
			name: "MD format renderer received alone object",
			fields: fields{
				writer: writer,
				options: RendererOptions{
					EmptyValue: DefaultEmptyValue,
					Format:     RenderFormatMD,
				},
			},
			args: args{
				obj: rendererTestStruct{
					ID: 1,
					User: &renderNestedTestStruct{
						UserID: 1,
					},
					Name: "test",
				},
			},
			expect: func() {
				writer.On("AppendHeader", headers).Return()
				writer.On("AppendRows", []table.Row{{1, 1, "test", DefaultEmptyValue}}).Return()
				writer.On("RenderMarkdown").Return("")
			},
		},
		{
			name: "Unknown format renderer received alone object",
			fields: fields{
				writer: writer,
				options: RendererOptions{
					EmptyValue: DefaultEmptyValue,
					Format:     RenderFormat("unknown"),
				},
			},
			args: args{
				obj: rendererTestStruct{
					ID: 1,
					User: &renderNestedTestStruct{
						UserID: 1,
					},
					Name: "test",
				},
			},
			expect: func() {
				writer.On("AppendHeader", headers).Return()
				writer.On("AppendRows", []table.Row{{1, 1, "test", DefaultEmptyValue}}).Return()
				writer.On("Render").Return("")
			},
		},
		{
			name:   "passed object is not a struct",
			fields: fields{writer: writer},
			args: args{
				obj: 1,
			},
			expect: func() {},
			err:    ObjectIsNotStruct,
		},
	}
	for _, tt := range tests {
		tt.expect()
		t.Run(tt.name, func(t *testing.T) {
			r := &tableRenderer{
				writer:  tt.fields.writer,
				options: tt.fields.options,
			}
			_, err := r.Render(tt.args.obj)
			if !reflect.DeepEqual(err, tt.err) {
				t.Errorf("Render() got = %v, want %v", err, tt.err)
			}
		})
	}
}

func TestNewRenderer(t *testing.T) {
	type args struct {
		writer  Writer
		options []RendererOptions
	}
	writer := mocks.NewWriterMock(t)
	tests := []struct {
		name string
		args args
		want Renderer
	}{
		{
			name: "without options",
			args: args{
				writer: writer,
			},
			want: &tableRenderer{
				writer:  writer,
				options: RendererOptions{EmptyValue: DefaultEmptyValue, Format: DefaultRenderFormat},
			},
		},
		{
			name: "with options",
			args: args{
				writer:  writer,
				options: []RendererOptions{{EmptyValue: "empty", Format: DefaultRenderFormat}},
			},
			want: &tableRenderer{
				writer:  writer,
				options: RendererOptions{EmptyValue: "empty", Format: DefaultRenderFormat},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewRenderer(tt.args.writer, tt.args.options...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewRenderer() = %v, want %v", got, tt.want)
			}
		})
	}
}
