package outputs

func OutputAs(format string, data []map[string]string) {
	if format == "table" {
		OutputAsTable(data)
	}
}
