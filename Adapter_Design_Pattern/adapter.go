package main

import "fmt"
type OldPrinter interface {
    Print(s string) string
}

type MyOldPrinter struct{}

func (p *MyOldPrinter) Print(s string) string {
    return "Legacy Printer: " + s
}

type NewPrinter interface {
    PrintFormatted(s string) string
}

type PrinterAdapter struct {
    OldPrinter OldPrinter
}

func (p *PrinterAdapter) PrintFormatted(s string) string {
    return p.OldPrinter.Print(s)
}
func main() {
    oldPrinter := &MyOldPrinter{}
    adapter := &PrinterAdapter{OldPrinter: oldPrinter}
    fmt.Println(adapter.PrintFormatted("Hello, World!"))
}