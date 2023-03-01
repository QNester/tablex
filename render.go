package tablex

import (
	"github.com/jedib0t/go-pretty/v6/table"
)

const DefaultEmptyValue = "null"
const RenderFormatTable RenderFormat = "table"
const RenderFormatCSV RenderFormat = "csv"
const RenderFormatHTML RenderFormat = "html"
const RenderFormatMD RenderFormat = "markdown"
const DefaultRenderFormat = RenderFormatTable

type RenderFormat string

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

//go:generate mockery --output ./mock --name Renderer --structname RendererMock
type Renderer interface {
	Render(interface{}) string
}

type RendererOptions struct {
	EmptyValue string
	Format     RenderFormat
}

type tableRenderer struct {
	writer  Writer
	options RendererOptions
}

func NewRenderer(writer table.Writer, options ...RendererOptions) Renderer {
	var renderOptions RendererOptions

	if len(options) != 0 {
		renderOptions = options[0]
	} else {
		renderOptions = RendererOptions{EmptyValue: DefaultEmptyValue, Format: DefaultRenderFormat}
	}

	return &tableRenderer{writer: writer, options: renderOptions}
}

func (r *tableRenderer) Render(obj interface{}) string {
	r.appendData(obj)
	return r.writerRender()
}

func (r *tableRenderer) appendData(obj interface{}) {
	tInfo := newTablexInfo(obj, r.options.EmptyValue)

	r.writer.AppendHeader(tInfo.headers)
	r.writer.AppendRows(tInfo.rowsForObject(obj))
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
