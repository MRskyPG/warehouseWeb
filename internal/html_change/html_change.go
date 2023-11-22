package html_change

import (
	"fmt"
	"io"
	"os"
	"strconv"
	Utils "warehouseWeb/internal/searchStruct"
)

func deleteLastChar(s string) string {
	sz := len(s)
	if sz > 0 {
		s = s[:sz-1]
	}
	return s
}

func copy(source *os.File, destination *os.File) error {
	buf := make([]byte, 1024)
	for {
		n, err := source.Read(buf)
		if err != nil && err != io.EOF {
			return err
		}
		if n == 0 {
			break
		}

		if _, err := destination.Write(buf[:n]); err != nil {
			return err
		}
	}
	return nil
}

func WriteListExpiredNotFound(fileName string, source *Utils.SearchResults) {
	f, err := os.OpenFile(fileName, os.O_WRONLY, 0666)
	if err != nil {
		fmt.Println("Error occurred when opening file.", err.Error())
		return
	}

	if err := os.Truncate(fileName, 0); err != nil {
		fmt.Println("Error occurred when truncate file.", err.Error())
		return
	}
	defer f.Close()

	tabs := ""
	_, err = f.WriteString(tabs + "<html>\n")
	tabs = tabs + "\t"

	//_, _ = f.WriteString(tabs + "<h1>Выберите товары для списания</h1>\n")

	tabs = deleteLastChar(tabs)
	_, _ = f.WriteString(tabs + "</html>\n")
}

func WriteListExpired(fileName string, source *Utils.SearchResults) {
	f, err := os.OpenFile(fileName, os.O_WRONLY, 0666)
	if err != nil {
		fmt.Println("Error occurred when opening file.", err.Error())
		return
	}

	if err := os.Truncate(fileName, 0); err != nil {
		fmt.Println("Error occurred when truncate file.", err.Error())
		return
	}
	defer f.Close()

	tabs := ""
	_, err = f.WriteString(tabs + "<html>\n")
	tabs = tabs + "\t"

	_, _ = f.WriteString(tabs + "<ol>\n")
	tabs = tabs + "\t"
	//Table start
	_, _ = f.WriteString(tabs + "<form>\n")
	tabs = tabs + "\t"
	_, _ = f.WriteString(tabs + "<table>\n")
	tabs = tabs + "\t"

	for _, pp := range *source {
		_, _ = f.WriteString(tabs + "<tr>\n")
		_, _ = f.WriteString(tabs + "<td>\n")
		str := fmt.Sprintf("<p style=\"width: 800px; height=100px; background-color: white;\">Товар: %s. Клиент: %s %s. Номер места на складе: %d </p>", pp.Name(), pp.ClName(), pp.Surname(), pp.Place())
		_, _ = f.WriteString(tabs + "<li> " + str + " </li>\n")
		_, _ = f.WriteString(tabs + "</td>\n")
		_, _ = f.WriteString(tabs + "</tr>\n")
	}

	_, _ = f.WriteString(tabs + "<tr>\n")
	_, _ = f.WriteString(tabs + "<td>\n")
	_, _ = f.WriteString(tabs + "<h3>Выведенные товары были списаны.</h3>\n")
	_, _ = f.WriteString(tabs + "</td>\n")
	_, _ = f.WriteString(tabs + "</tr>\n")

	tabs = deleteLastChar(tabs)
	_, _ = f.WriteString(tabs + "</table>\n")
	tabs = deleteLastChar(tabs)
	_, _ = f.WriteString(tabs + "</form>\n")

	//Table end
	tabs = deleteLastChar(tabs)
	_, _ = f.WriteString(tabs + "</ol>\n")

	tabs = deleteLastChar(tabs)
	_, _ = f.WriteString(tabs + "</html>\n")
}

func WriteListChangePlaceError(fileName string) {
	f, err := os.OpenFile(fileName, os.O_WRONLY, 0666)
	if err != nil {
		fmt.Println("Error occurred when opening file.", err.Error())
		return
	}

	if err := os.Truncate(fileName, 0); err != nil {
		fmt.Println("Error occurred when truncate file.", err.Error())
		return
	}
	defer f.Close()

	tabs := ""
	_, err = f.WriteString(tabs + "<html>\n")
	tabs = tabs + "\t"

	_, _ = f.WriteString(tabs + "<h1 style=\"font-size: 40px; font-family: Arial; margin: 10px 0 20px 10px\">Не удалось изменить месторасположение товара. Полный склад.</h1>\n")

	tabs = deleteLastChar(tabs)
	_, _ = f.WriteString(tabs + "</html>\n")
}

func WriteListNotFound(fileName string, source *Utils.SearchResults) {
	f, err := os.OpenFile(fileName, os.O_WRONLY, 0666)
	if err != nil {
		fmt.Println("Error occurred when opening file.", err.Error())
		return
	}

	if err := os.Truncate(fileName, 0); err != nil {
		fmt.Println("Error occurred when truncate file.", err.Error())
		return
	}
	defer f.Close()

	tabs := ""
	_, err = f.WriteString(tabs + "<html>\n")
	tabs = tabs + "\t"

	_, _ = f.WriteString(tabs + "<h1 style=\"font-size: 40px; font-family: Arial; margin: 10px 0 20px 10px\">Товары не найдены.</h1>\n")

	tabs = deleteLastChar(tabs)
	_, _ = f.WriteString(tabs + "</html>\n")
}

func WriteList(fileName string, source *Utils.SearchResults) {
	f, err := os.OpenFile(fileName, os.O_WRONLY, 0777)
	if err != nil {
		fmt.Println("Error occurred when opening file.", err.Error())
		return
	}
	script, err := os.Open("frontend/script.js")
	if err != nil {
		fmt.Println("Error occurred when opening script file.", err.Error())
		return
	}
	if err := os.Truncate(fileName, 0); err != nil {
		fmt.Println("Error occurred when truncate file.", err.Error())
		return
	}
	defer f.Close()

	tabs := ""
	_, err = f.WriteString(tabs + "<html>\n")
	_, err = f.WriteString(tabs + "<script src=\"https://ajax.googleapis.com/ajax/libs/jquery/1.12.4/jquery.min.js\"></script>")
	if err := copy(script, f); err != nil {
		fmt.Println("Error occurred when copy script ", err.Error())
		return
	}
	_, err = f.WriteString("\n")

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
		str := fmt.Sprintf("<p style=\"width: 500px; height=100px; background-color: white;\">Товар: %s. Клиент: %s %s. Номер места на складе: %d</p>", pp.Name(), pp.ClName(), pp.Surname(), pp.Place())
		_, _ = f.WriteString(tabs + "<li> " + str + " </li>\n")
		_, _ = f.WriteString(tabs + "</td>\n")
		addButtons(f, tabs, "Выдать товар", "\"buttonGive\"", pp.IdUniq())
		addButtons(f, tabs, "Изменить местоположение товара", "\"buttonChangePlacement\"", pp.IdUniq())
		_, _ = f.WriteString(tabs + "</tr>\n")

	}

	tabs = deleteLastChar(tabs)
	_, _ = f.WriteString(tabs + "</table>\n")
	tabs = deleteLastChar(tabs)
	_, _ = f.WriteString(tabs + "</form>\n")

}

func addButtons(f *os.File, tabs string, nameButton string, id string, id_uniq int) {
	_, _ = f.WriteString(tabs + "<td>")
	_, _ = f.WriteString("<button class=" + id + " type=\"submit\" value=\"" + strconv.Itoa(id_uniq) + "\">" + nameButton + "</button>")
	_, _ = f.WriteString(tabs + "</td>\n")
}
