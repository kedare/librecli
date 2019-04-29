package outputs

import (
	"fmt"
	"os"
	"text/tabwriter"
)

func OutputAsList(order []string, data []map[string]string) {
	w := tabwriter.NewWriter(os.Stdout, 1, 1, 1, ' ', tabwriter.TabIndent)

	for _, row := range data {
		for _, col := range order {
			fmt.Fprintf(w, "%s\t%s\r\n", col, row[col])
		}
		fmt.Fprintln(w, "-----------------------")
	}

	w.Flush()

}
