package main

import (
	"bytes"
	"github.com/SebastiaanKlippert/go-wkhtmltopdf"
	"github.com/julienschmidt/httprouter"
	"io"
	"log"
	"net/http"
)

// Generate takes in url as a query param and generates a document
func Generate(w http.ResponseWriter, req *http.Request, _ httprouter.Params) {
	url := req.URL.Query().Get("url")
	buf := generatePDF(url)
	w.Header().Set("Content-Disposition", "attachment; filename=doc.pdf")
	w.Header().Set("Content-Type", "application/pdf")
	io.Copy(w, buf)
}

func generatePDF(url string) *bytes.Buffer {
	pdfg, err := wkhtmltopdf.NewPDFGenerator()
	if err != nil {
		log.Fatal(err)
	}
	pdfg.AddPage(wkhtmltopdf.NewPage(url))
	err = pdfg.Create()
	if err != nil {
		log.Fatal(err)
	}
	return pdfg.Buffer()
}

func main() {
	router := httprouter.New()
	router.GET("/generate", Generate)
	log.Fatal(http.ListenAndServe(":8300", router))
}
