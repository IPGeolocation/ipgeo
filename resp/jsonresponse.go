package resp

import (
	"strings"

	"github.com/tidwall/pretty"
)

type Style struct {
	Key, String, Number [2]string
	True, False, Null   [2]string
	Escape              [2]string
	Brackets            [2]string
	Append              func(dst []byte, c byte) []byte
}

func GetResponseInJson(geoResp []byte, compact bool, colorize bool, keyColor string, valueColor string) string {
	jsonResp := pretty.Pretty(geoResp)
	if !compact && colorize && (keyColor != "" || valueColor != "") {
		colorOfKey := colorToANSICode(keyColor)
		colorOfValue := colorToANSICode(valueColor)
		if colorOfKey != "" && colorOfValue != "" {
			result := pretty.Color(jsonResp, CustomizedStyle(colorOfKey, colorOfValue))
			return string(result)
		} else {
			result := pretty.Color(jsonResp, nil)
			return string(result)
		}
	} else if !compact && colorize && keyColor == "" && valueColor == "" {
		return string(pretty.Color(jsonResp, nil))
	} else if !compact && !colorize {
		return string(pretty.Pretty(jsonResp))
	} else if compact {
		return string(pretty.UglyInPlace(jsonResp))
	} else {
		return string(pretty.Pretty(jsonResp))
	}
}

func CustomizedStyle(keyColor, valueColor string) *pretty.Style {
	var CustomStyle *pretty.Style
	CustomStyle = &pretty.Style{
		Key:      [2]string{keyColor, "\x1B[0m"},
		String:   [2]string{valueColor, "\x1B[0m"},
		Number:   [2]string{"\x1B[33m", "\x1B[0m"},
		True:     [2]string{"\x1B[36m", "\x1B[0m"},
		False:    [2]string{"\x1B[36m", "\x1B[0m"},
		Null:     [2]string{"\x1B[2m", "\x1B[0m"},
		Escape:   [2]string{"\x1B[35m", "\x1B[0m"},
		Brackets: [2]string{"\x1B[1m", "\x1B[0m"},
		Append:   nil,
	}
	return CustomStyle
}
func colorToANSICode(c string) string {
	c = strings.ToLower(c)
	code, ok := colorToANSICodeMap[c]
	if !ok {
		return ""
	}
	return code
}

// TODO: Add more colors and formatting later

var colorToANSICodeMap = map[string]string{
	"black":   "\x1B[30m",
	"red":     "\x1B[31;49;1m",
	"green":   "\x1B[32m",
	"yellow":  "\x1B[33m",
	"blue":    "\x1B[34m",
	"magenta": "\x1B[35m",
	"cyan":    "\x1B[36m",
	"white":   "\x1B[37m",
}
