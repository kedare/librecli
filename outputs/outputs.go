package outputs

func OutputAs(format string, order []string, data []map[string]string) {
	if format == "table" {
		OutputAsTable(order, data)
	}
}
