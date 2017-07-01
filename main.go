package main

func main() {
	pdf := initializePDF()

	addBackground(pdf, "page1.png")

	setName(pdf, "Peren Goodharvest")
	setClassAndLevel(pdf, "Cleric", 2)
	setBackground(pdf, "Hermit")
	setPlayerName(pdf, "No One")
	setRace(pdf, "Wood Elf")
	setAlignment(pdf, "Lawful Good")
	setXP(pdf, 100)

	generatePDF(pdf, "hello.pdf")
}
