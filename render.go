package tablex

import (
	"github.com/jedib0t/go-pretty/v6/table"
)

// DefaultEmptyValue default value for nullable columns
const DefaultEmptyValue = "null"

const RenderFormatTable RenderFormat = "table"
const RenderFormatCSV RenderFormat = "csv"
const RenderFormatHTML RenderFormat = "html"
const RenderFormatMD RenderFormat = "markdown"
const DefaultRenderFormat = RenderFormatTable

type RenderFormat string

// Writer is a slice of table.Writer interface
//
//go:generate mockery --output ./mock --name Writer --structname WriterMock
type Writer interface {
	AppendRow(row table.Row, configs ...table.RowConfig)
	AppendRows(rows []table.Row, configs ...table.RowConfig)
	Render() string
	RenderCSV() string
	RenderHTML() string
	RenderMarkdown() string
	AppendFooter(row table.Row, configs ...table.RowConfig)
	AppendHeader(row table.Row, configs ...table.RowConfig)
}

// Renderer is a main interface of the tablex package
//
//go:generate mockery --output ./mock --name Renderer --structname RendererMock
type Renderer interface {
	Render(interface{}) (string, error)
}

// RendererOptions is specific options for tablex behaviour
type RendererOptions struct {
	EmptyValue string
	Format     RenderFormat
}

type tableRenderer struct {
	writer  Writer
	options RendererOptions
}

// NewRenderer returns an structure implemented Renderer interface
// It dependents on Writer interfaces (part of interface in https://github.com/jedib0t/go-pretty/tree/main/table)
func NewRenderer(writer Writer, options ...RendererOptions) Renderer {
	var renderOptions RendererOptions

	if len(options) != 0 {
		renderOptions = options[0]
	} else {
		renderOptions = RendererOptions{EmptyValue: DefaultEmptyValue, Format: DefaultRenderFormat}
	}

	return &tableRenderer{writer: writer, options: renderOptions}
}

// Render is a function for build table from structs or collections.
func (r *tableRenderer) Render(obj interface{}) (string, error) {
	err := r.appendData(obj)
	if err != nil {
		return "", err
	}

	return r.writerRender(), nil
}

func (r *tableRenderer) appendData(obj interface{}) error {
	tInfo, err := newTablexInfo(obj, r.options.EmptyValue)
	if err != nil {
		return err
	}

	r.writer.AppendHeader(tInfo.headers)
	r.writer.AppendRows(tInfo.rowsForObject(obj))

	return nil
}

func (r *tableRenderer) writerRender() string {
	switch r.options.Format {
	case RenderFormatCSV:
		return r.writer.RenderCSV()
	case RenderFormatHTML:
		return r.writer.RenderHTML()
	case RenderFormatMD:
		return r.writer.RenderMarkdown()
	default:
		return r.writer.Render()
	}
}
