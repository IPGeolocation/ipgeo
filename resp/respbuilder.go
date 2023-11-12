package resp

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/IPGeolocation/ipgeo/models"
	"github.com/fatih/color"
)


func GetResponse(geoResp models.GeolocationResponse, keyColor *color.Color, valueColor *color.Color) string {
    white := color.New(color.FgWhite)
    bold := color.New(color.Bold)
    outputBuilder := &strings.Builder{}

    const keyWidth = 20

		outputBuilder.WriteString("\n")
    outputBuilder.WriteString(bold.Sprintf("GeoLocation: "))
    outputBuilder.WriteString("\n\n")

		appendGeoLocationSection(geoResp, outputBuilder, keyWidth, keyColor, valueColor, white, bold)

		appendCurrencySection(geoResp, outputBuilder, keyWidth, keyColor, valueColor, white, bold)

		appendTimeZoneSection(geoResp, outputBuilder, keyWidth, keyColor, valueColor, white, bold)

		appendSecuritySection(geoResp, outputBuilder, keyWidth, keyColor, valueColor, white, bold)

		appendUserAgentSection(geoResp, outputBuilder, keyWidth, keyColor, valueColor, white, bold)

    return outputBuilder.String()
}

func writeEntry(builder *strings.Builder, key string, value *string, width int, keyColor, valueColor, white *color.Color) {
    if value != nil {
        paddedKey := fmt.Sprintf("%-*s", width, key) 
        builder.WriteString(white.Sprintf("- "))
        builder.WriteString(keyColor.Sprintf(paddedKey))
        builder.WriteString(white.Sprintf(" "))
        builder.WriteString(valueColor.Sprintf("%s\n", *value))
    }
}
func appendGeoLocationSection(geoResp models.GeolocationResponse, outputBuilder *strings.Builder, keyWidth int, keyColor *color.Color, valueColor *color.Color, white *color.Color, bold *color.Color) {

	if(geoResp.IP != nil) {
			ip := geoResp.IP
			writeEntry(outputBuilder, "IP", ip, keyWidth, keyColor, valueColor, white)
		}
		if(geoResp.Hostname != nil) {
			hostname := geoResp.Hostname
			writeEntry(outputBuilder, "Hostname", hostname, keyWidth, keyColor, valueColor, white)
		}
		if(geoResp.Domain != nil) {
			domain := geoResp.Domain
			writeEntry(outputBuilder, "Domain", domain, keyWidth, keyColor, valueColor, white)
		}
		if(geoResp.ContinentCode != nil) {
			continentCode := geoResp.ContinentCode
			writeEntry(outputBuilder, "ContinentCode", continentCode, keyWidth, keyColor, valueColor, white)
		}
		if(geoResp.ContinentName != nil) {
			continentName := geoResp.ContinentName
			writeEntry(outputBuilder, "ContinentName", continentName, keyWidth, keyColor, valueColor, white)
		}
		if(geoResp.CountryCode2 != nil) {
			countryCode2 := geoResp.CountryCode2
			writeEntry(outputBuilder, "CountryCode2", countryCode2, keyWidth, keyColor, valueColor, white)
		}
		if(geoResp.CountryCode3 != nil) {
			countryCode3 := geoResp.CountryCode3
			writeEntry(outputBuilder, "CountryCode3", countryCode3, keyWidth, keyColor, valueColor, white)
		}
		if(geoResp.CountryName != nil) {
			countryName := geoResp.CountryName
			writeEntry(outputBuilder, "CountryName", countryName, keyWidth, keyColor, valueColor, white)
		}
		if(geoResp.CountryNameOfficial != nil) {
			countryNameOfficial := geoResp.CountryNameOfficial
			writeEntry(outputBuilder, "CountryNameOfficial", countryNameOfficial, keyWidth, keyColor, valueColor, white)
		}
		if(geoResp.CountryCapital != nil) {
			countryCapital := geoResp.CountryCapital
			writeEntry(outputBuilder, "CountryCapital", countryCapital, keyWidth, keyColor, valueColor, white)
		}
		if(geoResp.StateProv != nil) {
			stateProv := geoResp.StateProv
			writeEntry(outputBuilder, "StateProv", stateProv, keyWidth, keyColor, valueColor, white)
		}
		if(geoResp.StateCode != nil) {
			stateCode := geoResp.StateCode
			writeEntry(outputBuilder, "StateCode", stateCode, keyWidth, keyColor, valueColor, white)
		}
		if(geoResp.District != nil) {
			district := geoResp.District
			writeEntry(outputBuilder, "District", district, keyWidth, keyColor, valueColor, white)
		}
		if(geoResp.City != nil) {
			city := geoResp.City
			writeEntry(outputBuilder, "City", city, keyWidth, keyColor, valueColor, white)
		}
		if(geoResp.Zipcode != nil) {
			zipcode := geoResp.Zipcode
			writeEntry(outputBuilder, "Zipcode", zipcode, keyWidth, keyColor, valueColor, white)
		}
		if(geoResp.Latitude != nil) {
			latitude := geoResp.Latitude.String()
			writeEntry(outputBuilder, "Latitude", &latitude, keyWidth, keyColor, valueColor, white)
		}
		if(geoResp.Longitude != nil) {
			longitude := geoResp.Longitude.String()
			writeEntry(outputBuilder, "Longitude", &longitude, keyWidth, keyColor, valueColor, white)
		}
		if(geoResp.IsEU != nil) {
			isEU := strconv.FormatBool(*geoResp.IsEU)
			writeEntry(outputBuilder, "IsEU", &isEU, keyWidth, keyColor, valueColor, white)
		}
		if(geoResp.CallingCode != nil) {
			callingCode := geoResp.CallingCode
			writeEntry(outputBuilder, "CallingCode", callingCode, keyWidth, keyColor, valueColor, white)
		}
		if(geoResp.CountryTLD != nil) {
			countryTLD := geoResp.CountryTLD
			writeEntry(outputBuilder, "CountryTLD", countryTLD, keyWidth, keyColor, valueColor, white)
		}
		if(geoResp.Languages != nil) {
			languages := geoResp.Languages
			writeEntry(outputBuilder, "Languages", languages, keyWidth, keyColor, valueColor, white)
		}
		if(geoResp.CountryFlag != nil) {
			countryFlag := geoResp.CountryFlag
			writeEntry(outputBuilder, "CountryFlag", countryFlag, keyWidth, keyColor, valueColor, white)
		}
		if(geoResp.GeonameID != nil) {
			geonameID := geoResp.GeonameID
			writeEntry(outputBuilder, "GeonameID", geonameID, keyWidth, keyColor, valueColor, white)
		}
		if(geoResp.ISP != nil) {
			iSP := geoResp.ISP
			writeEntry(outputBuilder, "ISP", iSP, keyWidth, keyColor, valueColor, white)
		}
		if(geoResp.ConnectionType != nil) {
			connectionType := geoResp.ConnectionType
			writeEntry(outputBuilder, "ConnectionType", connectionType, keyWidth, keyColor, valueColor, white)
		}
		if(geoResp.Organization != nil) {
			organization := geoResp.Organization
			writeEntry(outputBuilder, "Organization", organization, keyWidth, keyColor, valueColor, white)
		}
		if(geoResp.ASN != nil) {
			aSN := geoResp.ASN
			writeEntry(outputBuilder, "ASN", aSN, keyWidth, keyColor, valueColor, white)
		}

}

