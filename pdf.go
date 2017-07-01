package main

import (
	"fmt"

	"github.com/jung-kurt/gofpdf"
)

const (
	TEXT_CHARNAME      = iota
	TEXT_CLASSANDLEVEL = iota
	TEXT_BACKGROUND    = iota
	TEXT_PLAYERNAME    = iota
	TEXT_RACE          = iota
	TEXT_ALIGNMENT     = iota
	TEXT_XP            = iota
)

type textInfo struct {
	size float64
	x    float64
	y    float64
}

var textLocations = map[int]textInfo{
	TEXT_CHARNAME:      textInfo{size: 16, x: 20, y: 19},
	TEXT_CLASSANDLEVEL: textInfo{size: 12, x: 92, y: 15},
	TEXT_BACKGROUND:    textInfo{size: 12, x: 131, y: 15},
	TEXT_PLAYERNAME:    textInfo{size: 12, x: 164, y: 15},
	TEXT_RACE:          textInfo{size: 12, x: 92, y: 25},
	TEXT_ALIGNMENT:     textInfo{size: 12, x: 131, y: 25},
	TEXT_XP:            textInfo{size: 12, x: 164, y: 25},
}

func initializePDF() *gofpdf.Fpdf {
	pdf := gofpdf.New("P", "mm", "A4", "")
	pdf.AddPage()
	pdf.SetFont("Arial", "", 12)

	return pdf
}

func addBackground(pdf *gofpdf.Fpdf, imagepath string) {
	width, height, _ := pdf.PageSize(0)

	pdf.ImageOptions(imagepath, 0, 0, width, height, false, gofpdf.ImageOptions{ImageType: "PNG", ReadDpi: true}, 0, "")
}

func setName(pdf *gofpdf.Fpdf, name string) {
	addText(pdf, name, textLocations[TEXT_CHARNAME])
}

func setClassAndLevel(pdf *gofpdf.Fpdf, class string, level int) {
	addText(pdf, fmt.Sprintf("%v %v", class, level), textLocations[TEXT_CLASSANDLEVEL])
}

func setBackground(pdf *gofpdf.Fpdf, background string) {
	addText(pdf, background, textLocations[TEXT_BACKGROUND])
}

func setPlayerName(pdf *gofpdf.Fpdf, playerName string) {
	addText(pdf, playerName, textLocations[TEXT_PLAYERNAME])
}

func setRace(pdf *gofpdf.Fpdf, race string) {
	addText(pdf, race, textLocations[TEXT_RACE])
}

func setAlignment(pdf *gofpdf.Fpdf, alignment string) {
	addText(pdf, alignment, textLocations[TEXT_ALIGNMENT])
}

func setXP(pdf *gofpdf.Fpdf, XP int) {
	addText(pdf, fmt.Sprintf("%v", XP), textLocations[TEXT_XP])
}

func addText(pdf *gofpdf.Fpdf, txtStr string, info textInfo) {
	pdf.SetXY(info.x, info.y)
	pdf.SetFontSize(info.size)
	pdf.Write(info.size, txtStr)
}

func generatePDF(pdf *gofpdf.Fpdf, filePath string) {
	pdf.OutputFileAndClose(filePath)
}
