package service

import (
	"EuprvaSaobracajna/data"
	"bytes"
	"fmt"
	"github.com/jung-kurt/gofpdf"
	"log"
	"strconv"
)

func GeneratePdf(nalog data.PrekrsajniNalog) ([]byte, error) {
	pdf := gofpdf.New("P", "mm", "A4", "")
	pdf.AddPage()
	pdf.AddUTF8Font("NotoSans", "", "./fonts/NotoSans-Regular.ttf")
	pdf.SetFont("NotoSans", "", 24)
	pdf.CellFormat(0, 20, "Prekršajni nalog:", "", 1, "C", false, 0, "")
	pdf.Ln(4)
	pdf.SetFont("NotoSans", "", 16)
	pdf.CellFormat(0, 10, fmt.Sprintf("Izdato dnana: %s", nalog.Datum.Format("2006-01-02")), "1", 1, "", false, 0, "")
	pdf.Ln(12)
	pdf.SetFont("NotoSans", "", 12)
	pdf.CellFormat(0, 50, fmt.Sprintf("Opis: %s", nalog.Opis), "1", 1, "T", false, 0, "")
	pdf.Ln(8)
	pdf.SetFont("NotoSans", "", 16)
	pdf.CellFormat(0, 10, fmt.Sprintf("Izdato od strane: %s,   JMBG Zaposlenog: %s", nalog.IzdatoOdStrane, nalog.JMBGSluzbenika), "1", 1, "", false, 0, "")
	pdf.Ln(8)
	pdf.CellFormat(0, 10, fmt.Sprintf("Izdato za: %s,   JMBG Zapisanog: %s", nalog.IzdatoZa, nalog.JMBGZapisanog), "1", 1, "", false, 0, "")
	pdf.Ln(8)
	if nalog.Vrednost != nil {
		pdf.CellFormat(0, 10, fmt.Sprintf("Tip prekršaja: %s,   Jedinica mere: %s,   Vrednost: %s", formateType(nalog.TipPrekrsaja), formateUnit(nalog.JedinicaMere), strconv.Itoa(*nalog.Vrednost)), "1", 1, "", false, 0, "")
	} else {
		pdf.CellFormat(0, 10, fmt.Sprintf("Tip prekršaja: %s,   Jedinica mere: %s,   Vrednost: %s", formateType(nalog.TipPrekrsaja), formateUnit(nalog.JedinicaMere), "n/a"), "1", 1, "", false, 0, "")

	}
	pdf.Ln(8)
	if len(nalog.Slike) > 0 {
		pdf.CellFormat(0, 20, "Slike:", "", 1, "C", false, 0, "")
		pdf.Ln(2)

		for i := 0; i < len(nalog.Slike); i++ {
			pdf.CellFormat(0, 10, fmt.Sprintf("%s", nalog.Slike[i]), "", 1, "", false, 0, fmt.Sprintf("http://localhost:8001/api/files/%s", nalog.Slike[i]))
			pdf.Ln(4)
		}
	}
	var buf bytes.Buffer
	if err := pdf.Output(&buf); err != nil {
		log.Fatal(err)
		return nil, err
	}

	return buf.Bytes(), nil
}

func formateType(nalogType string) string {
	switch nalogType {
	case "POJAS":
		return "Pojas nije vezan"
	case "PREKORACENJE_BRZINE":
		return "Prekoračenje brzine"
	case "PIJANA_VOZNJA":
		return "Vožnja pod dejstvom alkohola"
	case "TEHNICKA_NEISPRAVNOST":
		return "Tehnička neispravnost"
	case "PRVA_POMOC":
		return "Ne poseduje prvu pomoć"
	case "NEMA_VOZACKU":
		return "Ne poseduje vozačku dozvolu"
	case "REGISTRACIJA":
		return "Istekla registracija"
	default:
		return nalogType
	}
}

func formateUnit(unit *string) string {
	if unit != nil {
		return *unit
	}
	return "n/a"
}