func appendCurrencySection(geoResp models.GeolocationResponse, outputBuilder *strings.Builder, keyWidth int, keyColor *color.Color, valueColor *color.Color, white *color.Color, bold *color.Color){
	if(geoResp.Currency != nil) {
			outputBuilder.WriteString("\n")
    outputBuilder.WriteString(bold.Sprintf("Currency: "))
    outputBuilder.WriteString("\n\n")
			if(geoResp.Currency.Name != nil) {
			currencyName := geoResp.Currency.Name
			writeEntry(outputBuilder, "Name", currencyName, keyWidth, keyColor, valueColor, white)
		}
		if(geoResp.Currency.Code != nil) {
			currencyCode := geoResp.Currency.Code
			writeEntry(outputBuilder, "Code", currencyCode, keyWidth, keyColor, valueColor, white)
		}
		if(geoResp.Currency.Symbol != nil) {
			currencySymbol := geoResp.Currency.Symbol
			writeEntry(outputBuilder, "Symbol", currencySymbol, keyWidth, keyColor, valueColor, white)
		}
		}
}

func appendTimeZoneSection(geoResp models.GeolocationResponse, outputBuilder *strings.Builder, keyWidth int, keyColor *color.Color, valueColor *color.Color, white *color.Color, bold *color.Color){
		if(geoResp.TimeZone != nil){
			outputBuilder.WriteString("\n")
    outputBuilder.WriteString(bold.Sprintf("TimeZone: "))
    outputBuilder.WriteString("\n\n")
			if(geoResp.TimeZone.Name != nil) {
			timeZoneName := geoResp.TimeZone.Name
			writeEntry(outputBuilder, "Name", timeZoneName, keyWidth, keyColor, valueColor, white)
		}
		if(geoResp.TimeZone.Offset != nil) {
			timeZoneOffset := strconv.Itoa(*geoResp.TimeZone.Offset)
			writeEntry(outputBuilder, "Offset", &timeZoneOffset, keyWidth, keyColor, valueColor, white)
		}
		if(geoResp.TimeZone.OffsetWithDST != nil) {
			timeZoneOffsetWithDST := strconv.Itoa(*geoResp.TimeZone.OffsetWithDST)
			writeEntry(outputBuilder, "OffsetWithDST", &timeZoneOffsetWithDST, keyWidth, keyColor, valueColor, white)
		}
		if(geoResp.TimeZone.CurrentTime != nil) {
			timeZoneCurrentTime := geoResp.TimeZone.CurrentTime
			writeEntry(outputBuilder, "CurrentTime", timeZoneCurrentTime, keyWidth, keyColor, valueColor, white)
		}
		if(geoResp.TimeZone.CurrentTimeUnix != nil) {
			timeZoneCurrentTimeUnix := strconv.FormatFloat(*geoResp.TimeZone.CurrentTimeUnix, 'f', 6, 64)
			writeEntry(outputBuilder, "CurrentTimeUnix", &timeZoneCurrentTimeUnix, keyWidth, keyColor, valueColor, white)
		}
		if(geoResp.TimeZone.IsDST != nil) {
			timeZoneIsDST := strconv.FormatBool(*geoResp.TimeZone.IsDST)
			writeEntry(outputBuilder, "IsDST", &timeZoneIsDST, keyWidth, keyColor, valueColor, white)
		}
		if(geoResp.TimeZone.DSTSavings != nil) {
			timeZoneDSTSavings := strconv.Itoa(*geoResp.TimeZone.DSTSavings)
			writeEntry(outputBuilder, "DSTSavings", &timeZoneDSTSavings, keyWidth, keyColor, valueColor, white)
		}
		}
}
func appendSecuritySection(geoResp models.GeolocationResponse, outputBuilder *strings.Builder, keyWidth int, keyColor *color.Color, valueColor *color.Color, white *color.Color, bold *color.Color){
		if(geoResp.Security != nil) {
			outputBuilder.WriteString("\n")
    outputBuilder.WriteString(bold.Sprintf("Security: "))
    outputBuilder.WriteString("\n\n")
			if(geoResp.Security.ThreatScore != nil) {
			securityThreatScore := strconv.Itoa(*geoResp.Security.ThreatScore)
			writeEntry(outputBuilder, "ThreatScore", &securityThreatScore, keyWidth, keyColor, valueColor, white)
		}
		if(geoResp.Security.IsTor != nil) {
			securityIsTor := strconv.FormatBool(*geoResp.Security.IsTor)
			writeEntry(outputBuilder, "IsTor", &securityIsTor, keyWidth, keyColor, valueColor, white)
		}
		if(geoResp.Security.IsProxy != nil) {
			securityIsProxy := strconv.FormatBool(*geoResp.Security.IsProxy)
			writeEntry(outputBuilder, "IsProxy", &securityIsProxy, keyWidth, keyColor, valueColor, white)
		}
		if(geoResp.Security.ProxyType != nil) {
			securityProxyType := geoResp.Security.ProxyType
			writeEntry(outputBuilder, "ProxyType", securityProxyType, keyWidth, keyColor, valueColor, white)
		}
		if(geoResp.Security.IsAnonymous != nil) {
			securityIsAnonymous := strconv.FormatBool(*geoResp.Security.IsAnonymous)
			writeEntry(outputBuilder, "IsAnonymous", &securityIsAnonymous, keyWidth, keyColor, valueColor, white)
		}
		if(geoResp.Security.IsKnownAttacker != nil) {
			securityIsKnownAttacker := strconv.FormatBool(*geoResp.Security.IsKnownAttacker)
			writeEntry(outputBuilder, "IsKnownAttacker", &securityIsKnownAttacker, keyWidth, keyColor, valueColor, white)
		}
		if(geoResp.Security.IsSpam != nil) {
			securityIsSpam := strconv.FormatBool(*geoResp.Security.IsSpam)
			writeEntry(outputBuilder, "IsSpam", &securityIsSpam, keyWidth, keyColor, valueColor, white)
		}
		if(geoResp.Security.IsBot != nil) {
			securityIsBot := strconv.FormatBool(*geoResp.Security.IsBot)
			writeEntry(outputBuilder, "IsBot", &securityIsBot, keyWidth, keyColor, valueColor, white)
		}
		if(geoResp.Security.IsCloudProvider != nil) {
			securityIsCloudProvider := strconv.FormatBool(*geoResp.Security.IsCloudProvider)
			writeEntry(outputBuilder, "IsCloudProvider", &securityIsCloudProvider, keyWidth, keyColor, valueColor, white)
		}
		}
}

