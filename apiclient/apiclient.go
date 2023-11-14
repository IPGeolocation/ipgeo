package apiclient

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"

	"github.com/IPGeolocation/ipgeo/cache"
	"github.com/IPGeolocation/ipgeo/config"
	"github.com/IPGeolocation/ipgeo/error"
	"github.com/IPGeolocation/ipgeo/formatter"
)

type APIErrorResponse struct {
	Message string `json:"message"`
}

func HandleIPGeolocationLookup(ipOrDomain string, args []string) {
	apiKey, err := config.GetAPIKey()
	if err != nil {
		fmt.Println(err)
		return
	}

	params := url.Values{}
	params.Set("apiKey", apiKey)
	params.Set("ip", ipOrDomain)
	outputFormat, colorEnabled, compact, noCacheFlag, keyColorName, valueColorName := handleArgs(args, &params)
	cacheKey := "ipgeo:" + ipOrDomain + ":" + params.Encode()
	cacheEnabled, _ := cache.GetCacheStatus()

	if !noCacheFlag && cacheEnabled {
		if cachedResponse, found, err := cache.GetFromCache(cacheKey); found && err == nil {
			formattedOutput, err := formatter.FormatResponse(cachedResponse, outputFormat, colorEnabled, compact, keyColorName, valueColorName)
			if err != nil {
				fmt.Println("Error formatting output:", err)
				return
			}
			fmt.Println(formattedOutput)
			return
		}
	}

	fmt.Println("🔎  Looking up IP Geolocation details for", ipOrDomain)

	resp, err := http.Get("https://api.ipgeolocation.io/ipgeo?" + params.Encode())
	if err != nil {
		fmt.Println("Error making API request:", err)
		return
	}
	defer resp.Body.Close()
	if resp.StatusCode >= 400 {
		body, err := io.ReadAll(resp.Body)
		if err != nil {
			fmt.Println("Error reading response body:", err)
			return
		}

		var apiErr APIErrorResponse
		if err := json.Unmarshal(body, &apiErr); err != nil {
			fmt.Println("Error unmarshalling error response:", err)
			return
		}

		customErr := error.HandleErrorResponse(resp.StatusCode, apiErr.Message)
		jsonErr, err := json.Marshal(customErr)
		if err != nil {
			fmt.Println("Error marshaling the custom error:", jsonErr)
			return
		}
		fmt.Println(string(jsonErr))
		return
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error reading response body:", err)
		return
	}

	if resp.StatusCode == http.StatusOK && !noCacheFlag && cacheEnabled {
		err := cache.SaveToCache(cacheKey, body)
		if err != nil {
			fmt.Println("Error saving to cache:", err)
		}
	}

	formattedOutput, err := formatter.FormatResponse(body, outputFormat, colorEnabled, compact, keyColorName, valueColorName)
	if err != nil {
		fmt.Println("Error formatting output:", err)
		return
	}
	fmt.Println(formattedOutput)
}

func handleArgs(args []string, params *url.Values) (string, bool, bool, bool, string, string) {
	var outputFormat string
	var keyColorName string = ""
	var valueColorName string = ""
	colorEnabled := true
	compact := false
	noCacheFlag := false

	for _, arg := range args {
		switch {
		case strings.HasPrefix(arg, "--fields=") || strings.HasPrefix(arg, "-f="):
			fields := getArgValue(arg)
			params.Set("fields", fields)
		case strings.HasPrefix(arg, "--include=") || strings.HasPrefix(arg, "-i="):
			include := getArgValue(arg)
			params.Set("include", include)
		case strings.HasPrefix(arg, "--exclude=") || strings.HasPrefix(arg, "-e="):
			exclude := getArgValue(arg)
			params.Set("excludes", exclude)
		case arg == "--compact":
			compact = true
		case arg == "--json" || arg == "-j":
			outputFormat = "json"
		case arg == "--csv" || arg == "-c":
			outputFormat = "csv"
		case arg == "--yaml" || arg == "-y":
			outputFormat = "yaml"
		case arg == "--xml" || arg == "-x":
			outputFormat = "xml"
		case arg == "--table" || arg == "-t":
			outputFormat = "table"
		case arg == "--tree" || arg == "-h":
			outputFormat = "tree"
		case arg == "--nocolor":
			colorEnabled = false
		case arg == "--no-cache":
			noCacheFlag = true
		case strings.HasPrefix(arg, "--key-color="):
			keyColorName = getArgValue(arg)
		case strings.HasPrefix(arg, "--value-color="):
			valueColorName = getArgValue(arg)
		}
	}
	return outputFormat, colorEnabled, compact, noCacheFlag, keyColorName, valueColorName
}

func getArgValue(arg string) string {
	parts := strings.Split(arg, "=")
	if len(parts) > 1 {
		return parts[1]
	}
	return ""
}
