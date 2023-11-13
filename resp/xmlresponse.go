package resp

import (
	"encoding/json"
	"encoding/xml"
	"fmt"

	"github.com/IPGeolocation/ipgeo/models"
)

func GetResponseInXML(resp []byte) string {
	var geoResp models.GeolocationResponse
	if err := json.Unmarshal(resp, &geoResp); err != nil {
		fmt.Println("An error occured while converting response to xml", err)
		return ""
	}
	xmlStr, err := xml.MarshalIndent(geoResp, "", "  ")
	if err != nil {
		fmt.Println("An error occured while converting response to xml", err)
		return ""
	}
	return string(xmlStr)
}
