package resp

import (
	"strings"

	"github.com/buger/jsonparser"
	"github.com/fatih/color"
)

func GetResponseInTreeView(geoResp []byte, colorize bool, keyColor *color.Color, valueColor *color.Color) string {
	treeString := &strings.Builder{}
	treeString.WriteString(color.New(color.FgWhite, color.Bold).SprintFunc()("Geolocation\n"))
	var processObject func([]byte, string, int) error
	processObject = func(data []byte, currentHeader string, indentLevel int) error {
		return jsonparser.ObjectEach(data, func(key []byte, value []byte, dataType jsonparser.ValueType, offset int) error {
			keyStr := string(key)
			whiteKeys := []string{"currency", "time_zone", "security", "user_agent", "device", "engine", "operatingSystem", "tags"}
			isWhite := false
			for _, whiteKey := range whiteKeys {
				if keyStr == whiteKey {
					isWhite = true
					break
				}
			}
			switch keyStr {
			case "currency":
				keyStr = "Currency"
			case "time_zone":
				keyStr = "TimeZone"
			case "security":
				keyStr = "Security"
			case "user_agent":
				keyStr = "UserAgent"
			case "device":
				keyStr = "Device"
			case "engine":
				keyStr = "Engine"
			case "operatingSystem":
				keyStr = "OperatingSystem"
			case "tags":
				keyStr = "Tags"
			}
			if dataType == jsonparser.Object {
				if isWhite {
					treeString.WriteString(strings.Repeat(" ", indentLevel*4) + "├── " + color.New(color.FgWhite, color.Bold).SprintFunc()(keyStr) + "\n")
				} else if colorize {
					treeString.WriteString(strings.Repeat(" ", indentLevel*4) + "├── " + keyColor.SprintFunc()(keyStr) + "\n")
				} else {
					treeString.WriteString(strings.Repeat(" ", indentLevel*4) + "├── " + color.WhiteString(keyStr) + "\n")
				}
				return processObject(value, keyStr, indentLevel+1)
			} else {
				if colorize {
					treeString.WriteString(strings.Repeat(" ", indentLevel*4) + "│   ├── " + keyColor.SprintFunc()(keyStr) + ": " + valueColor.SprintFunc()(string(value)) + "\n")
				} else {
					treeString.WriteString(strings.Repeat(" ", indentLevel*4) + "│   ├── " + color.WhiteString(keyStr) + ": " + color.WhiteString(string(value)) + "\n")
				}
				return nil
			}
		})
	}
	processObject(geoResp, "", 1)
	return treeString.String()
}
