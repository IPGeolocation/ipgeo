package resp

import (
	"fmt"
	"strings"

	"github.com/buger/jsonparser"
	"github.com/olekukonko/tablewriter"
)

func GetResponseInTable(geoResp []byte) string {
	if tableStr, err := displayAsTable(geoResp, "GeoLocation", 0); err != nil {
		fmt.Println("Error displaying response as table:", err)
	} else {
		return tableStr
	}
	return ""
}
func displayAsTable(jsonData []byte, header string, indentLevel int) (string, error) {
	var table *tablewriter.Table
	tableString := &strings.Builder{}
	if indentLevel == 0 {
		tableString.WriteString(header + "\n")
	} else {
		tableString.WriteString(strings.Repeat(" ", indentLevel*4) + "↳ " + header + "\n")
	}

	startNewTable := func() {
		if table != nil {
			table.Render()
		}
		table = tablewriter.NewWriter(tableString)
		table.SetHeader([]string{"Field", "Value"})
		table.SetAlignment(tablewriter.ALIGN_LEFT)
		table.SetTablePadding("\t")
	}

	startNewTable()

	var processObject func([]byte, string, int) error
	processObject = func(data []byte, currentHeader string, indentLevel int) error {
		return jsonparser.ObjectEach(data, func(key []byte, value []byte, dataType jsonparser.ValueType, offset int) error {
			keyStr := string(key)
			if dataType == jsonparser.Object && len(value) > 2 {
				if table != nil && table.NumLines() > 0 {
					table.Render()
					table = nil
				}
				newIndentLevel := indentLevel + 1
				return processObject(value, strings.Title(keyStr), newIndentLevel)
			}
			if table == nil {
				if indentLevel > 0 {
					tableString.WriteString(strings.Repeat(" ", indentLevel*4) + "↳ " + currentHeader + "\n")
				} else {
					tableString.WriteString(currentHeader + "\n")
				}
				startNewTable()
			}
			table.Append([]string{keyStr, string(value)})

			return nil
		})
	}

	err := processObject(jsonData, header, 0)
	if err != nil {
		return "", err
	}

	if table != nil && table.NumLines() > 0 {
		table.Render()
	}

	return tableString.String(), nil
}
