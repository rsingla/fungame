package main

import (
	"fmt"

	"github.com/johnfercher/maroto/pkg/color"

	"github.com/johnfercher/maroto/pkg/consts"
	"github.com/johnfercher/maroto/pkg/pdf"
	"github.com/johnfercher/maroto/pkg/props"
)

func generatePDF(folder string, filename string) {
	fmt.Println("PDF generation is in progress...")

	m := pdf.NewMaroto(consts.Portrait, consts.A4)
	m.SetPageMargins(20, 10, 10)
	buildHeading(m)

	err := m.OutputFileAndClose(folder + "/" + filename + ".pdf")

	if err != nil {
		fmt.Println(err)
	}

}

func buildHeading(m pdf.Maroto) {

	heading := "Zapitel"

	m.RegisterHeader(func() {
		m.Row(4, func() {
			m.Col(4, func() {
				err := m.FileImage("public/zapitel_logo.png", props.Rect{
					Center:  true,
					Percent: 200,
				})
				if err != nil {
					fmt.Println(err)
				}
			})
		})

	})

	m.Row(4, func() {
		m.Col(8, func() {
			m.Text(heading, props.Text{
				Top:   0,
				Style: consts.Bold,
				Size:  15,
				Align: consts.Center,
				Color: getDarkGrayColor(),
			})
		})
	})

}

func getDarkGrayColor() color.Color {
	myColor := color.Color{
		Red:   180,
		Green: 180,
		Blue:  180,
	}

	return myColor
}