func appendUserAgentSection(geoResp models.GeolocationResponse, outputBuilder *strings.Builder, keyWidth int, keyColor *color.Color, valueColor *color.Color, white *color.Color, bold *color.Color){
		if(geoResp.UserAgent != nil) {
			outputBuilder.WriteString("\n")
    outputBuilder.WriteString(bold.Sprintf("UserAgent: "))
    outputBuilder.WriteString("\n\n")
			if(geoResp.UserAgent.UserAgentString != nil) {
			userAgentUserAgentString := geoResp.UserAgent.UserAgentString
			writeEntry(outputBuilder, "UserAgentString", userAgentUserAgentString, keyWidth, keyColor, valueColor, white)
		}
		if(geoResp.UserAgent.Name != nil) {
			userAgentName := geoResp.UserAgent.Name
			writeEntry(outputBuilder, "Name", userAgentName, keyWidth, keyColor, valueColor, white)
		}
		if(geoResp.UserAgent.Type != nil) {
			userAgentType := geoResp.UserAgent.Type
			writeEntry(outputBuilder, "Type", userAgentType, keyWidth, keyColor, valueColor, white)
		}
		if(geoResp.UserAgent.Version != nil) {
			userAgentVersion := geoResp.UserAgent.Version
			writeEntry(outputBuilder, "Version", userAgentVersion, keyWidth, keyColor, valueColor, white)
		}
		if(geoResp.UserAgent.VersionMajor != nil) {
			userAgentVersionMajor := geoResp.UserAgent.VersionMajor
			writeEntry(outputBuilder, "VersionMajor", userAgentVersionMajor, keyWidth, keyColor, valueColor, white)
		}
		if(geoResp.UserAgent.Device != nil) {
			outputBuilder.WriteString("\n")
			outputBuilder.WriteString(bold.Sprintf("UserAgent Device: "))
			outputBuilder.WriteString("\n\n")
			if(geoResp.UserAgent.Device.Name != nil) {
			userAgentDeviceName := geoResp.UserAgent.Device.Name
			writeEntry(outputBuilder, "Name", userAgentDeviceName, keyWidth, keyColor, valueColor, white)
		}
		if(geoResp.UserAgent.Device.Type != nil) {
			userAgentDeviceType := geoResp.UserAgent.Device.Type
			writeEntry(outputBuilder, "Type", userAgentDeviceType, keyWidth, keyColor, valueColor, white)
		}
		if(geoResp.UserAgent.Device.Brand != nil) {
			userAgentDeviceBrand := geoResp.UserAgent.Device.Brand
			writeEntry(outputBuilder, "Brand", userAgentDeviceBrand, keyWidth, keyColor, valueColor, white)
		}
		if(geoResp.UserAgent.Device.CPU != nil) {
			userAgentDeviceCPU := geoResp.UserAgent.Device.CPU
			writeEntry(outputBuilder, "CPU", userAgentDeviceCPU, keyWidth, keyColor, valueColor, white)
		}
		}
		if(geoResp.UserAgent.Engine != nil) {
			outputBuilder.WriteString("\n")
			outputBuilder.WriteString(bold.Sprintf("UserAgent Engine: "))
			outputBuilder.WriteString("\n\n")
			if(geoResp.UserAgent.Engine.Name != nil) {
			userAgentEngineName := geoResp.UserAgent.Engine.Name
			writeEntry(outputBuilder, "Name", userAgentEngineName, keyWidth, keyColor, valueColor, white)
		}
		if(geoResp.UserAgent.Engine.Type != nil) {
			userAgentEngineType := geoResp.UserAgent.Engine.Type
			writeEntry(outputBuilder, "Type", userAgentEngineType, keyWidth, keyColor, valueColor, white)
		}
		if(geoResp.UserAgent.Engine.Version != nil) {
			userAgentEngineVersion := geoResp.UserAgent.Engine.Version
			writeEntry(outputBuilder, "Version", userAgentEngineVersion, keyWidth, keyColor, valueColor, white)
		}
		if(geoResp.UserAgent.Engine.VersionMajor != nil) {
			userAgentEngineVersionMajor := geoResp.UserAgent.Engine.VersionMajor
			writeEntry(outputBuilder, "VersionMajor", userAgentEngineVersionMajor, keyWidth, keyColor, valueColor, white)
		}
		}
		if(geoResp.UserAgent.OperatingSystem != nil) {
			outputBuilder.WriteString("\n")
			outputBuilder.WriteString(bold.Sprintf("UserAgent OperatingSystem: "))
			outputBuilder.WriteString("\n\n")
			if(geoResp.UserAgent.OperatingSystem.Name != nil) {
			userAgentOperatingSystemName := geoResp.UserAgent.OperatingSystem.Name
			writeEntry(outputBuilder, "Name", userAgentOperatingSystemName, keyWidth, keyColor, valueColor, white)
		}
		if(geoResp.UserAgent.OperatingSystem.Type != nil) {
			userAgentOperatingSystemType := geoResp.UserAgent.OperatingSystem.Type
			writeEntry(outputBuilder, "Type", userAgentOperatingSystemType, keyWidth, keyColor, valueColor, white)
		}
		if(geoResp.UserAgent.OperatingSystem.Version != nil) {
			userAgentOperatingSystemVersion := geoResp.UserAgent.OperatingSystem.Version
			writeEntry(outputBuilder, "Version", userAgentOperatingSystemVersion, keyWidth, keyColor, valueColor, white)
		}
		if(geoResp.UserAgent.OperatingSystem.VersionMajor != nil) {
			userAgentOperatingSystemVersionMajor := geoResp.UserAgent.OperatingSystem.VersionMajor
			writeEntry(outputBuilder, "VersionMajor", userAgentOperatingSystemVersionMajor, keyWidth, keyColor, valueColor, white)
		}
		}
		}
}
