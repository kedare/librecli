package outputs

import (
	"fmt"
	"github.com/apcera/termtables"
)

func OutputAsTable(data []map[string]string) {
	headers := []interface{}{}
	for headerFromData := range data[0] {
		headers = append(headers, headerFromData)
	}

	table := termtables.CreateTable()
	table.AddHeaders(headers...)

	for _, row := range data {
		dataInterface := []interface{}{}
		for _, col := range headers {
			dataInterface = append(dataInterface, row[col.(string)])

		}
		table.AddRow(dataInterface...)
	}
	fmt.Println(table.Render())
}
