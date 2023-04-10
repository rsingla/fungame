package main

import (
	"fmt"
	"strings"
	"time"

	"github.com/johnfercher/maroto/pkg/color"

	"github.com/johnfercher/maroto/pkg/consts"
	"github.com/johnfercher/maroto/pkg/pdf"
	"github.com/johnfercher/maroto/pkg/props"
)

func generatePDF(folder string, filename string) {
	fmt.Println("PDF generation is in progress...")

	m := pdf.NewMaroto(consts.Portrait, consts.A4)
	m.SetAliasNbPages("5")
	m.SetPageMargins(10, 10, 10)
	m.SetFirstPageNb(1)
	buildHeading(m)
	buildBodyData(m)

	data := randomContractData()

	wordCount := 100

	dataArr := strings.Split(data, " ")
	myData := ""
	for i := 0; i < len(dataArr); i++ {
		myData = ""
		lastIndex := 0

		if i+wordCount > len(dataArr) {
			lastIndex = len(dataArr)
		} else {
			lastIndex = i + wordCount
		}
		for j := i; j < lastIndex; j++ {
			myData = myData + dataArr[j] + " "
			i++
		}
		fmt.Println(myData)
		buildContractData(m, myData)
	}

	//setBuyerSignature(m)
	//setSellerSignature(m)

	timeNow := time.Now()
	m.SetCreationDate(timeNow)

	err := m.OutputFileAndClose(folder + "/" + filename + ".pdf")

	if err != nil {
		fmt.Println(err)
	}

}

func buildHeading(m pdf.Maroto) {

	heading := "Zapitel"

	m.RegisterHeader(func() {
		m.Row(12, func() {
			m.Col(4, func() {
				err := m.FileImage("public/zapitel_logo.png", props.Rect{
					Top:     1,
					Left:    1,
					Center:  true,
					Percent: 500,
				})
				if err != nil {
					fmt.Println(err)
				}
			})
		})

	})

	m.Row(10, func() {
		m.Col(12, func() {
			m.Text(heading, props.Text{
				Top:    1,
				Style:  consts.Bold,
				Size:   15,
				Align:  consts.Center,
				Color:  color.NewBlack(),
				Family: consts.Courier,
			})
		})
	})

}

func buildBodyData(m pdf.Maroto) {

	//m.SetBackgroundColor(getTealColor())

	m.Row(12, func() {
		m.Col(12, func() {
			m.Text("Contract Info", props.Text{
				Top:    2,
				Style:  consts.Bold,
				Size:   13,
				Align:  consts.Center,
				Color:  color.NewBlack(),
				Family: consts.Courier,
			})
		})
	})
}

func buildContractData(m pdf.Maroto, data string) {

	m.Row(float64(15), func() {
		m.Col(12, func() {
			m.Text(data, props.Text{
				Size:            10.0,
				Style:           consts.BoldItalic,
				Family:          consts.Courier,
				Align:           consts.Center,
				Top:             1.0,
				Left:            1.0,
				Right:           1.0,
				Extrapolate:     false,
				VerticalPadding: 1.0,
				Color: color.Color{
					Red:   10,
					Green: 20,
					Blue:  30,
				},
			})
		})
	})

}

func setBuyerSignature(m pdf.Maroto) {
	m.Row(15, func() {
		m.Col(12, func() {
			m.Signature("Buyer Signature", props.Font{
				Size:   12.0,
				Style:  consts.BoldItalic,
				Family: consts.Courier,
				Color: color.Color{
					Red:   10,
					Green: 20,
					Blue:  30,
				},
			})
		})
	})
}

func setSellerSignature(m pdf.Maroto) {
	m.Row(15, func() {
		m.Col(12, func() {
			m.Signature("Seller Signature", props.Font{
				Size:   12.0,
				Style:  consts.BoldItalic,
				Family: consts.Courier,
				Color: color.Color{
					Red:   10,
					Green: 20,
					Blue:  30,
				},
			})
		})
	})
}

func getDarkGrayColor() color.Color {
	myColor := color.Color{
		Red:   88,
		Green: 80,
		Blue:  180,
	}

	return myColor
}

func getTealColor() color.Color {
	myColor := color.Color{
		Red:   0,
		Green: 128,
		Blue:  128,
	}

	return myColor
}
