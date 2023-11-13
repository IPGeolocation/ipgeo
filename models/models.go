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
	IP                  *string          `json:"ip" yaml:"ip,omitempty"`
	Hostname            *string          `json:"hostname" yaml:"hostname,omitempty"`
	Domain              *string          `json:"domain" yaml:"domain,omitempty"`
	ContinentCode       *string          `json:"continent_code" yaml:"continent_code,omitempty"`
	ContinentName       *string          `json:"continent_name" yaml:"continent_name,omitempty"`
	CountryCode2        *string          `json:"country_code2" yaml:"country_code2,omitempty"`
	CountryCode3        *string          `json:"country_code3" yaml:"country_code3,omitempty"`
	CountryName         *string          `json:"country_name" yaml:"country_name,omitempty"`
	CountryNameOfficial *string          `json:"country_name_official" yaml:"country_name_official,omitempty"`
	CountryCapital      *string          `json:"country_capital" yaml:"country_capital,omitempty"`
	StateProv           *string          `json:"state_prov" yaml:"state_prov,omitempty"`
	StateCode           *string          `json:"state_code" yaml:"state_code,omitempty"`
	District            *string          `json:"district" yaml:"district,omitempty"`
	City                *string          `json:"city" yaml:"city,omitempty"`
	Zipcode             *string          `json:"zipcode" yaml:"zipcode,omitempty"`
	Latitude            *decimal.Decimal `json:"latitude" yaml:"latitude,omitempty"`
	Longitude           *decimal.Decimal `json:"longitude" yaml:"longitude,omitempty"`
	IsEU                *bool            `json:"is_eu" yaml:"is_eu,omitempty"`
	CallingCode         *string          `json:"calling_code" yaml:"calling_code,omitempty"`
	CountryTLD          *string          `json:"country_tld" yaml:"country_tld,omitempty"`
	Languages           *string          `json:"languages" yaml:"languages,omitempty"`
	CountryFlag         *string          `json:"country_flag" yaml:"country_flag,omitempty"`
	GeonameID           *string          `json:"geoname_id" yaml:"geoname_id,omitempty"`
	ISP                 *string          `json:"isp" yaml:"isp,omitempty"`
	ConnectionType      *string          `json:"connection_type" yaml:"connection_type,omitempty"`
	Organization        *string          `json:"organization" yaml:"organization,omitempty"`
	ASN                 *string          `json:"asn" yaml:"asn,omitempty"`
	Currency            *Currency        `json:"currency" yaml:"currency,omitempty"`
	TimeZone            *TimeZone        `json:"time_zone" yaml:"time_zone,omitempty"`
	Security            *Security        `json:"security" yaml:"security,omitempty"`
	UserAgent           *UserAgent       `json:"user_agent" yaml:"user_agent,omitempty"`
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
	ThreatScore     *int    `json:"threat_score"`
	IsTor           *bool   `json:"is_tor"`
	IsProxy         *bool   `json:"is_proxy"`
	ProxyType       *string `json:"proxy_type"`
	IsAnonymous     *bool   `json:"is_anonymous"`
	IsKnownAttacker *bool   `json:"is_known_attacker"`
	IsSpam          *bool   `json:"is_spam"`
	IsBot           *bool   `json:"is_bot"`
	IsCloudProvider *bool   `json:"is_cloud_provider"`
}

type UserAgent struct {
	UserAgentString *string `json:"userAgentString"`
	Name            *string `json:"name"`
	Type            *string `json:"type"`
	Version         *string `json:"version"`
	VersionMajor    *string `json:"versionMajor"`
	Device          *Device `json:"device"`
	Engine          *Engine `json:"engine"`
	OperatingSystem *OS     `json:"operatingSystem"`
}

type Device struct {
	Name  *string `json:"name"`
	Type  *string `json:"type"`
	Brand *string `json:"brand"`
	CPU   *string `json:"cpu"`
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
func (g *GeolocationResponse) MarshalYAML() (interface{}, error) {
	type Alias GeolocationResponse
	alias := (*Alias)(g)

	serialized := make(map[string]interface{})

	v := reflect.ValueOf(alias).Elem()
	t := v.Type()

	for i := 0; i < v.NumField(); i++ {
		field := v.Field(i)
		if field.IsValid() && !field.IsZero() {
			yamlTag := t.Field(i).Tag.Get("yaml")
			if yamlTag == "" || yamlTag == "-" {
				continue
			}
			yamlFieldName := strings.Split(yamlTag, ",")[0]

			serialized[yamlFieldName] = field.Interface()
		}
	}

	return serialized, nil
}
