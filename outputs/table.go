package outputs

import (
	"fmt"
	"github.com/apcera/termtables"
)

func OutputAsTable(order []string, data []map[string]string) {
	table := termtables.CreateTable()

	orderInterface := []interface{}{}
	for _, orderItem := range order {
		orderInterface = append(orderInterface, orderItem)
	}
	table.AddHeaders(orderInterface...)

	for _, row := range data {
		dataInterface := []interface{}{}
		for _, col := range order {
			dataInterface = append(dataInterface, row[col])

		}
		table.AddRow(dataInterface...)
	}
	fmt.Println(table.Render())
}
