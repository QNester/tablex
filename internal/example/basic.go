package main

import (
	"github.com/QNester/tablex"
	"github.com/jedib0t/go-pretty/v6/table"
)

type aloneStruct struct {
	ID   int    `tablex:"header:#"`
	Name string `tablex:"header:name"`
}

func main() {
	writer := table.NewWriter()
	renderer := tablex.NewRenderer(writer)
	data := aloneStruct{ID: 1, Name: "Joh Watson"}
	renderedData := renderer.Render(data)
	print(renderedData)
}

// out:
//
// +---+------------+
// | # | NAME       |
// +---+------------+
// | 1 | Joh Watson |
// +---+------------+
