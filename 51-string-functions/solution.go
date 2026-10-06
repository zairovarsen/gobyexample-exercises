//go:build solution

package stringfunctions

import "strings"

func FileKind(name string) string {
	switch {
	case strings.HasPrefix(name, "."):
		return "hidden"
	case strings.Contains(name, "_gen."):
		return "generated"
	case strings.HasSuffix(name, "_test.go"):
		return "test"
	case strings.HasSuffix(name, ".go"):
		return "source"
	}
	return "other"
}

func Slug(title string) string {
	return strings.ReplaceAll(strings.ToLower(title), " ", "-")
}

func CountWord(text, word string) int {
	return strings.Count(strings.ToLower(text), strings.ToLower(word))
}

func Underline(title string) string {
	return title + "\n" + strings.Repeat("=", len(title))
}

func Domain(address string) string {
	i := strings.Index(address, "@")
	if i < 0 {
		return ""
	}
	return address[i+1:]
}

func CleanList(list string) string {
	var items []string
	for _, item := range strings.Split(list, ",") {
		item = strings.ToLower(strings.TrimSpace(item))
		if item != "" {
			items = append(items, item)
		}
	}
	return strings.Join(items, ",")
}
