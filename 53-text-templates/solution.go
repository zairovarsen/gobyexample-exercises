//go:build solution

package texttemplates

import (
	"io"
	"strings"
	"text/template"
)

type Person struct {
	Name string
	Age  int
}

type Order struct {
	Customer string
	Items    []OrderItem
	Paid     bool
}

type OrderItem struct {
	Name string
	Qty  int
}

var (
	valueTmpl = template.Must(template.New("value").Parse("Value is {{.}}\n"))

	orderTmpl = template.Must(template.New("order").Parse(`Order for {{.Customer}}
{{range .Items -}}
- {{.Qty}} x {{.Name}}
{{end -}}
Status: {{if .Paid}}paid{{else}}unpaid{{end}}
`))
)

func RenderValue(w io.Writer, value any) error {
	return valueTmpl.Execute(w, value)
}

func Render(text string, data any) (string, error) {
	tmpl, err := template.New("render").Parse(text)
	if err != nil {
		return "", err
	}
	var sb strings.Builder
	if err := tmpl.Execute(&sb, data); err != nil {
		return "", err
	}
	return sb.String(), nil
}

func NameTag(p Person) (string, error) {
	return Render("Hello, my name is {{.Name}} and I am {{.Age}} years old.", p)
}

func YesNo(v any) (string, error) {
	return Render("{{if .}}yes{{else}}no{{end}}", v)
}

func BulletList(items []string) (string, error) {
	return Render("{{range .}}- {{.}}\n{{end}}", items)
}

func WriteOrder(w io.Writer, o Order) error {
	return orderTmpl.Execute(w, o)
}
