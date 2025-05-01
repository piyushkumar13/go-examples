package main

import "fmt"

type Printer interface {
	Print()
}

type Scanner interface {
	Scan(s string)
}

type PrinterScanner interface {
	Printer
	Scanner
}

type PrinterScannerImpl struct {
	content string
}

func (p *PrinterScannerImpl) Print() {

	fmt.Println("Printing ::: ", p.content)
}

func (p *PrinterScannerImpl) Scan(s string) {
	fmt.Println("Scanning ::: ", p.content+s)
}

func main() {

	ps := &PrinterScannerImpl{content: "This is the content for printing and scanning"}

	ps.Print()
	ps.Scan(":::Scan::::")

}
