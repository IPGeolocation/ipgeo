package formatter

import (
	"encoding/json"
	"strings"

	"github.com/IPGeolocation/ipgeo/models"
	"github.com/IPGeolocation/ipgeo/resp"
	"github.com/fatih/color"
)

var defaultKeyColor = color.New(color.FgCyan)
var defaultValueColor = color.New(color.FgYellow)

func FormatResponse(response []byte, format string, colorize bool, compact bool, keyColorName string, valueColorName string) (string, error) {
	switch format {
	case "json":
		return formatJSON(response, compact, colorize, keyColorName, valueColorName), nil
	case "yaml":
		return resp.GetResponseInYAML(response), nil
	case "xml":
		return resp.GetResponseInXML(response), nil
	case "table":
		return resp.GetResponseInTable(response), nil
	case "tree":
		keyColor := getColorFromName(keyColorName, defaultKeyColor)
		valueColor := getColorFromName(valueColorName, defaultValueColor)
		return resp.GetResponseInTreeView(response, colorize, keyColor, valueColor), nil
	default:
		var geoResp models.GeolocationResponse
		if err := json.Unmarshal(response, &geoResp); err != nil {
			return "", err
		}
		keyColor := getColorFromName(keyColorName, defaultKeyColor)
		valueColor := getColorFromName(valueColorName, defaultValueColor)
		return formatResponseInBlocks(geoResp, colorize, keyColor, valueColor), nil
	}
}

func formatResponseInBlocks(geoResp models.GeolocationResponse, colorize bool, keyColor *color.Color, valueColor *color.Color) string {
	return resp.GetResponseInBlocks(geoResp, keyColor, valueColor, colorize)
}

func formatJSON(geoResp []byte, compact bool, colorize bool, keyColor string, valueColor string) string {
	return resp.GetResponseInJson(geoResp, compact, colorize, keyColor, valueColor)
}

func getColorFromName(colorName string, defaultColor *color.Color) *color.Color {
	colorName = strings.ToLower(colorName)
	if colorName == "" {
		return defaultColor
	}
	if c, ok := colorMap[colorName]; ok {
		return color.New(c)
	}
	return defaultColor
}

var colorMap = map[string]color.Attribute{
	"red":     color.FgRed,
	"green":   color.FgGreen,
	"blue":    color.FgBlue,
	"yellow":  color.FgYellow,
	"cyan":    color.FgCyan,
	"white":   color.FgWhite,
	"black":   color.FgBlack,
	"magenta": color.FgMagenta,
}
