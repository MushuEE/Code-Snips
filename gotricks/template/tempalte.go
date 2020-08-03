package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"os"
)

type Foo struct {
	Name string
	Type []*Bar
}

type Bar struct {
	Name2 string
	Type2 map[string]string
}

func main() {
	x := []Foo{
		{Name: "foo", Type: []*Bar{
			{Name2: "bar", Type2: map[string]string{"str1": "str1"}},
			{Name2: "bar2", Type2: map[string]string{"str2": "str2"}},
		},
		},
		{Name: "foo2", Type: []*Bar{
			{Name2: "2bar", Type2: map[string]string{"str1": "2str1"}},
			{Name2: "2bar2", Type2: map[string]string{"str2": "2str2"}},
		},
		},
	}
	fmt.Printf("Hello, playground: %+v", x)
	res1B, _ := json.Marshal(x)
	fmt.Println("")
	fmt.Println(string(res1B))

	const tpl = `
<!DOCTYPE html>
<html>
	<head>
		<meta charset="UTF-8">
		<title>Title</title>
	</head>
	<body>
		{{range .}}<div>{{ .Name }}|{{range .Type}}{{ .Name2 }}*{{index .Type2 "str1" }}|{{end}}</div>{{end}}
	</body>
</html>`

	t, _ := template.New("webpage").Parse(tpl)
	t.Execute(os.Stdout, x)

}
