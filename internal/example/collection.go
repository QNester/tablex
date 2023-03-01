package main

import (
	"os"

	"github.com/QNester/tablex"
	"github.com/jedib0t/go-pretty/v6/table"
)

type nestedData struct {
	UserID int `tablex:"header:user_id"`
}

type myData struct {
	ID       int         `tablex:"header:id"`
	User     *nestedData `tablex:"header:user"`
	Name     string      `tablex:"header:name"`
	NoHeader string      `tablex:"header:"`
	NoTablex string
}

func main() {
	writer := table.NewWriter()
	writer.SetOutputMirror(os.Stdout) // we can print table to STDOUT automatically via go-pretty/table
	renderer := tablex.NewRenderer(writer)

	data := []myData{
		{ID: 1, User: &nestedData{UserID: 1}, Name: "John Watson", NoHeader: "yes", NoTablex: "yes"},
		{ID: 2, User: &nestedData{UserID: 2}, Name: "Marry Watson", NoHeader: "yes", NoTablex: "yes"},
		{ID: 3, User: nil, Name: "Userless", NoHeader: "yes", NoTablex: "yes"},
	}

	renderer.Render(data)
}

// out
// +----+----------------+--------------+
// | ID | USER / USER_ID | NAME         |
// +----+----------------+--------------+
// |  1 | 1              | John Watson  |
// |  2 | 2              | Marry Watson |
// |  3 | null           | Userless     |
// +----+----------------+--------------+
