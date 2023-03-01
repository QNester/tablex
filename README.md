# Tablex

[![Tests](https://github.com/QNester/tablex/actions/workflows/tests.yml/badge.svg?branch=main)](https://github.com/QNester/tablex/actions/workflows/tests.yml)

Render your structs with [go-pretty/table](https://github.com/jedib0t/go-pretty/tree/main/table) in go-way tags.

## Installation 

`go get https://github.com/QNester/tablex`

## Usage

1. Set up your table (see more in [go-pretty/table doc](https://github.com/jedib0t/go-pretty/tree/main/table))
2. Add `tablex` as a tag to your struct type
3. Render the table via Renderer interface.

```go
import (
  "github.com/QNester/tablex"
  "github.com/jedib0t/go-pretty/v6/table"
)

type myData struct {
  ID int `tablex:"header:#"`
  Name string `table:"Name"`
}

func main() {
  writer := table.NewWriter()
  renderer := tablex.NewRenderer(writer)
  table := renderer.Render()
  print(table)
}
```

Result of the running this program:
```text
+---+------------+
| # | NAME       |
+---+------------+
| 1 | Joh Watson |
+---+------------+
```

Show [examples](./internal/examples) for more usage information.

### Additional information
1. Use `tablex.RenderOptions` to set default empty value and table format (Text Table/CSV/HTML/MD)
```go
renderer := tablex.NewRenderer(
  writer,
  tablex.RendererOptions{
    EmptyValue: "no data",               // what you want to see if some of your fields' data equals nil
    Format:     tablex.RenderFormatHTML, // rendering format
  },
)
```
2. You can use `RendererMock` in your tests.