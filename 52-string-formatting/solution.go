//go:build solution

package stringformatting

import (
	"fmt"
	"io"
)

type Point struct {
	X, Y int
}

type Item struct {
	Name  string
	Qty   int
	Price float64
}

func DescribePoint(p Point) string {
	return fmt.Sprintf("v=%v +v=%+v #v=%#v T=%T", p, p, p, p)
}

func DescribeInt(n int) string {
	return fmt.Sprintf("dec=%d bin=%b hex=%x char=%c", n, n, n, n)
}

func DescribeFloat(f float64) string {
	return fmt.Sprintf("f=%f 2f=%.2f e=%e E=%E", f, f, f, f)
}

func DescribeString(s string) string {
	return fmt.Sprintf("s=%s q=%q x=%x", s, s, s)
}

func TableRow(name string, qty int, price float64) string {
	return fmt.Sprintf("|%-10s|%5d|%8.2f|", name, qty, price)
}

func Receipt(w io.Writer, items []Item) {
	total := 0.0
	for i, item := range items {
		lineTotal := float64(item.Qty) * item.Price
		total += lineTotal
		fmt.Fprintf(w, "%03d %-10s %3d %8.2f\n", i+1, item.Name, item.Qty, lineTotal)
	}
	fmt.Fprintf(w, "%-18s %8.2f\n", "TOTAL", total)
}
