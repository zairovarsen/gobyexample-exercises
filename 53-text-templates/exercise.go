//go:build !solution

// Package texttemplates contains exercises for https://gobyexample.com/text-templates
//
// Run the tests for this section with:
//
//	go test ./53-text-templates/
package texttemplates

import "io"

// All exercises use the standard "text/template" package: parse a template
// text with template.New(name).Parse(text), then call Execute(w, data) on the
// result to write the output for some data.

// Person is used by NameTag.
type Person struct {
	Name string
	Age  int
}

// Order is used by WriteOrder.
type Order struct {
	Customer string
	Items    []OrderItem
	Paid     bool
}

// OrderItem is one position of an Order.
type OrderItem struct {
	Name string
	Qty  int
}

// RenderValue executes the template "Value is {{.}}\n" for value and writes
// the result to w (a real program would pass os.Stdout). The {{.}} action is
// replaced by the value the template is executed with.
//
// Example: RenderValue(w, 42) writes "Value is 42\n".
//
// The template text is a constant that is known to be valid, so this is a
// good place for template.Must, which panics when parsing fails.
func RenderValue(w io.Writer, value any) error {
	// TODO: parse the template, execute it into w and return Execute's error.
	return nil
}

// Render parses text as a template, executes it for data and returns the
// output as a string. Unlike RenderValue it gets its template from the
// caller, so it must not panic: when parsing or executing fails it returns
// an empty string and the error.
//
// Examples:
//
//	Render("Hi {{.Name}}!", Person{Name: "Ann"})              -> "Hi Ann!"
//	Render("Hi {{.name}}!", map[string]string{"name": "Ann"}) -> "Hi Ann!"
//	Render("Hi {{.Name", Person{Name: "Ann"})                 -> "" and an error
//
// Hint: Execute needs an io.Writer; a strings.Builder or bytes.Buffer
// collects the output in memory.
func Render(text string, data any) (string, error) {
	// TODO: parse, execute into a buffer, return the buffer's content.
	return "", nil
}

// NameTag returns "Hello, my name is <Name> and I am <Age> years old." for
// p, using a template that accesses the struct fields with {{.FieldName}}.
//
// Example: NameTag(Person{"Jane Doe", 30}) returns
// "Hello, my name is Jane Doe and I am 30 years old."
func NameTag(p Person) (string, error) {
	// TODO: write the template text and execute it (Render can help).
	return "", nil
}

// YesNo returns "yes" or "no" depending on how a template's
// {{if .}}...{{else}}...{{end}} judges v. In a template a value is false
// when it is the zero value of its type (false, 0, "", a nil pointer) or an
// empty slice or map; everything else is true.
//
// Examples: YesNo("not empty") -> "yes", YesNo("") -> "no", YesNo(0) -> "no".
func YesNo(v any) (string, error) {
	// TODO: execute a template with if/else for v.
	return "", nil
}

// BulletList returns one line per item of the form "- <item>\n", using the
// template action {{range .}}...{{end}}. Inside the range block {{.}} is the
// current item.
//
// Example: BulletList([]string{"Go", "Rust"}) returns "- Go\n- Rust\n".
// An empty list gives the empty string.
func BulletList(items []string) (string, error) {
	// TODO: execute a template that ranges over the items.
	return "", nil
}

// WriteOrder writes a summary of o to w with a single template that
// combines fields, range and if/else:
//
//	Order for Jane
//	- 2 x apple
//	- 1 x pear
//	Status: paid
//
// The first line names o.Customer, then there is one line per item with its
// Qty and Name, and the last line says "Status: paid" or "Status: unpaid"
// depending on o.Paid. Every line ends with "\n"; there are no blank lines.
//
// Hint: a newline in the template text right after an action such as
// {{range .Items}} is copied to the output. Either keep such actions on the
// same line as the text around them, or trim the white space next to an
// action with {{- and -}}.
func WriteOrder(w io.Writer, o Order) error {
	// TODO: one template, executed into w.
	return nil
}
