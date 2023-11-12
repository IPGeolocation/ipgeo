package models

import (
	"bytes"
	"encoding/json"
	"fmt"
	"reflect"
	"strings"

	"github.com/shopspring/decimal"
)


type GeolocationResponse struct {
    IP                   *string        `json:"ip"`
    Hostname             *string        `json:"hostname"`
    Domain               *string        `json:"domain"`
    ContinentCode        *string        `json:"continent_code"`
    ContinentName        *string        `json:"continent_name"`
    CountryCode2         *string        `json:"country_code2"`
    CountryCode3         *string        `json:"country_code3"`
    CountryName          *string        `json:"country_name"`
    CountryNameOfficial  *string        `json:"country_name_official"`
    CountryCapital       *string        `json:"country_capital"`
    StateProv            *string        `json:"state_prov"`
    StateCode            *string        `json:"state_code"`
    District             *string        `json:"district"`
    City                 *string        `json:"city"`
    Zipcode              *string        `json:"zipcode"`
    Latitude             *decimal.Decimal        `json:"latitude"`
    Longitude            *decimal.Decimal        `json:"longitude"`
    IsEU                 *bool          `json:"is_eu"`
    CallingCode          *string        `json:"calling_code"`
    CountryTLD           *string        `json:"country_tld"`
    Languages            *string        `json:"languages"`
    CountryFlag          *string        `json:"country_flag"`
    GeonameID            *string        `json:"geoname_id"`
    ISP                  *string        `json:"isp"`
    ConnectionType       *string        `json:"connection_type"`
    Organization         *string        `json:"organization"`
    ASN                  *string        `json:"asn"`
    Currency             *Currency      `json:"currency"`
    TimeZone             *TimeZone      `json:"time_zone"`
    Security             *Security      `json:"security"`
    UserAgent            *UserAgent     `json:"user_agent"`
}

type Currency struct {
    Code   *string `json:"code"`
    Name   *string `json:"name"`
    Symbol *string `json:"symbol"`
}

type TimeZone struct {
    Name            *string  `json:"name"`
    Offset          *int     `json:"offset"`
    OffsetWithDST   *int     `json:"offset_with_dst"`
    CurrentTime     *string  `json:"current_time"`
    CurrentTimeUnix *float64 `json:"current_time_unix"`
    IsDST           *bool    `json:"is_dst"`
    DSTSavings      *int     `json:"dst_savings"`
}

type Security struct {
    ThreatScore      *int    `json:"threat_score"`
    IsTor            *bool   `json:"is_tor"`
    IsProxy          *bool   `json:"is_proxy"`
    ProxyType        *string `json:"proxy_type"`
    IsAnonymous      *bool   `json:"is_anonymous"`
    IsKnownAttacker  *bool   `json:"is_known_attacker"`
    IsSpam           *bool   `json:"is_spam"`
    IsBot            *bool   `json:"is_bot"`
    IsCloudProvider  *bool   `json:"is_cloud_provider"`
}

type UserAgent struct {
    UserAgentString *string      `json:"userAgentString"`
    Name            *string      `json:"name"`
    Type            *string      `json:"type"`
    Version         *string      `json:"version"`
    VersionMajor    *string      `json:"versionMajor"`
    Device          *Device      `json:"device"`
    Engine          *Engine      `json:"engine"`
    OperatingSystem *OS          `json:"operatingSystem"`
}

type Device struct {
    Name *string `json:"name"`
    Type *string `json:"type"`
    Brand *string `json:"brand"`
    CPU  *string `json:"cpu"`
}

type Engine struct {
    Name         *string `json:"name"`
    Type         *string `json:"type"`
    Version      *string `json:"version"`
    VersionMajor *string `json:"versionMajor"`
}

type OS struct {
    Name         *string `json:"name"`
    Type         *string `json:"type"`
    Version      *string `json:"version"`
    VersionMajor *string `json:"versionMajor"`
}
func (resp *GeolocationResponse) MarshalJSON() ([]byte, error) {
	var buf bytes.Buffer
	buf.WriteString("{")

	v := reflect.ValueOf(resp).Elem()
	t := v.Type()
	first := true
	for i := 0; i < v.NumField(); i++ {
		field := v.Field(i)
		if field.IsValid() && !field.IsZero() {
			jsonTag := t.Field(i).Tag.Get("json")
			if jsonTag == "" || jsonTag == "-" {
				continue
			}
			jsonFieldName := strings.Split(jsonTag, ",")[0]

			if !first {
				buf.WriteString(", ")
			}
			first = false

			buf.WriteString(fmt.Sprintf("\"%s\": ", jsonFieldName))

			
			fieldBytes, err := json.Marshal(field.Interface())
			if err != nil {
				return nil, err
			}
			buf.Write(fieldBytes)
		}
	}

	buf.WriteString("}")
	return buf.Bytes(), nil
}
