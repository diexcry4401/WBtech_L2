package pattern

import "fmt"

// Visitor - интерфейс посетителя, объявляющий методы для каждого типа элемента.
type Visitor interface {
	VisitPDF(*PDFDocument)
	VisitDocx(*DocxDocument)
}

// Element - интерфейс элемента, объявляющий метод Accept.
type Element interface {
	Accept(Visitor)
}

// PDFDocument - конкретный элемент PDF-документа.
type PDFDocument struct {
	Content string
}

// Accept - реализация метода Accept из интерфейса Element для PDF-документа.
func (pdf *PDFDocument) Accept(vis Visitor) {
	vis.VisitPDF(pdf)
}

// DOCXDocument - конкретный элемент DOCX-документа.
type DocxDocument struct {
	Content string
}

// Accept - реализация метода Accept из интерфейса Element для DOCX-документа.
func (docx *DocxDocument) Accept(vis Visitor) {
	vis.VisitDocx(docx)
}

// ExportVisitor - конкретный посетитель, реализующий операцию экспорта.
type ExportVisitor struct{}

// VisistPDF - реализация метода VisistPDF из интерфейса Visitor
func (e *ExportVisitor) VisitPDF(pdf *PDFDocument) {
	fmt.Printf("Export PDF content: %s\n", pdf.Content)
}

// VisistDocx - реализация метода VisistPDF из интерфейса Visitor
func (e *ExportVisitor) VisitDocx(docx *DocxDocument) {
	fmt.Printf("Export DOCX content: %s\n", docx.Content)
}

// func main() {

// 	// Создаем конкретные элементы.
// 	pdfFile := &PDFDocument{Content: "Some text from PDF"}
// 	docxFile := &DocxDocument{Content: "Some text from DOCX"}

// 	// Создаем конкретного посетителя.
// 	concreteExportVisitor := &ExportVisitor{}

// 	// Посещаем элементы с конкретным посетителем.
// 	pdfFile.Accept(concreteExportVisitor)
// 	docxFile.Accept(concreteExportVisitor)
// }
