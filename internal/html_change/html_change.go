package html_change

import (
    "fmt"
    "log"
    "os"
	Utils "warehouseWeb/internal/searchStruct"
)
func deleteLastChar(s string) (string) {
	sz := len(s)
	if sz > 0 && s[sz-1] == '+' {
    	s = s[:sz-1]
	}
	return s
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

	for _, pp := range *source {
		str := fmt.Sprintf("Name: %s id: %d placement_id %d", pp.Name(), pp.IdUniq(), pp.Place())
		_, _ = f.WriteString(tabs + "<li> "+ str +" </li>\n")
	}

	tabs = deleteLastChar(tabs)
	_, _ = f.WriteString(tabs + "</ol>\n")
	
	tabs = deleteLastChar(tabs)
	_, _ = f.WriteString(tabs + "</html>\n")

}