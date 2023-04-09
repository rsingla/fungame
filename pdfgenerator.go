package main

import (
	"fmt"

	"github.com/johnfercher/maroto/pkg/consts"
	"github.com/johnfercher/maroto/pkg/pdf"
)

func generatePDF(folder string, filename string) {
	fmt.Println("PDF generation is in progress...")

	m := pdf.NewMaroto(consts.Portrait, consts.A4)
	m.SetPageMargins(20, 10, 10)

	err := m.OutputFileAndClose(folder + "/" + filename + ".pdf")

	if err != nil {
		fmt.Println(err)
	}

}
