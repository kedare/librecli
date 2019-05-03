package outputs

// OutputAs is the abstraction layer that will route to the proper display system
func OutputAs(format string, order []string, data []map[string]string) {
	if format == "table" {
		OutputAsTable(order, data)
	} else if format == "list" {
		OutputAsList(order, data)
	}
}
