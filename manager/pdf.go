package manager

import (
	"fmt"
	"io/ioutil"

	"github.com/jung-kurt/gofpdf"
)

const font = "Arial"
const headerFile = "header.txt"
const footerFile = "footer.txt"

func loadHeaderFooter() (h, f string, err error) {
	hr, err := ioutil.ReadFile(headerFile)
	if err != nil {
		return
	}

	fr, err := ioutil.ReadFile(footerFile)

	h = string(hr)
	f = string(fr)
	return
}

func (i Invoice) PDF(filename string) error {
	header, footer, err := loadHeaderFooter()
	if err != nil {
		return err
	}

	pdf := gofpdf.New("P", "mm", "A4", "")
	pdf.SetFont(font, "", 12)
	pdf.AddPage()
	tr := pdf.UnicodeTranslatorFromDescriptor("")

	pdf.SetXY(20, 20)
	pdf.MultiCell(0, 5, tr(header), "", "", false)
	pdf.Ln(10)
	pdf.MultiCell(180, 5, tr(i.Client), "", "R", false)

	pdf.SetXY(20, 100)
	pdf.SetFontSize(20)
	titleStr := "Facture"
	if i.Quote {
		titleStr = "Devis"
	}

	pdf.Write(0, tr(fmt.Sprintf("%s %06d", titleStr, i.ID)))

	pdf.SetXY(20, 110)
	pdf.SetFontSize(10)

	var dateStr string
	if !i.Quote {
		dateStr += "Établie le : " + i.Emitted.Format("02/01/2006") + "\n"
		dateStr += "Fin des prestations le : " + i.Delivered.Format("02/01/2006") + "\n"
		dateStr += "Échéance de paiement : " + i.Emitted.AddDate(0, 0, int(i.PaymentDays)).Format("02/01/2006")
	} else {
		dateStr += "Établi le : " + i.Emitted.Format("02/01/2006") + "\n"
		dateStr += "Durée de validité : 7 jours\n"
	}

	pdf.MultiCell(0, 5, tr(dateStr), "", "", false)

	lineNumber := 0
	makeLine := func(description, quantity, unitCost, amount string) {
		if lineNumber == 0 {
			pdf.SetFont(font, "B", 12)
		} else {
			pdf.SetFont(font, "", 12)
		}

		pdf.SetX(20)
		pdf.Cell(80, 5, tr(description))
		pdf.CellFormat(30, 5, tr(quantity), "", 0, "R", false, 0, "")
		pdf.CellFormat(30, 5, tr(unitCost), "", 0, "R", false, 0, "")
		pdf.CellFormat(30, 5, tr(amount), "", 0, "R", false, 0, "")

		if lineNumber == 0 {
			pdf.Ln(6)
			pdf.Line(20, pdf.GetY(), 190, pdf.GetY())
			pdf.Ln(2)
		} else {
			pdf.Ln(6)
		}
		lineNumber++
	}

	pdf.Ln(10)
	makeLine("Description", "Quantité", "Coût Unit.", "Montant")

	for _, s := range i.Services {
		makeLine(s.format(i.Currency))
	}

	pdf.Ln(10)
	pdf.SetFont(font, "B", 12)
	pdf.MultiCell(180, 5, tr("Total : "+formatMoney(i.Total(), i.Currency)), "", "R", false)

	pdf.SetFont(font, "", 10)
	pdf.MultiCell(180, 5, tr(i.Comment), "", "R", false)

	pdf.SetY(260)
	pdf.SetFont(font, "", 8)
	pdf.MultiCell(0, 3, tr(footer), "", "C", false)

	return pdf.OutputFileAndClose(filename)
}
