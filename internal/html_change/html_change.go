package html_change

import (
	"fmt"
	"log"
	"os"
	Utils "warehouseWeb/internal/searchStruct"
)

func deleteLastChar(s string) string {
	sz := len(s)
	if sz > 0 {
		s = s[:sz-1]
	}
	return s
}

func WriteListNotFound(fileName string, source *Utils.SearchResults) {
	f, err := os.OpenFile(fileName, os.O_WRONLY, 0666)
	if err != nil {
		fmt.Println("Error occurred when opening file.", err.Error())
		return
	}

	if err := os.Truncate(fileName, 0); err != nil {
		log.Printf("Failed to truncate: %v", err)
	}
	defer f.Close()

	tabs := ""
	_, err = f.WriteString(tabs + "<html>\n")
	tabs = tabs + "\t"

	_, _ = f.WriteString(tabs + "<h1>Товары не найдены.</h1>\n")

	tabs = deleteLastChar(tabs)
	_, _ = f.WriteString(tabs + "</html>\n")
}

func WriteList(fileName string, source *Utils.SearchResults) {
	f, err := os.OpenFile(fileName, os.O_WRONLY, 0666)
	if err != nil {
		fmt.Println("Error occurred when opening file.", err.Error())
		return
	}

	if err := os.Truncate(fileName, 0); err != nil {
		log.Printf("Failed to truncate: %v", err)
	}
	defer f.Close()

	tabs := ""
	_, err = f.WriteString(tabs + "<html>\n")
	tabs = tabs + "\t"

	_, _ = f.WriteString(tabs + "<ol>\n")
	tabs = tabs + "\t"

	createTable(f, source, tabs)

	tabs = deleteLastChar(tabs)
	_, _ = f.WriteString(tabs + "</ol>\n")

	tabs = deleteLastChar(tabs)
	_, _ = f.WriteString(tabs + "</html>\n")

}

func createTable(f *os.File, source *Utils.SearchResults, tabs string) {
	_, _ = f.WriteString(tabs + "<form>\n")
	tabs = tabs + "\t"
	_, _ = f.WriteString(tabs + "<table>\n")
	tabs = tabs + "\t"

	for _, pp := range *source {
		_, _ = f.WriteString(tabs + "<tr>\n")
		_, _ = f.WriteString(tabs + "<td>\n")
		str := fmt.Sprintf("Name: %s id: %d placement_id %d", pp.Name(), pp.IdUniq(), pp.Place())
		_, _ = f.WriteString(tabs + "<li> " + str + " </li>\n")
		_, _ = f.WriteString(tabs + "</td>\n")
		addButtons(f, tabs, "Выдать товар", "\"/buttonGive\"")
		addButtons(f, tabs, "Списать товар", "")
		addButtons(f, tabs, "Изменить местоположение товара", "")
		_, _ = f.WriteString(tabs + "</tr>\n")
	}

	tabs = deleteLastChar(tabs)
	_, _ = f.WriteString(tabs + "</table>\n")
	tabs = deleteLastChar(tabs)
	_, _ = f.WriteString(tabs + "</form>\n")

}

func addButtons(f *os.File, tabs string, nameButton string, formaction string) {
	_, _ = f.WriteString(tabs + "<td>")
	_, _ = f.WriteString("<button" + " formaction=" + formaction + " >" + nameButton + "</button>")
	_, _ = f.WriteString(tabs + "</td>\n")
}
