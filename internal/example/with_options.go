package main

import (
	"github.com/QNester/tablex"
	"github.com/jedib0t/go-pretty/v6/table"
)

type exampleStruct struct {
	ID   int    `tablex:"header:#"`
	Name string `tablex:"header:name"`
	Age  *int   `tablex:"header:age"`
}

func main() {
	writer := table.NewWriter()
	renderer := tablex.NewRenderer(
		writer,
		tablex.RendererOptions{
			EmptyValue: "no data",               // what you wanna see if some of your fields' data equals nil
			Format:     tablex.RenderFormatHTML, // rendering format
		},
	)
	data := exampleStruct{ID: 1, Name: "Joh Watson", Age: nil}
	renderedData := renderer.Render(data)
	print(renderedData)
}

// out:
//
// <table class="go-pretty-table">
// 		<thead>
// 			<tr>
// 				<th align="right">#</th>
// 				<th>name</th>
// 				<th>age</th>
// 			</tr>
// 		</thead>
// 		<tbody>
// 			<tr>
// 				<td align="right">1</td>
// 				<td>Joh Watson</td>
// 				<td>no data</td>
// 			</tr>
//		</tbody>
// </table>
