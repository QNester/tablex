package tablex

import (
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
	NoHeader     string                  `tablex:"header:"`
	NoTablex     string
	privateField int
}

func Test_tableRenderer_Render(t *testing.T) {
	type fields struct {
		writer  Writer
		options RendererOptions
	}
	writer := mocks.NewWriterMock(t)
	type args struct {
		obj interface{}
	}
	tests := []struct {
		name   string
		fields fields
		expect func()
		args   args
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
				headers := table.Row{tableFieldName("id"), tableFieldName("user / user_id"), tableFieldName("name")}
				writer.On("AppendHeader", headers).Return()
				writer.On("AppendRows", []table.Row{{1, 1, "test"}}).Return()
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
				headers := table.Row{tableFieldName("id"), tableFieldName("user / user_id"), tableFieldName("name")}
				writer.On("AppendHeader", headers).Return()
				writer.On("AppendRows", []table.Row{{1, DefaultEmptyValue, "name"}}).Return()
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
					ID:   1,
					User: nil,
					Name: "name",
				},
			},
			expect: func() {
				headers := table.Row{tableFieldName("id"), tableFieldName("user / user_id"), tableFieldName("name")}
				writer.On("AppendHeader", headers).Return()
				writer.On("AppendRows", []table.Row{{1, "ooops, it's empty!", "name"}}).Return()
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
				headers := table.Row{tableFieldName("id"), tableFieldName("user / user_id"), tableFieldName("name")}
				writer.On("AppendHeader", headers).Return()
				writer.On("AppendRows", []table.Row{{1, 1, "test"}, {2, 2, "test_2"}}).Return()
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
				headers := table.Row{tableFieldName("id"), tableFieldName("user / user_id"), tableFieldName("name")}
				writer.On("AppendHeader", headers).Return()
				writer.On("AppendRows", []table.Row{{1, 1, "test"}}).Return()
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
				headers := table.Row{tableFieldName("id"), tableFieldName("user / user_id"), tableFieldName("name")}
				writer.On("AppendHeader", headers).Return()
				writer.On("AppendRows", []table.Row{{1, 1, "test"}}).Return()
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
				headers := table.Row{tableFieldName("id"), tableFieldName("user / user_id"), tableFieldName("name")}
				writer.On("AppendHeader", headers).Return()
				writer.On("AppendRows", []table.Row{{1, 1, "test"}}).Return()
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
				headers := table.Row{tableFieldName("id"), tableFieldName("user / user_id"), tableFieldName("name")}
				writer.On("AppendHeader", headers).Return()
				writer.On("AppendRows", []table.Row{{1, 1, "test"}}).Return()
				writer.On("Render").Return("")
			},
		},
	}
	for _, tt := range tests {
		tt.expect()
		t.Run(tt.name, func(t *testing.T) {
			r := &tableRenderer{
				writer:  tt.fields.writer,
				options: tt.fields.options,
			}
			r.Render(tt.args.obj)
		})
	}
}
