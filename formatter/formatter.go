package formatter

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"strings"

	"github.com/IPGeolocation/ipgeo/models"
	"github.com/IPGeolocation/ipgeo/resp"
	"github.com/fatih/color"
	"github.com/ghodss/yaml"
	"github.com/gocarina/gocsv"
)

var defaultKeyColor = color.New(color.FgCyan)
var defaultValueColor = color.New(color.FgYellow)

func FormatResponse(response []byte, format string, colorize bool, compact bool, keyColorName string, valueColorName string) (string, error) {
	var geoResp models.GeolocationResponse
	if err := json.Unmarshal(response, &geoResp); err != nil {
		return "", err
	}

	switch format {
	case "json":
		return formatJSON(geoResp, compact), nil
	case "csv":
		return formatCSV(geoResp)
	case "yaml":
		return formatYAML(geoResp)
	case "xml":
		return formatXML(geoResp)
	default:
		return formatResponseInBlocks(response, colorize, keyColorName, valueColorName), nil
	}
}

func formatResponseInBlocks(response []byte, colorize bool, keyColorName string, valueColorName string) string {
	 var geoResp models.GeolocationResponse
    if err := json.Unmarshal(response, &geoResp); err != nil {
			fmt.Println("Error unmarshalling response:", err)
			return ""
    }
    keyColor := getColorFromName(keyColorName, defaultKeyColor)
    valueColor := getColorFromName(valueColorName, defaultValueColor)
		return resp.GetResponse(geoResp, keyColor, valueColor)
}

func formatJSON(resp models.GeolocationResponse, compact bool) string {
	var jsonResp []byte
	var err error

	if compact {
		jsonResp, err = json.Marshal(&resp)
	} else {
		jsonResp, err = json.MarshalIndent(&resp, " ", "\t")
		if err == nil {
			    lines := strings.Split(string(jsonResp), "\n")
    for i, line := range lines {
        if strings.Contains(line, ":") {
            // Split the line into key and value
            parts := strings.SplitN(line, ":", 2)

            // Color the key (before the ':') in blue
            key := color.New(color.FgMagenta).Sprint(strings.TrimSpace(parts[0]))

            // Apply color to the colon and space after key
            colonAndSpace := color.New(color.FgHiCyan, color.Bold).Sprint(":")

						value := color.New(color.FgCyan).Sprint(strings.TrimSpace(parts[1]))

            // Combine the colored elements with the value part
            lines[i] = "\t" + " " + key + colonAndSpace + " " + value
        }
    }

    // Join the lines back together
    return strings.Join(lines, "\n")
		}
	}

	if err != nil {
		return "Error formatting JSON: " + err.Error()
	}
	return string(jsonResp)
}

func formatCSV(resp models.GeolocationResponse) (string, error) {
	respSlice := []*models.GeolocationResponse{&resp}
	csvStr, err := gocsv.MarshalString(respSlice)
	if err != nil {
		return "", err
	}
	return csvStr, nil
}

func formatYAML(resp models.GeolocationResponse) (string, error) {
	yamlStr, err := yaml.Marshal(&resp)
	if err != nil {
		return "", err
	}
	return string(yamlStr), nil
}

func formatXML(resp models.GeolocationResponse) (string, error) {
	xmlStr, err := xml.MarshalIndent(&resp, "", "  ")
	if err != nil {
		return "", err
	}
	return string(xmlStr), nil
}

func getColorFromName(colorName string, defaultColor *color.Color) *color.Color {
    if colorName == "" {
        return defaultColor
    }
    if c, ok := colorMap[colorName]; ok {
        return color.New(c)
    }
    return defaultColor
}

var colorMap = map[string]color.Attribute{
    "red":    color.FgRed,
    "green":  color.FgGreen,
    "blue":   color.FgBlue,
    "yellow": color.FgYellow,
    "cyan":   color.FgCyan,
		"white":  color.FgWhite,
		"black":  color.FgBlack,
		"magenta": color.FgMagenta,
		"hi-red": color.FgHiRed,
		"hi-green": color.FgHiGreen,
		"hi-blue": color.FgHiBlue,
		"hi-yellow": color.FgHiYellow,
		"hi-cyan": color.FgHiCyan,
		"hi-white": color.FgHiWhite,
		"hi-black": color.FgHiBlack,
		"hi-magenta": color.FgHiMagenta,
}