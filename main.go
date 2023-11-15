package main

import (
	"fmt"
	"os"

	"github.com/IPGeolocation/ipgeo/apiclient"
	"github.com/IPGeolocation/ipgeo/ascii"
	"github.com/IPGeolocation/ipgeo/cache"
	"github.com/IPGeolocation/ipgeo/config"
	"github.com/IPGeolocation/ipgeo/dbmanager"
)

func main() {
	if len(os.Args) < 2 {
		printHelp()
		os.Exit(1)
	}
	if err := dbmanager.OpenDB(); err != nil {
		fmt.Println("Failed to open the database:", err)
		return
	}
	defer dbmanager.CloseDB()
	command := os.Args[1]
	switch command {
	case "config":
		config.HandleConfig(os.Args[2:])
	case "cache":
		cache.HandleCacheCommand(os.Args[2:])
	case "version":
		fmt.Println("ipgeo cli version 1.0.0")
	case "help":
		printHelp()
	case "ip":
		apiclient.HandleIPGeolocationLookup("", os.Args[2:])
	default:
		apiclient.HandleIPGeolocationLookup(command, os.Args[2:])
	}
}

func printHelp() {
	fmt.Println(ascii.GetAsciiArt() + "\n")
	fmt.Println("\033[1m" + "IP Geolocation CLI" + "\033[0m")
	fmt.Println("Usage: ipgeo <cmd> [<opts>]")
	fmt.Println("Commands:")
	fmt.Println("  <ip/domain> - Look up details for an IP address or domain, e.g., 1.1.1.1 or google.com")
	fmt.Println("  ip          - Look up details for your own IP address.")
	fmt.Println("  config      - Manage the configuration. Use 'login' or 'logout'.")
	fmt.Println("  cache       - Manage the cache. Use 'clear', 'count', 'enable', 'disable', or 'status'.")
	fmt.Println("  version     - Show current version.")
	fmt.Println("  help        - Show this help information.")

	fmt.Println("\nOptions:")
	fmt.Println("  login       - Log in and save API key for the session. Need to use with ipgeo config.")
	fmt.Println("  logout      - Log out and clear the saved API key. Need to use with ipgeo config.")
	fmt.Println("  clear       - Clear the cache. Need to use with ipgeo cache.")
	fmt.Println("  count       - Count the number of entries in the cache. Need to use with ipgeo cache.")
	fmt.Println("  enable      - Enable caching. Need to use with ipgeo cache. Need to use in case user has disabled caching.")
	fmt.Println("  disable     - Disable caching. Need to use with ipgeo cache. By default, caching is enabled.")
	fmt.Println("  status      - Show the status of caching. Need to use with ipgeo cache.")
	fmt.Println("  --fields <fields>, -f <fields>      - Specify fields for output filtering. Multiple fields separated by commas.")
	fmt.Println("  --include <fields>, -i <fields>     - Include additional query parameters in API request.")
	fmt.Println("  --exclude <fields>, -e <fields>     - Exclude query parameters from API request.")
	fmt.Println("  --key-color <color>           - Specify the color for keys in output.")
	fmt.Println("  --value-color <color>         - Specify the color for values in output.")
	fmt.Println("  --compact                  	 - Output json in raw form.")
	fmt.Println("  --json, -j                    - Output in JSON format.")
	fmt.Println("  --yaml, -y                    - Output in YAML format.")
	fmt.Println("  --xml, -x                     - Output in XML format.")
	fmt.Println("  --table, -t                   - Output in CSV format.")
	fmt.Println("  --no-cache                    - Bypass the cache and make a new API request.")
	fmt.Println("  --no-color                    - Disable colorized output.")
}
