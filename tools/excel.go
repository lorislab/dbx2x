package tools

import (
	"os"
	"path/filepath"

	log "github.com/sirupsen/logrus"
	"github.com/xuri/excelize/v2"
)

func ConvertExcel2Xml(file, output string) {
	ext := filepath.Ext(file)
	if ext == ".xls" {
		log.WithField("file", file).Error("Not supported format. Please convert xls to xlsx format")
	} else {
		convertXlsx2Xml(file, output)
	}
}

func convertXlsx2Xml(file, output string) {

	// open excel file
	excel, err := excelize.OpenFile(file)
	if err != nil {
		log.WithField("file", file).Fatal(err)
	}
	defer func() {
		// Close the spreadsheet.
		if err := excel.Close(); err != nil {
			log.WithField("file", file).Error(err)
		}
	}()

	content := `<?xml version="1.0" encoding="UTF-8"?>` + "\n<dataset>\n"

	for _, sheet := range excel.GetSheetMap() {

		rows, err := excel.GetRows(sheet)
		if err != nil {
			log.WithFields(log.Fields{"sheet": sheet}).Error(err)
			continue
		}

		if len(rows) <= 1 {
			log.WithFields(log.Fields{"sheet": sheet}).Warn("No rows found for the sheet")
			continue
		}

		content += "\t<!-- " + sheet + " -->\n"
		header := rows[0]
		for ri := 1; ri < len(rows); ri++ {
			row := rows[ri]
			xml_row := "\t<" + sheet
			for hi, colCell := range row {
				xml_row += ` ` + header[hi] + `="` + colCell + `"`
			}
			xml_row += "/>\n"
			content += xml_row
		}
		content += "\n"
	}

	content += "</dataset>\n"

	writeOutput(output, content)
}

func writeOutput(output, content string) {
	// create output file
	err := os.WriteFile(output, []byte(content), 0644)
	if err != nil {
		log.WithField("file", output).Fatal(err)
	}
}
